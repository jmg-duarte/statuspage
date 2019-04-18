
While not all the API's were covered, I chose to cover those whose provider is
the same as `https://status.bitbucket.org`.

The provider for these status pages is [Statuspage](https://www.statuspage.io/) 
and offers a REST API for each page, as documented 
[here](https://help.statuspage.io/knowledge_base/topics/api-information?from_search=true)
under "Status API".

In order to accomodate for such the tool is REST-oriented, meaning it will only
work with REST APIs, furthermore it will only work with the Statuspage API.

The code itself can be reused to create a REST client for Statuspage in Go.

The other service on the list using the same provider is GitHub, however the
tool is flexible enough to work with other clients such as Canary and Reddit.

Each time the program is run it loads the configuration and the local storage 
(defined in the configuration file).

## Configuration

In order to configure the program before running a JSON file must exist in 
`$HOME/.statuspage` or its path can be defined with the flag `--config`.

An example `config.json` file would be:

```json
{
  "services": [
    {
      "id": "github",
      "name": "GitHub",
      "endpoint": "https://kctbh9vrtdwd.statuspage.io/api/v2/"
    },
    {
      "id": "bitbucket",
      "name": "BitBucket",
      "endpoint": "https://bqlf8qjztdtr.statuspage.io/api/v2/"
    },
    {
      "id": "reddit",
      "name": "Reddit",
      "endpoint": "https://2kbc0d48tv3j.statuspage.io/api/v2/"
    },
    {
      "id": "canary",
      "name": "Canary",
      "endpoint": "https://jfhnkhllbpdl.statuspage.io/api/v2/"
    }
  ],
  "output": "./out.json",
  "backup_location": "./backup.json"
}
```

This defines four services, each with an ID, Name and Endpoint for the 
Statuspage REST API.

> Note that in order to find out the endpoints for a Statuspage page you
> simply need to put `/api` in front of the link

## `poll` & `fetch`

While implementing the `poll` and `fetch` commands, I took the liberty to add
the `--brief` flag which, when set to true, shows only the overall status of the 
service.

In order to implement the bonus filters `--only` and `--exclude` I used a set, 
which in Go can be expressed with a map whose value is a 0-byte struct.
The `--only` flag takes precendence over `--exclude`.

This solution was general enough to be encapsulated in its own function and 
"shared" by both `poll` and `fetch`.
Furthermore the `fetch` command is nothing more than a `poll` command in an
endless loop (until the user presses `Ctrl-C`).

## `history`

The history command shows the history contained in local storage, it is able
of filtering using the flag `--only` and `--exclude` just like the commands 
`poll` and `fetch`.

## `backup`

The backup command is able to output JSON and CSV files, the default being JSON
and the CSV being enabled with `--format=csv`.

In order to be flexible enough to handle different services and components
the CSV prints a line identifying the service, then the field IDs and finally
the values.

Such as:

```csv
[bitbucket]
time_utc,Git LFS,Git via HTTPS,Marketplace Apps,Pipelines,Source downloads,Website,Webhooks,API,AWS CodeDeploy App,Atlassian account signup and login,Email delivery,Mercurial via HTTPS,SSH
2019-04-17 12:53:13.951220913 +0000 UTC,operational,operational,operational,operational,operational,operational,operational,operational,operational,operational,operational,operational,operational
```

## `services`

The services command does as stated, showing the available services and 
respective endpoints.

## `help` 

The help command is enabled in an UNIX style, by calling the program with the 
flag `-h` or with the command `help` however this requires a valid configuration
file.

## Compiling

Due to the way Go works, ideally you need to put this project in the following
directory in order to compile it.

```
~/go/src/github.com/jmg-duarte/statuspage
```

However there is already an ELF 64-bit binary `statuspage` ready to use!