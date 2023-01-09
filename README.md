# metalbeat

[![Build Status](https://github.com/devops-metalflow/metalbeat/workflows/ci/badge.svg?branch=main&event=push)](https://github.com/devops-metalflow/metalbeat/actions?query=workflow%3Aci)
[![codecov](https://codecov.io/gh/devops-metalflow/metalbeat/branch/main/graph/badge.svg?token=El8oiyaIsD)](https://codecov.io/gh/devops-metalflow/metalbeat)
[![Go Report Card](https://goreportcard.com/badge/github.com/devops-metalflow/metalbeat)](https://goreportcard.com/report/github.com/devops-metalflow/metalbeat)
[![License](https://img.shields.io/github/license/devops-metalflow/metalbeat.svg)](https://github.com/devops-metalflow/metalbeat/blob/main/LICENSE)
[![Tag](https://img.shields.io/github/tag/devops-metalflow/metalbeat.svg)](https://github.com/devops-metalflow/metalbeat/tags)



## Introduction

*metalbeat* is the agent of [metalflow](https://github.com/devops-metalflow/metalflow) written in Go.



## Prerequisites

- Go >= 1.18.0



## Run

```bash
version=latest make build
./metalbeat --config-file="config.yml" --output-file="output.json"
```



## Usage

```
usage: metalbeat [<flags>]

metalbeat

Flags:
  --help                     Show context-sensitive help (also try --help-long and --help-man).
  --version                  Show application version.
  --config-file=CONFIG-FILE  Config file (.yml)
```



## Settings

*metalbeat* parameters can be set in the directory [conf](https://github.com/devops-metalflow/metalbeat/blob/main/initialize/conf).

An example of configuration in [config.yml](https://github.com/devops-metalflow/metalbeat/blob/main/initialize/conf/config.yml):

```yaml
```



## License

Project License can be found [here](LICENSE).



## Reference
