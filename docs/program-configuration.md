# Program Configuration

Queued uses external YAML files to define the programs and their process behaviors. These files are referenced via the `include` directive in your main configuration, typically using a glob path like:

```yaml
include: "programs/*.yml"
```

Each program configuration defines how a specific command should run, how many instances to spawn, startup behavior, environment variables, logging, and more.

## Example Configuration


```arduino
arduino:
  group: "web"
  command: "/usr/bin/my-server --port=8080"
  autoStart: true
  autoRestart: true
  startSecs: 5
  slowStart: 2
  numProcs: 3
  user: "appuser"
  stdout: "/var/log/my-server/stdout.log"
  stderr: "/var/log/my-server/stderr.log"
  env:
    - "ENV=production"
    - "DEBUG=false"
```

---

## Program Fields

|Field|Type|Description|
|---|---|---|
|`group`|string|Group label for this program|
|`command`|string|The full command to run|
|`autoStart`|boolean|Start the program automatically on server start|
|`autoRestart`|boolean|Restart the process automatically if it stops|
|`startSecs`|int|Seconds to wait before considering the process successfully started|
|`slowStart`|int|Delay (in seconds) between starting multiple processes|
|`numProcs`|int|Number of process instances to spawn|
|`user`|string|Run the process as this user|
|`stdout`|string|File path to write standard output logs|
|`stderr`|string|File path to write error output logs|
|`env`|[]string|List of environment variables (`KEY=value`)|

---

## Process Metadata (Read-Only)

Queued manages runtime information per process, available via the API or CLI. These fields are not meant to be configured manually:

|Field|Type|Description|
|---|---|---|
|`lastStart`|timestamp|Last time the process was started|
|`processName`|string|Full process name (`<group>:<index>`)|
|`programName`|string|Parent program name|
|`processIndex`|int|Index of the process in the group|
|`pid`|int?|Current PID of the running process|
|`status`|string|Current status: `stopped`, `running`, etc.|

> 📝 These fields are managed internally and typically returned by the API or CLI when querying process status.

---

You can create multiple YAML files for different program groups (e.g., workers, jobs, services) and organize them inside the directory you specify in `include`.

Example structure:

```
config.yml
programs/
├── web-server.yml
├── worker.yml
└── scheduler.yml
```
