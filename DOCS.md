Use this plugin for sending build status notifications via Email. You can
override the default configuration with the following parameters:

* `from` - Send notifications from this address
* `host` - SMTP server host
* `port` - SMTP server port, defaults to `587`
* `username` - SMTP username
* `password` - SMTP password
* `recipients` - List of recipients, defaults to commit email

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
notify:
  email:
    from: noreply@github.com
    host: smtp.mailgun.org
    username: octocat
    password: 12345
    recipients:
      - octocat@github.com
```
