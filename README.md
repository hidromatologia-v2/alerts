# alerts

[![Coverage](https://github.com/hidromatologia-v2/alerts/actions/workflows/codecov.yaml/badge.svg)](https://github.com/hidromatologia-v2/alerts/actions/workflows/codecov.yaml)
[![Release](https://github.com/hidromatologia-v2/alerts/actions/workflows/release.yaml/badge.svg)](https://github.com/hidromatologia-v2/alerts/actions/workflows/release.yaml)
[![Tagging](https://github.com/hidromatologia-v2/alerts/actions/workflows/tagging.yaml/badge.svg)](https://github.com/hidromatologia-v2/alerts/actions/workflows/tagging.yaml)
[![Test](https://github.com/hidromatologia-v2/alerts/actions/workflows/testing.yaml/badge.svg)](https://github.com/hidromatologia-v2/alerts/actions/workflows/testing.yaml)
[![codecov](https://codecov.io/gh/hidromatologia-v2/alerts/branch/main/graph/badge.svg?token=KUQFMZ4IR2)](https://codecov.io/gh/hidromatologia-v2/alerts)
[![Go Report Card](https://goreportcard.com/badge/github.com/hidromatologia-v2/alerts)](https://goreportcard.com/report/github.com/hidromatologia-v2/alerts)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/hidromatologia-v2/alerts)

## Documentation

| Document                                                     | Description          |
| ------------------------------------------------------------ | -------------------- |
| [CONTRIBUTING.md](CONTRIBUTING.md)                           | Contribution manual. |
| [CICD.md](https://github.com/hidromatologia-v2/docs/blob/main/CICD.md) | CI/CD documentation. |

## Docker

```shell
docker pull ghcr.io/hidromatologia-v2/alerts:latest
```

### Compose example

```shell
docker compose up -d
```

## Binary

You can use the binary present in [Releases](https://github.com/hidromatologia-v2/alerts/releases/latest). Or compile your own with.

```shell
go install github.com/hidromatologia-v2/alerts@latest
```

## Config

List of environment variables used by the binary. Make sure to setup them as well in your deployments

| Variable                      | Description                                                  | Example                                                      |
| ----------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `MEMPHIS_CONSUMER_STATION`    | Name for the station to **CREATE**/**CONNECT**               | `alerts`                                                     |
| `MEMPHIS_CONSUMER_CONSUMER`   | Alerts consumer name                                         | `alerts-consumer`                                            |
| `MEMPHIS_CONSUMER_HOST`       | Host or IP of the Memphis service                            | `10.10.10.10`                                                |
| `MEMPHIS_CONSUMER_USERNAME`   | Memphis Username                                             | `root`                                                       |
| `MEMPHIS_CONSUMER_PASSWORD`   | Memphis password, if ignored `MEMPHIS_CONSUMER_CONN_TOKEN` will be used | `memphis`                                                    |
| `MEMPHIS_CONSUMER_CONN_TOKEN` | Memphis connection token, if ignored `MEMPHIS_CONSUMER_PASSWORD` will be used | `ABCD`                                                       |
| `MEMPHIS_PRODUCER_STATION`    | Name for the station to **CREATE**/**CONNECT**               | `messages`                                                   |
| `MEMPHIS_PRODUCER_PRODUCER`   | Messages producer name                                       | `messages-producer                                           |
| `MEMPHIS_PRODUCER_HOST`       | Host or IP of the Memphis service                            | `10.10.10.10`                                                |
| `MEMPHIS_PRODUCER_USERNAME`   | Memphis Username                                             | `root`                                                       |
| `MEMPHIS_PRODUCER_PASSWORD`   | Memphis password, if ignored `MEMPHIS_PRODUCER_CONN_TOKEN` will be used | `memphis`                                                    |
| `MEMPHIS_PRODUCER_CONN_TOKEN` | Memphis connection token, if ignored `MEMPHIS_PRODUCER_PASSWORD` will be used | `ABCD`                                                       |
| `POSTGRES_DSN`                | Postgres DSN to be used                                      | `host=127.0.0.1 user=sulcud password=sulcud dbname=sulcud port=5432 sslmode=disable` |

## Coverage

| [![coverage](https://codecov.io/gh/hidromatologia-v2/alerts/branch/main/graphs/sunburst.svg?token=KUQFMZ4IR2)](https://app.codecov.io/gh/hidromatologia-v2/alerts) | [![coverage](https://codecov.io/gh/hidromatologia-v2/alerts/branch/main/graphs/tree.svg?token=KUQFMZ4IR2)](https://app.codecov.io/gh/hidromatologia-v2/alerts) |
| ------------------------------------------------------------ | ------------------------------------------------------------ |