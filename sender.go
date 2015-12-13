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
	if len(c.Vargs.Recipients) == 0 {
		c.Vargs.Recipients = []string{
			c.Build.Email,
		}
	}

	var auth smtp.Auth

	var addr = net.JoinHostPort(
		c.Vargs.Host,
		strconv.Itoa(c.Vargs.Port))

	// setup the authentication to the smtp server
	// if the username and password are provided.
	if len(c.Vargs.Username) > 0 {
		auth = smtp.PlainAuth(
			"",
			c.Vargs.Username,
			c.Vargs.Password,
			c.Vargs.Host)
	}

	// genereate the raw email message
	var to = strings.Join(
		c.Vargs.Recipients,
		",")

	var raw = fmt.Sprintf(
		rawMessage,
		c.Vargs.From,
		to,
		subject,
		body)

	return smtp.SendMail(
		addr,
		auth,
		c.Vargs.From,
		c.Vargs.Recipients,
		[]byte(raw))
}
