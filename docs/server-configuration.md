# Server Configuration

Queued requires a YAML configuration file to define how the server should behave. This includes HTTP API settings, logging, notifications, and more.

To start the server with a specific configuration file, use the following command:
`queued server -c /path/to/config.yml`

## Configuration Structure

Below is the top-level structure of the configuration file:
```server config
queued:
  api:
    httpListen: ":3200"
    cors: "*"
    authEnabled: true
    authAdmin: "admin:password"
    authReadOnly: "viewer:password"
    accessToken: "your-token-here"

  notification:
    level: "info"
    telegram:
      botToken: "your-bot-token"
      chatId: "your-chat-id"
    lark:
      webhookUrl: "https://open.larksuite.com/..."
      secret: "your-lark-secret"
    slack:
      webhookUrl: "https://hooks.slack.com/..."
      channel: "#alerts"
      user: "QueuedBot"

  log:
    level: "info"
    location: "/var/log/queued.log"

  include: "programs/*.yml"
```
## `api` Section

Configure the HTTP server and access settings:

|Field|Type|Description|
|---|---|---|
|`httpListen`|string|Address and port to listen on (e.g., `:3200`)|
|`cors`|string|CORS policy (e.g., `"*"` for open access)|
|`authEnabled`|boolean|Enable basic authentication for the API|
|`authAdmin`|string|Admin credentials in `username:password` format|
|`authReadOnly`|string|Read-only credentials in `username:password` format|
|`accessToken`|string|Optional bearer token for additional security|

## `notification` Section

Set up log level and third-party notification services (Telegram, Lark, Slack):

|Field|Type|Description|
|---|---|---|
|`level`|string|Log level: `debug`, `info`, `warning`, `error`, `panic`, `fatal`|

Each service (Telegram, Lark, Slack) has its own subfields:

### Telegram

- `botToken`: Your Telegram bot token.
- `chatId`: Chat or group ID to receive messages.

### Lark

- `webhookUrl`: Lark webhook URL.
- `secret`: Optional signing secret for verification.

### Slack

- `webhookUrl`: Slack webhook URL.
- `channel`: Target channel (e.g., `#alerts`).
- `user`: Display name of the sender bot.

## `log` Section

Control how Queued logs information:

|Field|Type|Description|
|---|---|---|
|`level`|string|Log level: `debug`, `info`, `warning`, `error`, `panic`, `fatal`|
|`location`|string|File path to write logs (e.g., `/var/log/queued.log`)|

## `include` Field

Define an external glob path or file to load program definitions. This helps organize process configs separately.
