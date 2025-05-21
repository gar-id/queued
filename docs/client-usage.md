# Client CLI Usage

The `qctl` CLI tool allows users to interact with the Queued server to manage, inspect, and update process configurations. This includes starting, stopping, restarting processes, fetching logs, checking status, and performing configuration updates.

## General Usage

```bash
qctl <command> [flags]
```

Each command supports a `--config` (`-c`) flag to specify a custom configuration file for the CLI.

## 🔍 Status

Get the status of all managed processes.

```bash
qctl status [flags]
```

**Flags:**

| Flag           | Description             |
| -------------- | ----------------------- |
| `-c, --config` | Select your config file |

---

## 📄 Logs

Retrieve logs (stdout/stderr) for a specific process.

```bash
qctl logs [flags]
```

**Flags:**

|Flag|Description|
|---|---|
|`-c, --config`|Select your config file|
|`-n, --process-name`|Insert the process name|

---

## ▶️ Start

Start processes by group, program, or specific process name.

```bash
qctl start [flags]
```

**Flags:**

|Flag|Description|
|---|---|
|`-c, --config`|Select your config file|
|`-g, --group-name`|Insert group name|
|`-p, --program-name`|Insert program name|
|`-n, --process-name`|Insert process name|

---

## ⏹️ Stop

Stop processes by group, program, or specific process name.

```bash
qctl stop [flags]
```

**Flags:**

|Flag|Description|
|---|---|
|`-c, --config`|Select your config file|
|`-g, --group-name`|Insert group name|
|`-p, --program-name`|Insert program name|
|`-n, --process-name`|Insert process name|

---

## 🔁 Restart

Restart processes by group, program, or specific process name.

```bash
qctl restart [flags]
```

**Flags:**

|Flag|Description|
|---|---|
|`-c, --config`|Select your config file|
|`-g, --group-name`|Insert group name|
|`-p, --program-name`|Insert program name|
|`-n, --process-name`|Insert process name|

---

## 🛠️ Update

Update program configuration without restarting the Queued server. Useful for applying config changes on the fly.

```bash
qctl update [flags]
```

**Flags:**

|Flag|Description|
|---|---|
|`-c, --config`|Select your config file|
