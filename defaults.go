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
            width: 100% !important;
            height: 100%;
            line-height: 1.6;
            padding-bottom: 20px;
            padding: 0;
            margin: 0;

            /* Not working in GMail */
            -webkit-font-smoothing: antialiased;
            -webkit-text-size-adjust: none;
            top: 0;
            left: 0;
            right: 0;
            bottom: 0;
        }

        @media (prefers-color-scheme: dark) {
            body {
                background-color: rgb(30, 30, 30);
                color: rgb(208, 208, 208);
            }

            .bg-secondary {
                background-color: #262626;
            }

            .badge {
                background-color: #393939;
            }

            a {
                color: rgb(94, 172, 244);
            }

            .card-header {
                background-color: #343434;
            }
        }

        @media (prefers-color-scheme: light) {
            body {
                background-color: rgb(241, 241, 241);
                color: black;
            }

            .bg-secondary {
                background-color: #ececec;

                /* Not working in GMail */
                box-shadow: #a5a5a5a5 0px 0px 25px;
            }

            .badge {
                background-color: #d2d2d2;
            }

            .card-header {
                background-color: #240e62;
            }
        }

        .content {
            margin: 5px;
            margin-bottom: 15px;
            padding: 15px;
        }

        .card {
            border-radius: 5px;
            width: 100%;
            margin: 15px 0px;
        }

        .card-header {
            border-radius: 5px 5px 0px 0px;
            padding: 5px 10px;
            color: white;
        }

        .card-body {
            padding: 10px;
        }

        .badge {
            border-radius: 2px;
            padding: 3px;
        }

        .m-2 {
            margin: 15px;
        }

        .w-full {
            width: 100%;
        }

        .alert {
            width: 100%;
            font-size: 16px;
            color: #fff;
            font-weight: 500;
            padding: 20px 0px;
            text-align: center;
            border-radius: 5px;
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
    </style>
</head>

<body>
    <div class="content">
        <p style="margin-top: 0;">
            This is the report of your most recent drone pipeline build.
            It finished {{ datetime build.finished "Mon Jan 2 15:04:05 MST 2006" "Local" }}.
        </p>

        <p>
            {{#success build.status}}
             <div class="alert alert-good">
               <span>Successful build #{{ build.number }}</span>
             </div>
           {{else}}
             <div class="alert alert-bad">
               <span>Failed build #{{ build.number }}</span>
             </div>
           {{/success}}
        </p>

        <p>
        <div class="bg-secondary card">
            <div class="card-header">
                The build was based on the following commit:
            </div>
            <div class="card-body">
                <table class="w-full" cellpadding="0" cellspacing="0">
                    <tbody>
                        <tr>
                            <td><strong><a href="{{commit.link}}">{{ truncate commit.sha 8 }}: {{commit.message}}</a></strong></td>
                            <td style="text-align: right;"><small class="badge">{{commit.branch}}</small></td>
                        </tr>
                    </tbody>
                </table>
                <table class="w-full" cellpadding="0" cellspacing="0">
                    <tbody>
                        <tr>
                            <td style="width: 35px"><img class="m-2" src="{{commit.author.avatar}}" style="border-radius: 50%;"
                                    width="40px" height="40px" alt="Avatar of {{commit.author.name}}"></td>
                            <td>
                                <div>
                                    <strong>{{commit.author.name}}</strong><br>
                                    {{commit.author.email}}
                                </div>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
        </p>

        <div class="bg-secondary card">
            <div class="card-header">
                Here are some useful links:
            </div>
            <div class="card-body">
                <table class="w-full" cellpadding="0" cellspacing="0">
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
        </div>
    </div>
</body>

</html>
`