package main

const (
	// DefaultPort is the default SMTP port to use
	DefaultPort = 587
	// DefaultOnlyRecipients controls wether to exclude the commit author by default
	DefaultOnlyRecipients = false
	// DefaultSkipVerify controls wether to skip SSL verification for the SMTP server
	DefaultSkipVerify = false
)

// DefaultSubject is the default subject template to use for the email
const DefaultSubject = `
[{{ build.status }}] {{ repo.owner }}/{{ repo.name }} ({{ build.branch }} - {{ truncate build.commit 8 }})
`

// DefaultTemplate is the default body template to use for the email
const DefaultTemplate = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta name="viewport" content="width=device-width" />
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <style>
      * {
        margin: 0;
        padding: 0;
        font-family: "Helvetica Neue", "Helvetica", Helvetica, Arial, sans-serif;
        box-sizing: border-box;
        font-size: 14px;
      }
      body {
        -webkit-font-smoothing: antialiased;
        -webkit-text-size-adjust: none;
        width: 100% !important;
        height: 100%;
        line-height: 1.6;
        background-color: #f6f6f6;
      }
      table td {
        vertical-align: top;
      }
      .body-wrap {
        background-color: #f6f6f6;
        width: 100%;
      }
      .container {
        display: block !important;
        max-width: 600px !important;
        margin: 0 auto !important;
        /* makes it centered */
        clear: both !important;
      }
      .content {
        max-width: 600px;
        margin: 0 auto;
        display: block;
        padding: 20px;
      }
      .main {
        background: #fff;
        border: 1px solid #e9e9e9;
        border-radius: 3px;
      }
      .content-wrap {
        padding: 20px;
      }
      .content-block {
        padding: 0 0 20px;
      }
      .header {
        width: 100%;
        margin-bottom: 20px;
      }
      h1, h2, h3 {
        font-family: "Helvetica Neue", Helvetica, Arial, "Lucida Grande", sans-serif;
        color: #000;
        margin: 40px 0 0;
        line-height: 1.2;
        font-weight: 400;
      }
      h1 {
        font-size: 32px;
        font-weight: 500;
      }
      h2 {
        font-size: 24px;
      }
      h3 {
        font-size: 18px;
      }
      hr {
        border: 1px solid #e9e9e9;
        margin: 20px 0;
        height: 1px;
        padding: 0;
      }
      p,
      ul,
      ol {
        margin-bottom: 10px;
        font-weight: normal;
      }
      p li,
      ul li,
      ol li {
        margin-left: 5px;
        list-style-position: inside;
      }
      a {
        color: #348eda;
        text-decoration: underline;
      }
      .last {
        margin-bottom: 0;
      }
      .first {
        margin-top: 0;
      }
      .padding {
        padding: 10px 0;
      }
      .aligncenter {
        text-align: center;
      }
      .alignright {
        text-align: right;
      }
      .alignleft {
        text-align: left;
      }
      .clear {
        clear: both;
      }
      .alert {
        font-size: 16px;
        color: #fff;
        font-weight: 500;
        padding: 20px;
        text-align: center;
        border-radius: 3px 3px 0 0;
      }
      .alert a {
        color: #fff;
        text-decoration: none;
        font-weight: 500;
        font-size: 16px;
      }
      .alert.alert-warning {
        background: #ff9f00;
      }
      .alert.alert-bad {
        background: #d0021b;
      }
      .alert.alert-good {
        background: #68b90f;
      }
      @media only screen and (max-width: 640px) {
        h1,
        h2,
        h3 {
          font-weight: 600 !important;
          margin: 20px 0 5px !important;
        }
        h1 {
          font-size: 22px !important;
        }
        h2 {
          font-size: 18px !important;
        }
        h3 {
          font-size: 16px !important;
        }
        .container {
          width: 100% !important;
        }
        .content,
        .content-wrapper {
          padding: 10px !important;
        }
      }
    </style>
  </head>
  <body>
    <table class="body-wrap">
      <tr>
        <td></td>
        <td class="container" width="600">
          <div class="content">
            <table class="main" width="100%" cellpadding="0" cellspacing="0">
              <tr>
                {{#success build.status}}
                  <td class="alert alert-good">
                    <a href="{{ build.link }}">
                      Successful build #{{ build.number }}
                    </a>
                  </td>
                {{else}}
                  <td class="alert alert-bad">
                    <a href="{{ build.link }}">
                      Failed build #{{ build.number }}
                    </a>
                  </td>
                {{/success}}
              </tr>
              <tr>
                <td class="content-wrap">
                  <table width="100%" cellpadding="0" cellspacing="0">
                    <tr>
                      <td>
                        Repo:
                      </td>
                      <td>
                        {{ repo.owner }}/{{ repo.name }}
                      </td>
                    </tr>
                    <tr>
                      <td>
                        Author:
                      </td>
                      <td>
                        {{ build.author.name }} ({{ build.author.email }})
                      </td>
                    </tr>
                    <tr>
                      <td>
                        Branch:
                      </td>
                      <td>
                        {{ build.branch }}
                      </td>
                    </tr>
                    <tr>
                      <td>
                        Commit:
                      </td>
                      <td>
                        {{ truncate build.commit 8 }}
                      </td>
                    </tr>
                    <tr>
                      <td>
                        Started at:
                      </td>
                      <td>
                        {{ datetime build.started "Mon Jan 2 15:04:05 MST 2006" "Local" }}
                      </td>
                    </tr>
                  </table>
                  <hr>
                  <table width="100%" cellpadding="0" cellspacing="0">
                    <tr>
                      <td>
                        {{ build.message }}
                      </td>
                    </tr>
                  </table>
                </td>
              </tr>
            </table>
          </div>
        </td>
        <td></td>
      </tr>
    </table>
  </body>
</html>
`
