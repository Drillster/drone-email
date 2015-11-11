package main

import (
	"fmt"
	"os"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

type Email struct {
	Recipients []string `json:"recipients"`

	Host     string `json:"host"`
	Port     int    `json:"port"`
	From     string `json:"from"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Context struct {
	Email  Email
	Build  drone.Build
	Repo   drone.Repo
	System drone.System
}

func main() {
	repo := drone.Repo{}
	build := drone.Build{}
	system := drone.System{}
	email := Email{}

	plugin.Param("system", &system)
	plugin.Param("build", &build)
	plugin.Param("repo", &repo)
	plugin.Param("vargs", &email)
	err := plugin.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// default to the commit email address
	// if no email recipient list is provided
	if len(email.Recipients) == 0 {
		email.Recipients = []string{
			build.Email,
		}
	}

	// default smtp port
	if email.Port == 0 {
		email.Port = 587
	}

	err = Send(&Context{
		Email:  email,
		Build:  build,
		Repo:   repo,
		System: system,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
