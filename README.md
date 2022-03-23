# Spocker

A tool that makes creating mock API endpoints as simple as filling out a small form on a dashboard.

A work in progress. Built with Svelte and Go.

## Downloads

- [go](https://go.dev/doc/install)
- [gin](https://github.com/gin-gonic/gin#installation)
- [MongoDB](https://docs.mongodb.com/manual/installation/)

## Run Locally

Make sure your local instance of MongoDB is runnning.

Run the API:

```
gin --appPort 5001 --port 5005 -x ./web --immediate
```

Run the frontend:

```
npm run dev
```

## Backlog

Priority

- MongoDB instead of JSON file (lol)
- Error handling in go
- Create via frontend
- Delete via frontend
- Update via API
- Update via frontend
- Package spocker somehow, Docker?

Future

- Authentication for mock endpoints
- Wildcard endpoints
- Configurable response headers for mock endpoints
- Expiration for mock endpoints
- Log requests received by mock endpoints