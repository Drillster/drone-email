Use the Email plugin for sending build status notifications via email.

## Config
You can configure the plugin using the following parameters:

* **from** - Send notifications from this address
* **host** - SMTP server host
* **port** - SMTP server port, defaults to `587`
* **username** - SMTP username
* **password** - SMTP password
* **skip_verify** - Skip verification of SSL certificates, defaults to `false`
* **recipients** - List of recipients to send this mail to (besides the commit author)
* **recipients_only** - Do not send mails to the commit author, but only to **recipients**, defaults to `false`
* **subject** - The subject line template
* **body** - The email body template

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
pipeline:
  notify:
    image: drillster/drone-email
    from: noreply@github.com
    host: smtp.mailgun.org
    username: octocat
    password: 12345
    recipients:
      - octocat@github.com
```

### Secrets
The Email plugin supports reading credentials and other parameters from the Drone secret store. This is strongly recommended instead of storing credentials in the pipeline configuration in plain text.

```diff
pipeline:
  notify:
    image: drillster/drone-email
    from: noreply@github.com
    host: smtp.mailgun.org
-   username: octocat
-   password: 12345
    recipients:
      - octocat@github.com
```

Use the command line utility to add the secrets to the store:

```sh
drone secret add --image=drillster/drone-email \
    octocat/hello-world EMAIL_USERNAME octocat
drone secret add --image=drillster/drone-email \
    octocat/hello-world EMAIL_PASSWORD 12345
```

Then sign the YAML file after all secrets are added:

```sh
drone sign octocat/hello-world
```

The following secret values can be set to configure the plugin:
* **EMAIL_HOST** - corresponds to **host**
* **EMAIL_PORT** - corresponds to **port**
* **EMAIL_USERNAME** - corresponds to **username**
* **EMAIL_PASSWORD** - corresponds to **password**
* **EMAIL_RECIPIENTS** - corresponds to **recipients**

See [Secret Guide](http://readme.drone.io/usage/secret-guide/) for additional information on secrets.

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
pipeline:
  notify:
    image: drillster/drone-email
    from: noreply@github.com
    host: smtp.mailgun.org
    username: octocat
    password: 12345
    subject: >
      [{{ build.status }}]
      {{ repo.owner }}/{{ repo.name }}
      ({{ build.branch }} - {{ truncate build.commit 8 }})
    template: >
      https://git.io/vgvPz
```

### Skip SSL verify

In some cases you may want to skip SSL verification, even if we discourage that
as it leads to an unsecure environment. Please use this option only within your
intranet and/or with truested resources. For this use case we expose the
following additional parameter:

* **skip_verify** - Skip verification of SSL certificates

Example configuration that skips SSL verification:

```yaml
pipeline:
  notify:
    image: drillster/drone-email
    from: noreply@github.com
    host: smtp.mailgun.org
    username: octocat
    password: 12345
    skip_verify: true
```
