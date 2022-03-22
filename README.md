# Spocker

A tool that runs locally and allows you to create a mock API endpoint by filling out a small form on a dashboard.

A work in progress.

Written in Svelte and Go.

## Downloads

- [go](https://go.dev/doc/install)
- gin

`go get -u github.com/codegangsta/gin`

For Mac: add `alias gin='[path to go]/go/bin/gin` if `gin -h` yields `command not found: gin`

## Run Locally

The API is working. Run it with this command. There's a postman collection in the repo as well.

```
gin --appPort 5001 --port 5005 -x ./web --immediate
```

## Backlog

- TODO