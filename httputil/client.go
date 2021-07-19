// Copyright (C) 2021  Shanhu Tech Inc.
//
// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU Affero General Public License as published by the
// Free Software Foundation, either version 3 of the License, or (at your
// option) any later version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU Affero General Public License
// for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package httputil

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client performs client that calls to a remote server with an optional token.
type Client struct {
	Server *url.URL

	// TokenSource is an optional token source to proivde bearer token.
	TokenSource TokenSource
	// Token is the optional token to use a bearer token, used only when
	// TokenSource is nil.
	Token string

	UserAgent string // Optional User-Agent for each request.
	Accept    string // Optional Accept header.

	Transport http.RoundTripper
}

func (c *Client) addAuth(req *http.Request) error {
	if c.TokenSource != nil {
		ctx := req.Context()
		tok, err := c.TokenSource.Token(ctx)
		if err != nil {
			return err
		}
		SetAuthToken(req.Header, tok)
		return nil
	}
	SetAuthToken(req.Header, c.Token)
	return nil
}

func (c *Client) addHeaders(h http.Header) {
	setHeader(h, "User-Agent", c.UserAgent)
	setHeader(h, "Accept", c.Accept)
}

func (c *Client) makeClient() *http.Client {
	return &http.Client{Transport: c.Transport}
}

func (c *Client) doRaw(req *http.Request) (*http.Response, error) {
	return c.makeClient().Do(req)
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	resp, err := c.doRaw(req)
	if err != nil {
		return nil, err
	}
	if !isSuccess(resp) {
		defer resp.Body.Close()
		return nil, RespError(resp)
	}
	return resp, nil
}

func (c *Client) req(m, p string, r io.Reader) (*http.Request, error) {
	u, err := makeURL(c.Server, p)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(m, u, r)
	if err != nil {
		return nil, err
	}
	if err := c.addAuth(req); err != nil {
		return nil, err
	}
	c.addHeaders(req.Header)
	return req, nil
}

func (c *Client) reqJSON(m, p string, r io.Reader) (*http.Request, error) {
	req, err := c.req(m, p, r)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

// Put puts a stream to a path on the server.
func (c *Client) Put(p string, r io.Reader) error {
	req, err := c.req(http.MethodPut, p, r)
	if err != nil {
		return err
	}

	resp, err := c.do(req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}

// PutBytes puts bytes to a path on the server.
func (c *Client) PutBytes(p string, bs []byte) error {
	return c.Put(p, bytes.NewBuffer(bs))
}

// JSONPut puts an object in JSON encoding.
func (c *Client) JSONPut(p string, v interface{}) error {
	bs, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return c.PutBytes(p, bs)
}

func (c *Client) poke(m, p string) error {
	req, err := c.req(m, p, nil)
	if err != nil {
		return err
	}

	resp, err := c.do(req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}

// GetCode gets a response from a route and returns the
// status code.
func (c *Client) GetCode(p string) (int, error) {
	req, err := c.req(http.MethodGet, p, nil)
	if err != nil {
		return 0, err
	}
	resp, err := c.doRaw(req)
	if err != nil {
		return 0, err
	}
	code := resp.StatusCode
	resp.Body.Close()
	return code, nil
}

// Poke posts a signal to the given route on the server.
func (c *Client) Poke(p string) error {
	return c.poke(http.MethodPost, p)
}

// Get gets a response from a route on the server.
func (c *Client) Get(p string) (*http.Response, error) {
	req, err := c.req(http.MethodGet, p, nil)
	if err != nil {
		return nil, err
	}
	return c.do(req)
}

// GetString gets the string response from a route on the server.
func (c *Client) GetString(p string) (string, error) {
	resp, err := c.Get(p)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return respString(resp)
}

// GetInto gets the specified path and writes everything from the body to the
// given writer.
func (c *Client) GetInto(p string, w io.Writer) (int64, error) {
	resp, err := c.Get(p)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return io.Copy(w, resp.Body)
}

// GetBytes gets the byte array from a route on the server.
func (c *Client) GetBytes(p string) ([]byte, error) {
	resp, err := c.Get(p)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// JSONGet gets the content of a path and decodes the response
// into resp as JSON.
func (c *Client) JSONGet(p string, resp interface{}) error {
	req, err := c.reqJSON(http.MethodGet, p, nil)
	if err != nil {
		return nil
	}
	httpResp, err := c.do(req)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	dec := json.NewDecoder(httpResp.Body)
	if err := dec.Decode(resp); err != nil {
		return err
	}
	return httpResp.Body.Close()
}

func copyRespBody(resp *http.Response, w io.Writer) error {
	defer resp.Body.Close()
	if w == nil {
		return nil
	}
	if _, err := io.Copy(w, resp.Body); err != nil {
		return err
	}
	return resp.Body.Close()
}

// Post posts with request body from r, and copies the response body
// to w.
func (c *Client) Post(p string, r io.Reader, w io.Writer) error {
	if r != nil {
		r = ioutil.NopCloser(r)
	}
	req, err := c.req(http.MethodPost, p, r)
	if err != nil {
		return err
	}
	resp, err := c.do(req)
	if err != nil {
		return err
	}
	return copyRespBody(resp, w)
}

func (c *Client) jsonPost(p string, req interface{}) (*http.Response, error) {
	bs, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := c.reqJSON(http.MethodPost, p, bytes.NewBuffer(bs))
	if err != nil {
		return nil, err
	}
	return c.do(httpReq)
}

// JSONPost posts a JSON object as the request body and writes the body
// into the given writer.
func (c *Client) JSONPost(p string, req interface{}, w io.Writer) error {
	resp, err := c.jsonPost(p, req)
	if err != nil {
		return err
	}
	return copyRespBody(resp, w)
}

// JSONCall performs a call with the request as a marshalled JSON object,
// and the response unmarhsalled as a JSON object.
func (c *Client) JSONCall(p string, req, resp interface{}) error {
	httpResp, err := c.jsonPost(p, req)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if resp == nil {
		return nil
	}
	dec := json.NewDecoder(httpResp.Body)
	if err := dec.Decode(resp); err != nil {
		return err
	}
	return httpResp.Body.Close()
}

// Call is an alias to JSONCall.
func (c *Client) Call(p string, req, resp interface{}) error {
	return c.JSONCall(p, req, resp)
}

// Delete sends a delete message to the particular path.
func (c *Client) Delete(p string) error {
	return c.poke(http.MethodDelete, p)
}
