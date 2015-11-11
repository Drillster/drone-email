# drone-email

Drone plugin for sending email notifications.

## Overview

This plugin is responsible for sending build notifications via email:

```
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
        "recipients": ["octocat@github.com"]
    }
}
EOF
```
