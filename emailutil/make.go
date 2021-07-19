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

package emailutil

import (
	"bytes"
	"fmt"
	"html/template"
	"time"
)

// Address creates a new email address.
func Address(name, email string) string {
	if name == "" {
		return email
	}
	return fmt.Sprintf("%s <%s>", name, email)
}

// Header is an email header.
type Header struct {
	From    string
	To      string
	Cc      string
	Bcc     string
	Subject string
	Time    time.Time
}

func printHeader(b *bytes.Buffer, k, v string) {
	fmt.Fprintf(b, "%s: %s\r\n", k, v)
}

// Make creates an email with the given header and body.
func Make(h *Header, body []byte) []byte {
	b := new(bytes.Buffer)
	printHeader(b, "Date", h.Time.String())
	printHeader(b, "From", h.From)
	printHeader(b, "To", h.To)
	if h.Cc != "" {
		printHeader(b, "CC", h.Cc)
	}
	if h.Bcc != "" {
		printHeader(b, "BCC", h.Bcc)
	}
	printHeader(b, "Subject", h.Subject)
	printHeader(b, "MIME-Version", "1.0;")
	printHeader(b, "Content-Type", `text/html; charset="UTF-8"`)
	fmt.Fprint(b, "\r\n")
	b.Write(body)
	return b.Bytes()
}

// TemplateMake creates an email using the given template
func TemplateMake(
	h *Header, t *template.Template, dat interface{},
) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := t.Execute(buf, dat); err != nil {
		return nil, err
	}
	return Make(h, buf.Bytes()), nil
}
