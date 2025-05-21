# Introduction

**Queued** is a lightweight and configurable process manager written in Go. Designed for simplicity and performance, Queued provides powerful features for managing and monitoring processes in both development and production environments.

Whether you're orchestrating background workers, daemons, or CLI-based services, Queued gives you fine-grained control over how processes are grouped, started, and maintained—while exposing a REST API and a dedicated control client for seamless management.

## Key Features

- **Process Grouping**: Organize and manage processes under logical groups.
- **Slow Start**: Stagger process startups to prevent resource spikes or race conditions.
- **Scalable Workers**: Define the number of process instances to run per program.
- **Output Management**: Configure custom directories for `stdout` and `stderr` logs.
- **User Execution**: Run processes under specific user permissions for security and isolation.
- **Auto Start & Restart**: Automatically start or restart processes on failure or boot.
- **Environment Configuration**: Set custom environment variables for each process.
- **REST API**: Easily manage and query processes via a built-in HTTP API.
- **Separate Control App**: A standalone CLI to interact with the Queued server.

Queued is ideal for users who want a lightweight alternative to traditional supervisors like `systemd`, `supervisord`, or `runit`, with a modern API-driven architecture.
