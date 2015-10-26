package main

import (
	"html/template"
)

// raw email message template
var rawMessage = `From: %s
To: %s
Subject: %s
MIME-version: 1.0
Content-Type: text/html; charset="UTF-8"
%s`

// default success email template
var successTemplate = template.Must(template.New("_").Parse(`
<p>
	<b>Build was Successful</b>
	(<a href="{{.System.Link}}/{{.Repo.Owner}}/{{.Repo.Name}}/{{.Build.Number}}">see results</a>)
</p>
<p>Repository : {{.Repo.Owner}}/{{.Repo.Name}}</p>
<p>Commit     : {{.Build.Commit}}</p>
<p>Author     : {{.Build.Author}}</p>
<p>Branch     : {{.Build.Branch}}</p>
<p>Message:</p>
<p>{{ .Build.Message }}</p>
`))

// default failure email template
var failureTemplate = template.Must(template.New("_").Parse(`
<p>
	<b>Build Failed</b>
	(<a href="{{.System.Link}}/{{.Repo.Owner}}/{{.Repo.Name}}/{{.Build.Number}}">see results</a>)
</p>
<p>Repository : {{.Repo.Owner}}/{{.Repo.Name}}</p>
<p>Commit     : {{.Build.Commit}}</p>
<p>Author     : {{.Build.Author}}</p>
<p>Branch     : {{.Build.Branch}}</p>
<p>Message:</p>
<p>{{ .Build.Message }}</p>
`))
