# countdown [![Build Status](https://travis-ci.org/justincampbell/countdown.svg?branch=master)](https://travis-ci.org/justincampbell/countdown)

> Counts down for a duration (like `sleep`, but with progress)

## Installation

1. Download the latest package for your platform from the [Releases page](https://github.com/justincampbell/countdown/releases/latest).
2. Untar the package with `tar -zxvf cachout*.tar.gz`.
3. Move the extracted `countdown` file to a directory in your `$PATH` (for most systems, this will be `/usr/local/bin/`).

Or, if you have a [Go development environment](https://golang.org/doc/install):

```
go get github.com/justincampbell/countdown
```

## Usage

```
countdown 5m
```

The duration is parsed with [Go's time.ParseDuration](https://golang.org/pkg/time/#ParseDuration).

Example durations are `30s`, `5m`, `2h`, or `1h30m`. Omitting the time unit will default to seconds.
