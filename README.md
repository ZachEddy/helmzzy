# Helmzzy

Helmzzy is a simple command-line tool that makes it easier to manage Helm releases by providing fuzzy completion for frequently-used `helm` commands.

## Usage ⛵

### Rollback Releases

Revert to an earlier release _nanoseconds_ faster!

![https://asciinema.org/a/215012](https://i.imgur.com/h6nbTkv.gif)

### Delete Releases

Remove `whispering-frog` or `belching-goose` in record time!

![https://asciinema.org/a/215017](https://imgur.com/Lt1gqQx.gif)

## Prerequisites

You'll need the following tools before you set sail.

- [Go](https://golang.org/doc/install#download)
- [fzf](https://github.com/junegunn/fzf#installation)

## Installing

Use `go get` to install the `helmzzy` binary locally.

```sh
$ go get github.com/ZachEddy/helmzzy
```

Next, verify that `helmzzy` was set up correctly installed.

```sh
$ helmzzy
```