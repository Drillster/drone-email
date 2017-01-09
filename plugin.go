package main

import (
	"crypto/tls"
	log "github.com/Sirupsen/logrus"
	"github.com/aymerick/douceur/inliner"
	"github.com/drone/drone-go/template"
	"github.com/jaytaylor/html2text"
	"gopkg.in/gomail.v2"
)

type (
	Repo struct {
		Owner string
		Name  string
	}

	Author struct {
		Name   string
		Email  string
		Avatar string
	}

	Build struct {
		Tag     string
		Event   string
		Number  int
		Commit  string
		Ref     string
		Branch  string
		Author  Author
		Message string
		Status  string
		Link    string
		Started int64
		Created int64
	}

	Config struct {
		From           string
		Host           string
		Port           int
		Username       string
		Password       string
		SkipVerify     bool
		Recipients     []string
		RecipientsOnly bool
		Subject        string
		Body           string
	}

	Job struct {
		Started int64
	}

	Plugin struct {
		Repo   Repo
		Build  Build
		Config Config
		Job    Job
	}
)

// Exec will send emails over SMTP
func (p Plugin) Exec() error {
	var dialer *gomail.Dialer

	if !p.Config.RecipientsOnly {
		p.Config.Recipients = append(p.Config.Recipients, p.Build.Author.Email)
	}

	if p.Config.Username == "" && p.Config.Password == "" {
		dialer = &gomail.Dialer{Host: p.Config.Host, Port: p.Config.Port}
	} else {
		dialer = gomail.NewDialer(p.Config.Host, p.Config.Port, p.Config.Username, p.Config.Password)
	}
	if p.Config.SkipVerify {
		dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	closer, err := dialer.Dial()
	if err != nil {
		log.Errorf("Error while dialing SMTP server: %v", err)
		return err
	}

	type Context struct {
		Job    interface{}
		Repo   interface{}
		Build  interface{}
		Config interface{}
	}
	ctx := Context{
		Job:    p.Job,
		Repo:   p.Repo,
		Build:  p.Build,
		Config: p.Config,
	}

	// Render body in HTML and plain text
	renderedBody, err := template.RenderTrim(p.Config.Body, ctx)
	if err != nil {
		log.Errorf("Could not render body template: %v", err)
		return err
	}
	html, err := inliner.Inline(renderedBody)
	if err != nil {
		log.Errorf("Could not inline rendered body: %v", err)
		return err
	}
	plainBody, err := html2text.FromString(html)
	if err != nil {
		log.Errorf("Could not convert html to text: %v", err)
		return err
	}

	// Render subject
	subject, err := template.RenderTrim(p.Config.Subject, ctx)
	if err != nil {
		log.Errorf("Could not render subject template: %v", err)
		return err
	}

	// Send emails
	message := gomail.NewMessage()
	for _, recipient := range p.Config.Recipients {
		message.SetHeader("From", p.Config.From)
		message.SetAddressHeader("To", recipient, "")
		message.SetHeader("Subject", subject)
		message.AddAlternative("text/plain", plainBody)
		message.AddAlternative("text/html", html)

		if err := gomail.Send(closer, message); err != nil {
			log.Errorf("Could not send email to %q: %v", recipient, err)
		}
		message.Reset()
	}

	return nil
}
