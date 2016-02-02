package main

import (
	"crypto/tls"

	"github.com/aymerick/douceur/inliner"
	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/template"
	"github.com/go-gomail/gomail"
	"github.com/jaytaylor/html2text"
)

func Send(context *Context) error {
	payload := &drone.Payload{
		System: &context.System,
		Repo:   &context.Repo,
		Build:  &context.Build,
	}

	subject, plain, html, err := build(
		payload,
		context,
	)

	if err != nil {
		return err
	}

	return send(
		subject,
		plain,
		html,
		context,
	)
}

func build(payload *drone.Payload, context *Context) (string, string, string, error) {
	subject, err := template.RenderTrim(
		context.Vargs.Subject,
		payload)

	if err != nil {
		return "", "", "", err
	}

	body, err := template.RenderTrim(
		context.Vargs.Template,
		payload,
	)

	if err != nil {
		return "", "", "", err
	}

	html, err := inliner.Inline(body)

	if err != nil {
		return "", "", "", err
	}

	plain, err := html2text.FromString(
		html,
	)

	if err != nil {
		return "", "", "", err
	}

	return subject, plain, html, nil
}

func send(subject, plainBody, htmlBody string, c *Context) error {
	if len(c.Vargs.Recipients) == 0 {
		c.Vargs.Recipients = []string{
			c.Build.Email,
		}
	}

	m := gomail.NewMessage()

	m.SetHeader(
		"To",
		c.Vargs.Recipients...,
	)

	m.SetHeader(
		"From",
		c.Vargs.From,
	)

	m.SetHeader(
		"Subject",
		subject,
	)

	m.AddAlternative(
		"text/plain",
		plainBody,
	)

	m.AddAlternative(
		"text/html",
		htmlBody,
	)

	d := gomail.NewPlainDialer(
		c.Vargs.Host,
		c.Vargs.Port,
		c.Vargs.Username,
		c.Vargs.Password,
	)

	if c.Vargs.SkipVerify {
		d.TLSConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	return d.DialAndSend(m)
}
