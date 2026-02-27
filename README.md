# Mirth Channel Exporter

Export [Mirth Connect](https://en.wikipedia.org/wiki/Mirth_Connect) channel
statistics to [Prometheus](https://prometheus.io).

Metrics are retrieved using the Mirth Connect REST API.
This has been tested with Mirth Connect >=4.5.

To run it:

```bash
make run
```

Or, if you just want to build it:

```bash
make build
```

The executable will be placed in `./build/mirth-channel-exporter`.

## Docker setup

You can also build a lightweight Docker image using the provided [Dockerfile](./Dockerfile) by running:

```bash
make container
```

This only works if you have Docker installed on your system, and will build an image `mirth-channel-exporter:latest`.

## Exported Metrics

| Metric                        | Description                             | ConstLabels | VariableLabels |
| ----------------------------- | --------------------------------------- | ----------- | -------------- |
| mirth_up                      | Was the last Mirth CLI query successful | status      |                |
| mirth_messages_received_total | How many messages have been received    | channels    | channel        |
| mirth_messages_filtered_total | How many messages have been filtered    | channels    | channel        |
| mirth_messages_queued         | How many messages are currently queued  | channels    | channel        |
| mirth_messages_sent_total     | How many messages have been sent        | channels    | channel        |
| mirth_messages_errored_total  | How many messages have errored          | channels    | channel        |
| mirth_cpu_usage_pct           | CPU usage percentage                    | system      |                |
| mirth_allocated_memory_bytes  | Allocated memory in bytes               | system      |                |
| mirth_free_memory_bytes       | Free memory in bytes                    | system      |                |
| mirth_disk_total_bytes        | Total disk space in bytes               | system      |                |
| mirth_disk_free_bytes         | Free disk space in bytes                | system      |                |
| mirth_num_db_tasks            | Number of database tasks                | database    |                |

## Flags

To get information about command line arguments:

```bash
./build/mirth-channel-exporter --help
```

| Flag               | Description                        | Default    |
| ------------------ | ---------------------------------- | ---------- |
| log.level          | Logging level                      | `info`     |
| web.listen-address | Address to listen on for telemetry | `:9141`    |
| web.telemetry-path | Path under which to expose metrics | `/metrics` |
| config.file-path   | Optional environment file path     | `None`     |

## Env Variables

Use a .env file in the local folder, /etc/sysconfig/mirth_channel_exporter, or
use the --config.file-path command line flag to provide a path to your
environment file

```env
MIRTH_ENDPOINT=https://mirth-connect.yourcompane.com
MIRTH_USERNAME=admin
MIRTH_PASSWORD=admin
```

## Notice

This exporter is inspired by the [consul_exporter](https://github.com/prometheus/consul_exporter)
and has some common code. Any new code here is Copyright &copy; 2020 TeamZero, Inc. See the included
LICENSE file for terms and conditions.
