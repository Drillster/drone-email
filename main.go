package main

import (
	"os"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

type Email struct {
	Recipients []string `json:"recipients"`

	Host     string `json:"host"`
	Port     string `json:"port"`
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

	err := Send(&Context{
		Email: email,
		Build: build,
		Repo:  repo,
	})
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
