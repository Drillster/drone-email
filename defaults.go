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
      {{#success build.status}}Successful build on {{ commit.branch }}{{else}}Failed build on {{ commit.branch }}{{/success}} for {{ repo.owner }}/{{ repo.name }}
`

// DefaultTemplate is the default body template to use for the email
const DefaultTemplate = `
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta name="viewport" content="width=device-width" />
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <meta name="color-scheme" content="dark light">
    <style>
        * {
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
          padding-bottom: 20px;
          top: 0;
          left: 0;
          right: 0;
          bottom: 0;
          padding: 0;
          margin: 0;
        }

        @media (prefers-color-scheme: dark) {
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

            a {
              color: rgb(94, 172, 244);
            }
        }

        @media (prefers-color-scheme: light) {
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

        .content {
          margin: 5px;
          margin-bottom: 15px;
        }

        .card {
            border-radius: 5px;
            width: calc(100% - 30px);
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

        .justify-center {
            justify-content: center;
        }

        .m-2 {
            margin: 15px;
        }
	
        .alert {
          width: calc(100% - 30px);
          font-size: 16px;
          color: #fff;
          font-weight: 500;
          padding: 20px;
          text-align: center;
          border-radius: 3px;
        }

        .alert.alert-bad {
          background-color: #d0021b;
        }

        .alert.alert-good {
          background-color: #68b90f;
        }

        td {
          margin-right: 5px;
        }

        .confetti {
            width: 0;
            height: 0;
        }

        .confetti::before {
            content: '';
            position: absolute;
            top: 0;

            background-image: url(https://github.com/JonasBernard/drone-email/raw/master/img/confetti-transparent.gif);
            background-repeat: no-repeat;
            background-size: contain;

            width: 100vw;
            height: 100vw;

            pointer-events: none;
        }
    </style>
</head>
<body>
    <div class="content">
      <p>
        This is the report of your most recent drone pipeline build.
        It finished {{ datetime build.finished "Mon Jan 2 15:04:05 MST 2006" "Local" }}.
      </p>
      <!--<img src="https://github.com/JonasBernard/drone-email/raw/master/img/{{build.status}}.png" />-->
      
         <p>
         {{#success build.status}}
          <div class="m-2 alert alert-good flex justify-center">
            <span>Successful build #{{ build.number }}</span>
          </div>
        {{else}}
          <div class="m-2 alert alert-bad flex justify-center">
            <span>Failed build #{{ build.number }}</span>
          </div>
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
      
      <table width="100%" cellpadding="0" cellspacing="0">
      <tr>
        <td>
              See the build on drone:
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
        Started at:
        </td>
        <td>
        {{ datetime build.created "Mon Jan 2 15:04:05 MST 2006" "Local" }}
        </td>
      </tr>
        </table>
    </div>
    {{#success build.status}}
    <div class="confetti"></div>
    {{else}}

    {{/success}}
</body>
</html>
`
