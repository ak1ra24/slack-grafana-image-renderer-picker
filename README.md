## Slack Grafana Image Renderer Picker

Pick graph with cli tool from Grafana Image Renderer and post graph image to Slack.

### Dependencies

- Grafana
  - Grafana must be accessible with [API Key](https://grafana.com/docs/grafana/latest/http_api/auth/) or [Auth Proxy Authentication](https://grafana.com/docs/grafana/latest/auth/auth-proxy/#auth-proxy-authentication)
    - If you use Auth Proxy Authentication, the reverse proxy must support client certificate authentication.
- [Grafana Image Renderer](https://grafana.com/grafana/plugins/grafana-image-renderer)
  - Grafana needs installed this plugin.

### Deployment

- Docker image from [Packages](https://github.com/ak1ra24/slack-grafana-image-renderer-picker/packages) [DockerHub](https://hub.docker.com/r/akiranet24/gfslack)
- Get binary from [Releases](https://github.com/ak1ra24/slack-grafana-image-renderer-picker/releases)


### Configuration

#### Basic

You need register an Slack Application for Slash Command and files:write permission token.

cli can be configured as follows:

```text
❯ ./gfslack -h
NAME:
   gfslack - upload grafana graph image to slack

USAGE:
   gfslack [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHOR:
   ak1ra24 <ak1ra24net@gmail.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config value, -c value  Specify the Config file. (default: "config.yaml")
   --name value              Specify the dashboard name.
   --time value, -d value    Specify from time e.g.) 1h
   --tz value                Specify timezone (default: "JST")
   --help, -h                show help (default: false)
   --version, -v             print the version (default: false)

```

Configuration file be specified as follows:

```yaml
slack:
    token: "xoxb-test"
    channel: "#test"
grafana:
    endpoint: "http://localhost:3000"
    apikey: "grafana-apikey"
    use_client_auth: true
    client_auth_p12: "/ssl/key.p12"
dashboards:
   -  name: cpu
      dashboardName: alerts-linux-nodes
      orgId: 1
      panelId: 2
   -  name: disk
      dashboardName: alerts-linux-nodes
      orgId: 1
      panelId: 3
```

`dashboards` specify a graph panel to be upload with cli. You can get the parameters of the graph panel by selecting the panel in Grafana and clicking on the share button.

`name` specifies the alias of a graph. So you can get a graph in Slack like `/graph cpu`.

#### Use Auth Proxy Authentication with Client Certificate

This application needs PKCS12 File (.p12) and password, and you need to enable `use_client_auth` and specify p12 file path on `client_auth_p12` at `config.yaml`.

Run with environment: `CONFIG_FILE=config.yaml CLIENT_AUTH_PASSWORD=p12_password`

#### Use API Key 

Run with environment: `CONFIG_FILE=config.yaml GRAFANA_API_KEY=apikey`

### Usage

```
gfslack -c config.yaml --name cpu -d 1h
```

Example `<from_time_range>`: `15m` `3h` `1d` `1M`