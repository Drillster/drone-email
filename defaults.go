package main

const (
	// DefaultPort is the default SMTP port to use
	DefaultPort = 587
	// DefaultOnlyRecipients controls wether to exclude the commit author by default
	DefaultOnlyRecipients = false
	// DefaultSkipVerify controls wether to skip SSL verification for the SMTP server
	DefaultSkipVerify = false
  // DefaultClientHostname is the client hostname used in the HELO command sent to the SMTP server
  DefaultClientHostname = "localhost"
)

// DefaultSubject is the default subject template to use for the email
const DefaultSubject = `
[{{ build.status }}] {{ repo.owner }}/{{ repo.name }} ({{ commit.branch }} - {{ truncate commit.sha 8 }})
`

// DefaultTemplate is the default body template to use for the email
const DefaultTemplate = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta name="viewport" content="width=device-width" />
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <style>
        @media (prefers-color-scheme: light) {
            body {
                background-color: rgb(30, 30, 30);
                color: rgb(208, 208, 208);
            }

            .bg-secondary {
                background-color: #3e3e3e;
            }

            .badge {
                background-color: #505050;
            }
        }

        @media (prefers-color-scheme: dark) {
            body {
                background-color: rgb(241, 241, 241);
                color: black;
            }

            .bg-secondary {
                background-color: #ececec;
                box-shadow: #a5a5a5a5 0px 0px 25px;
            }

            .badge {
                background-color: #d2d2d2;
            }
        }

        .card {
            border-radius: 5px;
            width: calc(100% - 40px);
        }

        .badge {
            border-radius: 2px;
            padding: 3px;
        }

        .flex {
            display: flex;
        }

        .flex-col {
            flex-direction: column;
        }

        .align-start {
            align-self: flex-start;
        }

        .items-start {
            align-items: flex-start;
        }

        .align-center {
            align-items: center;
        }

        .justify-between {
            justify-content: space-between;
        }

        .m-2 {
            margin: 15px;
        }
	
        .alert {
          font-size: 16px;
          color: #fff;
          font-weight: 500;
          padding: 20px;
          text-align: center;
          border-radius: 3px;
        }

        .alert.alert-bad {
          background: #d0021b;
        }

        .alert.alert-good {
          background: #68b90f;
        }

        td {
          margin-right: 5px;
        }
    </style>
</head>
<body>
    <p>
      This is the report of your most recent drone pipeline build.
      It finished {{ datetime build.finished "Mon Jan 2 15:04:05 MST 2006" "Local" }}
    </p>
   
    <!--<img src="https://github.com/JonasBernard/drone-email/raw/master/img/{{build.status}}.png" />-->

   <p>
       {{#success build.status}}
        <td class="alert alert-good">
          Successful build #{{ build.number }}
        </td>
      {{else}}
        <td class="alert alert-bad">
          Failed build #{{ build.number }}
        </td>
      {{/success}}
   </p>
    
    <p>
        The build was based on the following commit:
        <div class="bg-secondary card m-2 flex flex-col">
            <div class="flex justify-between items-start m-2">
                <strong><a href="{{commit.link}}">{{ truncate commit.sha 8 }}: {{commit.message}}</a></strong>
                <small class="badge">{{commit.branch}}</small>
            </div>
            <div class="flex align-center m-2">
                <img class="align-start m-2" src="{{commit.author.avatar}}" style="border-radius: 50%;" width="30px" height="30px" alt="Avatar of {{commit.author.name}}">
                <div class="flex flex-col">
                    <strong>{{commit.author.name}}</strong>
                    {{commit.author.email}}
                </div>
            </div>
        </div>
    </p>
    
    <ul>
        <li>See the build on drone: <a href="{{build.link}}">Build #{{build.number}}, {{build.status}}</a></li>
        <li>See the commit on GitHub: <a href="{{commit.link}}">{{commit.message}}</a></li>
        <li>See the repository on GitHub: <a href="{{repo.link}}">{{repo.fullName}}</a></li>
    </ul>
    <table width="100%" cellpadding="0" cellspacing="0">
    <tr>
      <td>
	See the build on drone::
      </td>
      <td>
	<a href="{{build.link}}">Build #{{build.number}}, {{build.status}}</a>
      </td>
    </tr>
    <tr>
      <td>
	Link to the commit:
      </td>
      <td>
	<a href="{{commit.link}}">{{ truncate commit.sha 8 }}: {{commit.message}}</a>
      </td>
    </tr>
    <tr>
      <td>
	Link to the repository:
      </td>
      <td>
	<a href="{{repo.link}}">{{repo.fullName}}</a>
      </td>
    </tr>
    <tr>
      <td>
	Commit:
      </td>
      <td>
	{{ truncate commit.sha 8 }}
      </td>
    </tr>
    <tr>
      <td>
	Started at:
      </td>
      <td>
	{{ datetime build.created "Mon Jan 2 15:04:05 MST 2006" "Local" }}
      </td>
    </tr>
  </table>
</body>

  <head>
    
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
</html>
`
