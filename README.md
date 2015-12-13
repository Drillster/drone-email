# drone-email

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-email/status.svg)](http://beta.drone.io/drone-plugins/drone-email)
[![](https://badge.imagelayers.io/plugins/drone-email:latest.svg)](https://imagelayers.io/?images=plugins/drone-email:latest 'Get your own badge on imagelayers.io')

Drone plugin for sending build status notifications via Email

## Usage

```sh
./drone-email <<EOF
{
    "repo" : {
        "owner": "foo",
        "name": "bar",
        "full_name": "foo/bar"
    },
    "system": {
        "link_url": "http://drone.mycompany.com"
    },
    "build" : {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "commit": "64908ed2414b771554fda6508dd56a0c43766831",
        "branch": "master",
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
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

Build the Docker container using `make`:

```sh
make deps build
docker build --rm=true -t plugins/drone-email .
```

### Example

```sh
docker run -i plugins/drone-email <<EOF
{
    "repo" : {
        "owner": "foo",
        "name": "bar",
        "full_name": "foo/bar"
    },
    "system": {
        "link_url": "http://drone.mycompany.com"
    },
    "build" : {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "commit": "64908ed2414b771554fda6508dd56a0c43766831",
        "branch": "master",
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
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
