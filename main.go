package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

func main() {
	// Load env-file if it exists first
	envFile, envFileSet := os.LookupEnv("PLUGIN_ENV_FILE")
	if !envFileSet {
		envFile = "/run/drone/env"
	}
	if _, err := os.Stat(envFile); err == nil {
		godotenv.Overload(envFile)
	}

	app := cli.NewApp()
	app.Name = "email plugin"
	app.Usage = "email plugin"
	app.Action = run
	app.Version = "2.0.2"
	app.Flags = []cli.Flag{
		// Plugin environment
		cli.StringFlag{
			Name:   "from",
			Usage:  "from address",
			EnvVar: "PLUGIN_FROM",
		},
		cli.StringFlag{
			Name:   "from.address",
			Usage:  "from address",
			EnvVar: "PLUGIN_FROM.ADDRESS",
		},
		cli.StringFlag{
			Name:   "from.name",
			Usage:  "from name",
			EnvVar: "PLUGIN_FROM.NAME",
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
		cli.BoolFlag{
			Name:   "no.starttls",
			Usage:  "Enable/Disable STARTTLS",
			EnvVar: "PLUGIN_NO_STARTTLS",
		},
		cli.StringFlag{
			Name:   "recipients.file",
			Usage:  "file to read recipients from",
			EnvVar: "EMAIL_RECIPIENTS_FILE,PLUGIN_RECIPIENTS_FILE",
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
		cli.StringFlag{
			Name:   "attachment",
			Usage:  "attachment filename",
			EnvVar: "PLUGIN_ATTACHMENT",
		},
		cli.StringSliceFlag{
			Name:   "attachments",
			Usage:  "attachment filename(s)",
			EnvVar: "PLUGIN_ATTACHMENTS",
		},
		cli.StringFlag{
			Name:   "clienthostname",
			Value:  DefaultClientHostname,
			Usage:  "smtp client hostname",
			EnvVar: "EMAIL_CLIENTHOSTNAME,PLUGIN_CLIENTHOSTNAME",
		},

		// Drone environment
		// Repo
		cli.StringFlag{
			Name:   "repo.fullName",
			Usage:  "repository full name",
			EnvVar: "DRONE_REPO",
		},
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
			Name:   "repo.scm",
			Value:  "git",
			Usage:  "respository scm",
			EnvVar: "DRONE_REPO_SCM",
		},
		cli.StringFlag{
			Name:   "repo.link",
			Usage:  "repository link",
			EnvVar: "DRONE_REPO_LINK",
		},
		cli.StringFlag{
			Name:   "repo.avatar",
			Usage:  "repository avatar",
			EnvVar: "DRONE_REPO_AVATAR",
		},
		cli.StringFlag{
			Name:   "repo.branch",
			Value:  "master",
			Usage:  "repository default branch",
			EnvVar: "DRONE_REPO_BRANCH",
		},
		cli.BoolFlag{
			Name:   "repo.private",
			Usage:  "repository is private",
			EnvVar: "DRONE_REPO_PRIVATE",
		},
		cli.BoolFlag{
			Name:   "repo.trusted",
			Usage:  "repository is trusted",
			EnvVar: "DRONE_REPO_TRUSTED",
		},

		// Remote
		cli.StringFlag{
			Name:   "remote.url",
			Usage:  "repository clone url",
			EnvVar: "DRONE_REMOTE_URL",
		},

		// Commit
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
			Name:   "commit.link",
			Usage:  "commit link",
			EnvVar: "DRONE_COMMIT_LINK",
		},
		cli.StringFlag{
			Name:   "commit.message",
			Usage:  "git commit message",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "commit.author.name",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR,DRONE_COMMIT_AUTHOR_NAME",
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

		// Build
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
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
			Name:   "build.created",
			Usage:  "build created",
			EnvVar: "DRONE_BUILD_CREATED",
		},
		cli.Int64Flag{
			Name:   "build.started",
			Usage:  "build started",
			EnvVar: "DRONE_BUILD_STARTED",
		},
		cli.Int64Flag{
			Name:   "build.finished",
			Usage:  "build finished",
			EnvVar: "DRONE_BUILD_FINISHED",
		},

		// Prev
		cli.StringFlag{
			Name:   "prev.build.status",
			Usage:  "prior build status",
			EnvVar: "DRONE_PREV_BUILD_STATUS",
		},
		cli.IntFlag{
			Name:   "prev.build.number",
			Usage:  "prior build number",
			EnvVar: "DRONE_PREV_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "prev.commit.sha",
			Usage:  "prior commit sha",
			EnvVar: "DRONE_PREV_COMMIT_SHA",
		},

		// Job
		cli.IntFlag{
			Name:   "job.number",
			Usage:  "job number",
			EnvVar: "DRONE_JOB_NUMBER",
		},
		cli.StringFlag{
			Name:   "job.status",
			Usage:  "job status",
			EnvVar: "DRONE_JOB_STATUS",
		},
		cli.IntFlag{
			Name:   "job.exitCode",
			Usage:  "job exit code",
			EnvVar: "DRONE_JOB_EXIT_CODE",
		},
		cli.Int64Flag{
			Name:   "job.started",
			Usage:  "job started",
			EnvVar: "DRONE_JOB_STARTED",
		},
		cli.Int64Flag{
			Name:   "job.finished",
			Usage:  "job finished",
			EnvVar: "DRONE_JOB_FINISHED",
		},

		// Yaml
		cli.BoolFlag{
			Name:   "yaml.signed",
			Usage:  "yaml is signed",
			EnvVar: "DRONE_YAML_SIGNED",
		},
		cli.BoolFlag{
			Name:   "yaml.verified",
			Usage:  "yaml is signed and verified",
			EnvVar: "DRONE_YAML_VERIFIED",
		},

		// Tag
		cli.StringFlag{
			Name:   "tag",
			Usage:  "git tag",
			EnvVar: "DRONE_TAG",
		},

		// PullRequest
		cli.IntFlag{
			Name:   "pullRequest",
			Usage:  "pull request number",
			EnvVar: "DRONE_PULL_REQUEST",
		},

		// DeployTo
		cli.StringFlag{
			Name:   "deployTo",
			Usage:  "deployment target",
			EnvVar: "DRONE_DEPLOY_TO",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func run(c *cli.Context) error {

	var fromAddress string = c.String("from")
	if fromAddress == "" {
		fromAddress = c.String("from.address")
	}

	plugin := Plugin{
		Repo: Repo{
			FullName: c.String("repo.fullName"),
			Owner:    c.String("repo.owner"),
			Name:     c.String("repo.name"),
			SCM:      c.String("repo.scm"),
			Link:     c.String("repo.link"),
			Avatar:   c.String("repo.avatar"),
			Branch:   c.String("repo.branch"),
			Private:  c.Bool("repo.private"),
			Trusted:  c.Bool("repo.trusted"),
		},
		Remote: Remote{
			URL: c.String("remote.url"),
		},
		Commit: Commit{
			Sha:     c.String("commit.sha"),
			Ref:     c.String("commit.ref"),
			Branch:  c.String("commit.branch"),
			Link:    c.String("commit.link"),
			Message: c.String("commit.message"),
			Author: Author{
				Name:   c.String("commit.author.name"),
				Email:  c.String("commit.author.email"),
				Avatar: c.String("commit.author.avatar"),
			},
		},
		Build: Build{
			Number:   c.Int("build.number"),
			Event:    c.String("build.event"),
			Status:   c.String("build.status"),
			Link:     c.String("build.link"),
			Created:  c.Int64("build.created"),
			Started:  c.Int64("build.started"),
			Finished: c.Int64("build.finished"),
		},
		Prev: Prev{
			Build: PrevBuild{
				Status: c.String("prev.build.status"),
				Number: c.Int("prev.build.number"),
			},
			Commit: PrevCommit{
				Sha: c.String("prev.commit.sha"),
			},
		},
		Job: Job{
			Status:   c.String("job.status"),
			ExitCode: c.Int("job.exitCode"),
			Started:  c.Int64("job.started"),
			Finished: c.Int64("job.finished"),
		},
		Yaml: Yaml{
			Signed:   c.Bool("yaml.signed"),
			Verified: c.Bool("yaml.verified"),
		},
		Tag:         c.String("tag"),
		PullRequest: c.Int("pullRequest"),
		DeployTo:    c.String("deployTo"),
		Config: Config{
			FromAddress:    fromAddress,
			FromName:       c.String("from.name"),
			Host:           c.String("host"),
			Port:           c.Int("port"),
			Username:       c.String("username"),
			Password:       c.String("password"),
			SkipVerify:     c.Bool("skip.verify"),
			NoStartTLS:     c.Bool("no.starttls"),
			Recipients:     c.StringSlice("recipients"),
			RecipientsFile: c.String("recipients.file"),
			RecipientsOnly: c.Bool("recipients.only"),
			Subject:        c.String("template.subject"),
			Body:           c.String("template.body"),
			Attachment:     c.String("attachment"),
			Attachments:    c.StringSlice("attachments"),
			ClientHostname: c.String("clienthostname"),
		},
	}

	return plugin.Exec()
}
