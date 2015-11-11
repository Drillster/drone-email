Use the Email plugin to send an email to a recipient list when a build completes. The following parameters are used to configuration the email plugin:

* **from** - sends email from this address
* **host** - smtp server host
* **port** - smtp server port (defaults to `587`)
* **username** - smtp server username
* **password** - smtp server password
* **recipient** - list of email recipients (defaults to commit email)

Same email configuration:

```yaml
notify:
  email:
    from: noreply@github.com
    host: smtp.mailgun.org
    port: 587
    username: octocat
    password: 12345
    recipients:
      - octocat@github.com
```
