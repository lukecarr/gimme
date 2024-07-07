# Gimme

[![Maintainability](https://api.codeclimate.com/v1/badges/846923e672645e210e76/maintainability)](https://codeclimate.com/github/lukecarr/gimme/maintainability)

Gimme is a minimal HTTP API for fetching various random values (e.g. UUIDs, names, colours).

## Requirements

Gimme has no requirements: simply grab a binary for your respective operating system/architecture from the [GitHub releases page](https://github.com/lukecarr/gimme/releases).

### Building from source/development

If you'd like to develop locally or wish to build from source, you'll need an install of [Go](https://go.dev) (check the [`go.mod`](go.mod) file for the version we're using).

Once you have Go installed, you can then run `go run main.go` to launch Gimme directly, or run `go build` (which will produce a `gimme` executable in the working directory).

## API reference

Below you can find reference documentation for the different routes and functionality that Gimme provides.

### GET `/uuid/{version}`

Generates a UUID string for the provided version.

Supported versions are: `v1`, `v4`, `v6`, and `v7`.

By default, only one UUID is generated. This behaviour can be configured by providing a `?n=123` query parameter as part of your request. Supported values for `n` are integers between 1 and 1000.

The response structure (irrespective of `n`) is always an array of strings.

## CS50x final project

This repository also serves as my entry for the final project of CS50x.

#### Video demo

The below video is a short demo I created to highlight the functionality of this repository. This video is a requirement for the CS50x final project submission.

