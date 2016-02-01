package main

import (
	"github.com/drone/drone-go/drone"
)

type Params struct {
	Recipients []string `json:"recipients"`
	Host       string   `json:"host"`
	Port       int      `json:"port"`
	From       string   `json:"from"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	Subject    string   `json:"subject"`
	Template   string   `json:"template"`
	SkipVerify bool     `json:"skip_verify"`
}

type Context struct {
	System drone.System
	Repo   drone.Repo
	Build  drone.Build
	Vargs  Params
}
