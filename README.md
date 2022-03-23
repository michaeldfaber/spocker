# Spocker

A tool that makes creating mock API endpoints as simple as filling out a small form on a dashboard.

A work in progress. Built with Svelte and Go.

## Downloads

- [go](https://go.dev/doc/install)
- gin

`go get -u github.com/codegangsta/gin`

For Mac: add `alias gin='[path to go]/go/bin/gin` to `.bash_profile` if `gin -h` yields `command not found: gin`

## Run Locally

The API is working. Run it with this command. There's a postman collection in the repo as well.

```
gin --appPort 5001 --port 5005 -x ./web --immediate
```

The frontend is also working, but it isn't "connected" to the API yet. Run it with this command.

```
npm run dev
```

## Backlog

Priority

- MongoDB instead of JSON file (lol)
- Create via frontend
- Delete via frontend
- Update via API
- Update via frontend
- Package spocker somehow, Docker?

Future

- Authentication for mock endpoints
- Configurable response headers for mock endpoints
- Expiration for mock endpoints
- Log requests received by mock endpoints