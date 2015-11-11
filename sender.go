package main

import (
	"bytes"
	"fmt"
	"net"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/drone/drone-go/drone"
)

const (
	Subject = "[%s] %s/%s (%s - %s)"
)

func Send(context *Context) error {
	switch context.Build.Status {
	case drone.StatusSuccess:
		return SendSuccess(context)
	default:
		return SendFailure(context)
	}
}

// SendFailure sends email notifications to the list of
// recipients indicating the build failed.
func SendFailure(context *Context) error {

	// generate the email failure template
	var buf bytes.Buffer
	err := failureTemplate.ExecuteTemplate(&buf, "_", context)
	if err != nil {
		return err
	}

	// generate the email subject
	var subject = fmt.Sprintf(
		Subject,
		context.Build.Status,
		context.Repo.Owner,
		context.Repo.Name,
		context.Build.Branch,
		context.Build.Commit[:8],
	)

	return send(subject, buf.String(), context)
}

// SendSuccess sends email notifications to the list of
// recipients indicating the build was a success.
func SendSuccess(context *Context) error {

	// generate the email success template
	var buf bytes.Buffer
	err := successTemplate.ExecuteTemplate(&buf, "_", context)
	if err != nil {
		return err
	}

	// generate the email subject
	var subject = fmt.Sprintf(
		Subject,
		context.Build.Status,
		context.Repo.Owner,
		context.Repo.Name,
		context.Build.Branch,
		context.Build.Commit[:8],
	)

	return send(subject, buf.String(), context)
}

func send(subject, body string, c *Context) error {

	if len(c.Email.Recipients) == 0 {
		c.Email.Recipients = []string{
			c.Build.Email,
		}
	}

	var auth smtp.Auth
	var addr = net.JoinHostPort(c.Email.Host, strconv.Itoa(c.Email.Port))

	// setup the authentication to the smtp server
	// if the username and password are provided.
	if len(c.Email.Username) > 0 {
		auth = smtp.PlainAuth("", c.Email.Username, c.Email.Password, c.Email.Host)
	}

	// genereate the raw email message
	var to = strings.Join(c.Email.Recipients, ",")
	var raw = fmt.Sprintf(rawMessage, c.Email.From, to, subject, body)

	return smtp.SendMail(addr, auth, c.Email.From, c.Email.Recipients, []byte(raw))
}
