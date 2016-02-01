# drone-email

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-email/status.svg)](http://beta.drone.io/drone-plugins/drone-email)
[![Coverage Status](https://aircover.co/badges/drone-plugins/drone-email/coverage.svg)](https://aircover.co/drone-plugins/drone-email)
[![](https://badge.imagelayers.io/plugins/drone-email:latest.svg)](https://imagelayers.io/?images=plugins/drone-email:latest 'Get your own badge on imagelayers.io')

Drone plugin to send build status notifications via Email

## Binary

Build the binary using `make`:

```
make deps build
```

### Example

```sh
./drone-email <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "drone",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com",
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "from": "noreply@foo.com",
        "host": "smtp.mailgun.org",
        "port": 587,
        "username": "",
        "password": "",
        "recipients": [
            "octocat@github.com"
        ]
    }
}
EOF
```

## Docker

Build the container using `make`:

```
make deps docker
```

### Example

```sh
docker run -i plugins/drone-email <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "drone",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "from": "noreply@foo.com",
        "host": "smtp.mailgun.org",
        "port": 587,
        "username": "",
        "password": "",
        "recipients": [
            "octocat@github.com"
        ]
    }
}
EOF
```
