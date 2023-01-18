# metalbeat

[![Build Status](https://github.com/devops-metalflow/metalbeat/workflows/ci/badge.svg?branch=main&event=push)](https://github.com/devops-metalflow/metalbeat/actions?query=workflow%3Aci)
[![codecov](https://codecov.io/gh/devops-metalflow/metalbeat/branch/main/graph/badge.svg?token=El8oiyaIsD)](https://codecov.io/gh/devops-metalflow/metalbeat)
[![Go Report Card](https://goreportcard.com/badge/github.com/devops-metalflow/metalbeat)](https://goreportcard.com/report/github.com/devops-metalflow/metalbeat)
[![License](https://img.shields.io/github/license/devops-metalflow/metalbeat.svg)](https://github.com/devops-metalflow/metalbeat/blob/main/LICENSE)
[![Tag](https://img.shields.io/github/tag/devops-metalflow/metalbeat.svg)](https://github.com/devops-metalflow/metalbeat/tags)



> [English](README.md) | 中文



## 介绍

*metalbeat* 作为 [metalflow](https://github.com/devops-metalflow/metalflow) 服务代理，用于向服务注册中心 `consul` 进行服务注册。



## 前提

- Go >= 1.18.0



## 运行

```bash
make build
./metalbeat --config-file=config.yml
```



## 用法

```
usage: metalbeat [<flags>]

MetalBeat

Flags:
  --help                     Show context-sensitive help (also try --help-long and --help-man).
  --version                  Show application version.
  --config-file=CONFIG-FILE  Config file (.yml)
```



## 配置

*metalbeat* 相关配置参数见 [conf](https://github.com/devops-metalflow/metalbeat/blob/main/initialize/conf).

配置文件示例见 [config.yml](https://github.com/devops-metalflow/metalbeat/blob/main/initialize/conf/config.yml).



## 协议

本项目协议声明见 [here](LICENSE).



## 引用

- [consul](https://github.com/hashicorp/consul): A distributed, highly available, and data center aware solution to connect and configure applications across dynamic, distributed infrastructure.
- [lumberjack](https://github.com/natefinch/lumberjack):  A log rolling package for Go.
- [viper](https://github.com/spf13/viper): Go configuration with fangs.
- [zap](https://github.com/uber-go/zap): Blazing fast, structured, leveled logging in Go.
