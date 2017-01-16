package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

func main() {
	// Load env-file if it exists first
	if env := os.Getenv("PLUGIN_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	app := cli.NewApp()
	app.Name = "email plugin"
	app.Usage = "email plugin"
	app.Action = run
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		// Plugin environment
		cli.StringFlag{
			Name:   "from",
			Usage:  "from address",
			EnvVar: "PLUGIN_FROM",
		},
		cli.StringFlag{
			Name:   "host",
			Usage:  "smtp host",
			EnvVar: "EMAIL_HOST,PLUGIN_HOST",
		},
		cli.IntFlag{
			Name:   "port",
			Value:  DefaultPort,
			Usage:  "smtp port",
			EnvVar: "EMAIL_PORT,PLUGIN_PORT",
		},
		cli.StringFlag{
			Name:   "username",
			Usage:  "smtp server username",
			EnvVar: "EMAIL_USERNAME,PLUGIN_USERNAME",
		},
		cli.StringFlag{
			Name:   "password",
			Usage:  "smtp server password",
			EnvVar: "EMAIL_PASSWORD,PLUGIN_PASSWORD",
		},
		cli.BoolFlag{
			Name:   "skip.verify",
			Usage:  "skip tls verify",
			EnvVar: "PLUGIN_SKIP_VERIFY",
		},
		cli.StringSliceFlag{
			Name:   "recipients",
			Usage:  "recipient addresses",
			EnvVar: "EMAIL_RECIPIENTS,PLUGIN_RECIPIENTS",
		},
		cli.BoolFlag{
			Name:   "recipients.only",
			Usage:  "send to recipients only",
			EnvVar: "PLUGIN_RECIPIENTS_ONLY",
		},
		cli.StringFlag{
			Name:   "template.subject",
			Value:  DefaultSubject,
			Usage:  "subject template",
			EnvVar: "PLUGIN_SUBJECT",
		},
		cli.StringFlag{
			Name:   "template.body",
			Value:  DefaultTemplate,
			Usage:  "body template",
			EnvVar: "PLUGIN_BODY",
		},

		// Drone environment
		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "repository owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "git commit sha",
			EnvVar: "DRONE_COMMIT_SHA",
		},
		cli.StringFlag{
			Name:   "commit.ref",
			Value:  "refs/heads/master",
			Usage:  "git commit ref",
			EnvVar: "DRONE_COMMIT_REF",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "commit.author.name",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "commit.author.email",
			Usage:  "git author email",
			EnvVar: "DRONE_COMMIT_AUTHOR_EMAIL",
		},
		cli.StringFlag{
			Name:   "commit.author.avatar",
			Usage:  "git author avatar",
			EnvVar: "DRONE_COMMIT_AUTHOR_AVATAR",
		},
		cli.StringFlag{
			Name:   "commit.message",
			Usage:  "git commit message",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			Value:  "success",
			EnvVar: "DRONE_BUILD_STATUS",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
		cli.Int64Flag{
			Name:   "build.started",
			Usage:  "build started",
			EnvVar: "DRONE_BUILD_STARTED",
		},
		cli.Int64Flag{
			Name:   "build.created",
			Usage:  "build created",
			EnvVar: "DRONE_BUILD_CREATED",
		},
		cli.StringFlag{
			Name:   "build.tag",
			Usage:  "build tag",
			EnvVar: "DRONE_TAG",
		},
		cli.Int64Flag{
			Name:   "job.started",
			Usage:  "job started",
			EnvVar: "DRONE_JOB_STARTED",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Repo: Repo{
			Owner: c.String("repo.owner"),
			Name:  c.String("repo.name"),
		},
		Build: Build{
			Tag:    c.String("build.tag"),
			Number: c.Int("build.number"),
			Event:  c.String("build.event"),
			Status: c.String("build.status"),
			Commit: c.String("commit.sha"),
			Ref:    c.String("commit.ref"),
			Branch: c.String("commit.branch"),
			Author: Author{
				Name:   c.String("commit.author.name"),
				Email:  c.String("commit.author.email"),
				Avatar: c.String("commit.author.avatar"),
			},
			Message: c.String("commit.message"),
			Link:    c.String("build.link"),
			Started: c.Int64("build.started"),
			Created: c.Int64("build.created"),
		},
		Job: Job{
			Started: c.Int64("job.started"),
		},
		Config: Config{
			From:           c.String("from"),
			Host:           c.String("host"),
			Port:           c.Int("port"),
			Username:       c.String("username"),
			Password:       c.String("password"),
			SkipVerify:     c.Bool("skip.verify"),
			Recipients:     c.StringSlice("recipients"),
			RecipientsOnly: c.Bool("recipients.only"),
			Subject:        c.String("template.subject"),
			Body:           c.String("template.body"),
		},
	}

	return plugin.Exec()
}
