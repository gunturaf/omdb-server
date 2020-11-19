# OMDB Server in Go

An attempt to implement clean architecture for simple OMDB Server
that supports HTTP JSON API and GRPC-Protobuf frontends.

Author: Guntur A. Fauzi

## Prerequisites

- go 1.12+
- mysql server

## Configuration

This app using [spf13/viper](https://github.com/spf13/viper), so you can put the configuration values into environment variable or through a `local-config.yaml` file.

Refer to `config/consts.go` for all supported configuration keys.

## Running the App

run `make`

## Test

run `make test`
