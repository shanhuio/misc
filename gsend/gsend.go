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

package gsend

import (
	"bytes"
	"fmt"
	"net/smtp"
)

// Mail is a simple email.
type Mail struct {
	From    string
	To      string
	Subject string
	Message string
}

const server = "smtp.gmail.com"

// Send sends an email via gmail server.
func Send(m *Mail, password string) error {
	body := new(bytes.Buffer)
	fmt.Fprintf(body, "To: %s\r\n", m.To)
	fmt.Fprintf(body, "Subject: %s\r\n", m.Subject)
	fmt.Fprintf(body, "\r\n")
	fmt.Fprintf(body, "%s", m.Message)

	auth := smtp.PlainAuth("", m.From, password, server)
	return smtp.SendMail(
		fmt.Sprintf("%s:587", server),
		auth,
		m.From, []string{m.To},
		body.Bytes(),
	)
}
