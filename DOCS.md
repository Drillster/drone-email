Use the Email plugin for sending build status notifications via email.

## Config
You can configure the plugin using the following parameters:

* **from.address** - Send notifications from this address
* **from.name** - Notifications sender name
* **host** - SMTP server host
* **port** - SMTP server port, defaults to `587`
* **username** - SMTP username
* **password** - SMTP password
* **skip_verify** - Skip verification of SSL certificates, defaults to `false`
* **no_starttls** - Enable/Disable STARTTLS
* **recipients** - List of recipients to send this mail to (besides the commit author)
* **recipients_file** - Filename to load additional recipients from (textfile with one email per line) (besides the commit author)
* **recipients_only** - Do not send mails to the commit author, but only to **recipients**, defaults to `false`
* **subject** - The subject line template
* **body** - The email body template
* **attachment** - An optional file to attach to the sent mail(s), can be an absolute path or relative to the working directory.

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
kind: pipeline
type: docker
name: default

steps:
  - name: notify
    image: drillster/drone-email
    settings:
      from.address: noreply@github.com
      from.name: John Smith
      host: smtp.mailgun.org
      username: octocat
      password: 12345
      recipients:
        - octocat@github.com
```

### Secrets
The Email plugin supports reading credentials and other parameters from the Drone secret store. This is strongly recommended instead of storing credentials in the pipeline configuration in plain text.

```diff
steps:
  - name: notify:
    image: drillster/drone-email
    settings:
      from.address: noreply@github.com
      host: smtp.mailgun.org
+     username:
+       from_secret: email_username
+     password: 12345
+       from_secret: email_password
      recipients:
        - octocat@github.com
```

Use the command line utility to add the secrets to the store:

```sh
drone secret add \
  --repository octocat/hello-world \
  --name email_username \
  --data octocat
drone secret add \
  --repository octocat/hello-world \
  --name email_password \
  --data 12345
```

See [Secret Guide](https://docs.drone.io/secret/) for additional information on secrets.

### Custom Templates

In some cases you may want to customize the look and feel of the email message
so you can use custom templates. For the use case we expose the following
additional parameters, all of the accept a custom handlebars template, directly
provided as a string or as a remote URL which gets fetched and parsed:

* **subject** - A handlebars template to create a custom subject. For more
  details take a look at the [docs](http://handlebarsjs.com/). You can see the
  default template [here](https://github.com/Drillster/drone-email/blob/master/defaults.go#L14)
* **body** - A handlebars template to create a custom template. For more
  details take a look at the [docs](http://handlebarsjs.com/). You can see the
  default template [here](https://github.com/Drillster/drone-email/blob/master/defaults.go#L19-L267)

Example configuration that generate a custom email:

```yaml
steps:
  - name: notify
    image: drillster/drone-email
    settings:
      from.address: noreply@github.com
      host: smtp.mailgun.org
      username: octocat
      password: 12345
      subject: >
        [{{ build.status }}]
        {{ repo.owner }}/{{ repo.name }}
        ({{ build.branch }} - {{ truncate build.commit 8 }})
      body:
        https://git.io/vgvPz
```

### Skip SSL verify

In some cases you may want to skip SSL verification, even if we discourage that
as it leads to an unsecure environment. Please use this option only within your
intranet and/or with truested resources. For this use case we expose the
following additional parameter:

* **skip_verify** - Skip verification of SSL certificates

Example configuration that skips SSL verification:

```diff
steps:
  - name: notify
    image: drillster/drone-email
    settings:
      from: noreply@github.com
      host: smtp.mailgun.org
      username: octocat
      password: 12345
+     skip_verify: true
```

### STARTTLS

By default, STARTTLS is being used opportunistically meaning, if advertised
by the server, traffic is going to be encrypted.

You may want to disable STARTTLS, e.g., with faulty and/or internal servers:

```diff
steps:
  - name: notify
    image: drillster/drone-email
    settings:
      from: noreply@github.com
      host: smtp.mailgun.org
      username: octocat
      password: 12345
+     no_starttls: true
```
