package main

import (
	"bufio"
	"crypto/tls"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/aymerick/douceur/inliner"
	"github.com/drone/drone-go/template"
	"github.com/jaytaylor/html2text"
	gomail "gopkg.in/mail.v2"
)

type (
	Repo struct {
		FullName string
		Owner    string
		Name     string
		SCM      string
		Link     string
		Avatar   string
		Branch   string
		Private  bool
		Trusted  bool
	}

	Remote struct {
		URL string
	}

	Author struct {
		Name   string
		Email  string
		Avatar string
	}

	Commit struct {
		Sha     string
		Ref     string
		Branch  string
		Link    string
		Message string
		Author  Author
	}

	Build struct {
		Number   int
		Event    string
		Status   string
		Link     string
		Created  int64
		Started  int64
		Finished int64
	}

	PrevBuild struct {
		Status string
		Number int
	}

	PrevCommit struct {
		Sha string
	}

	Prev struct {
		Build  PrevBuild
		Commit PrevCommit
	}

	Job struct {
		Status   string
		ExitCode int
		Started  int64
		Finished int64
	}

	Yaml struct {
		Signed   bool
		Verified bool
	}

	Config struct {
		FromAddress    string
		FromName       string
		Host           string
		Port           int
		Username       string
		Password       string
		SkipVerify     bool
		NoStartTLS     bool
		Recipients     []string
		RecipientsFile string
		RecipientsOnly bool
		Subject        string
		Body           string
		Attachment     string
		Attachments    []string
		ClientHostname string
	}

	Plugin struct {
		Repo        Repo
		Remote      Remote
		Commit      Commit
		Build       Build
		Prev        Prev
		Job         Job
		Yaml        Yaml
		Tag         string
		PullRequest int
		DeployTo    string
		Config      Config
	}
)

// Exec will send emails over SMTP
func (p Plugin) Exec() error {
	var dialer *gomail.Dialer

	if !p.Config.RecipientsOnly {
		exists := false
		for _, recipient := range p.Config.Recipients {
			if recipient == p.Commit.Author.Email {
				exists = true
			}
		}

		if !exists {
			p.Config.Recipients = append(p.Config.Recipients, p.Commit.Author.Email)
		}
	}

	if p.Config.RecipientsFile != "" {
		f, err := os.Open(p.Config.RecipientsFile)
		if err == nil {
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				p.Config.Recipients = append(p.Config.Recipients, scanner.Text())
			}
		} else {
			log.Errorf("Could not open RecipientsFile %s: %v", p.Config.RecipientsFile, err)
		}
	}

	if p.Config.Username == "" && p.Config.Password == "" {
		dialer = &gomail.Dialer{Host: p.Config.Host, Port: p.Config.Port}
	} else {
		dialer = gomail.NewDialer(p.Config.Host, p.Config.Port, p.Config.Username, p.Config.Password)
	}

	if p.Config.SkipVerify {
		dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	if p.Config.NoStartTLS {
		dialer.StartTLSPolicy = gomail.NoStartTLS
	}

	dialer.LocalName = p.Config.ClientHostname

	closer, err := dialer.Dial()
	if err != nil {
		log.Errorf("Error while dialing SMTP server: %v", err)
		return err
	}

	type Context struct {
		Repo        Repo
		Remote      Remote
		Commit      Commit
		Build       Build
		Prev        Prev
		Job         Job
		Yaml        Yaml
		Tag         string
		PullRequest int
		DeployTo    string
	}
	ctx := Context{
		Repo:        p.Repo,
		Remote:      p.Remote,
		Commit:      p.Commit,
		Build:       p.Build,
		Prev:        p.Prev,
		Job:         p.Job,
		Yaml:        p.Yaml,
		Tag:         p.Tag,
		PullRequest: p.PullRequest,
		DeployTo:    p.DeployTo,
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
		if len(recipient) == 0 {
			continue
		}
		message.SetAddressHeader("From", p.Config.FromAddress, p.Config.FromName)
		message.SetAddressHeader("To", recipient, "")
		message.SetHeader("Subject", subject)
		message.AddAlternative("text/plain", plainBody)
		message.AddAlternative("text/html", html)

		if p.Config.Attachment != "" {
			attach(message, p.Config.Attachment)
		}

		for _, attachment := range p.Config.Attachments {
			attach(message, attachment)
		}

		if err := gomail.Send(closer, message); err != nil {
			log.Errorf("Could not send email to %q: %v", recipient, err)
			return err
		}
		message.Reset()
	}

	return nil
}

func attach(message *gomail.Message, attachment string) {
	if _, err := os.Stat(attachment); err == nil {
		message.Attach(attachment)
	}
}
