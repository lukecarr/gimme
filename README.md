# Gimme

[![Maintainability](https://api.codeclimate.com/v1/badges/846923e672645e210e76/maintainability)](https://codeclimate.com/github/lukecarr/gimme/maintainability)

Gimme is a minimal HTTP API for fetching various random values (e.g. UUIDs, names, colours).

## Installation

### Requirements

Gimme has no requirements!

### Binaries

Simply grab a binary for your respective operating system/architecture from the [GitHub releases page](https://github.com/lukecarr/gimme/releases).

Then, all you need to do is run `./gimme`, and the server will be running on `http://127.0.0.1:5000` (see [Configuration](#configuration) for details on changing the address).

### Docker image

Gimme can also be deployed from its official Docker image. The Docker image is automatically built using GitHub Actions and pushed to this repository's container registry.

```shell
docker run -d p 5000:5000 ghcr.io/lukecarr/gimme:latest
```

## Configuration

Gimme uses environment variables to configure the server's behaviour:

* `ADDR`: used to configure the address that Gimme will bind to. Defaults to `:5000` (port 5000 on all addresses).

## API reference

Below you can find reference documentation for the different routes and functionality that Gimme provides.

### GET `/uuid/{version}`

Generates a UUID string for the provided version.

Supported versions are: `v1`, `v4`, `v6`, and `v7`.

By default, only one UUID is generated. This behaviour can be configured by providing a `?n=123` query parameter as part of your request. Supported values for `n` are integers between 1 and 1000.

The response structure (irrespective of `n`) is always an array of strings.

## Building from source/development

If you'd like to develop locally or wish to build from source, you'll need an install of [Go](https://go.dev) (check the [`go.mod`](go.mod) file for the version we're using).

Once you have Go installed, you can then run `go run main.go` to launch Gimme directly, or run `go build` (which will produce a `gimme` executable in the working directory).

## CS50x final project

This repository also serves as my entry for the final project of [CS50x](https://cs50.harvard.edu/x/2024/).

### Video demo

The below video is a short demo I created to highlight the functionality of this repository. This video is a requirement for the CS50x final project submission.

### Project structure

The main entrypoint for the application is the `main.go` file. In this file, the `main` function is responsible for registering the different API route handlers and then launching the HTTP server.

The different "modules" (e.g. UUID, colours, names) are split into separate directories within the `internal/` directory.

In the UUID generation logic (`internal/uuid/generate.go`), I use concurrency (channels and goroutines) to simultaneously generate UUIDs within the route handler function. This is to improve performance when the `n` query parameter provided by the requester is large.
