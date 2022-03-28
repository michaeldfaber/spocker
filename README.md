# Spocker

A tool that makes creating mock API endpoints as simple as filling out a small form on a dashboard.

A work in progress. Built with Svelte and Go.

## Run with Docker Compose (Recommended)

### Downloads
- [Docker](https://docs.docker.com/get-docker/)

### Commands

Working on it!

## Run Manually

### Downloads

- [npm](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm)
- [go](https://go.dev/doc/install)
- [gin](https://github.com/gin-gonic/gin#installation)
- [MongoDB](https://docs.mongodb.com/manual/installation/)

### Commands

- Make sure your local instance of MongoDB is running
- Run `gin --appPort 5001 --port 5005 --immediate --build . --path main.go --bin spocker` in the root folder of this repository
- Run `npm run dev` in the `web` folder of this repository

## Backlog

Priority

- Run frontend, api, and mongo in docker containers
- docker compose to run all components together

Future

- Improve error handling
- Update via API
- Update via frontend
- Authentication for mock endpoints
- Wildcard endpoints
- Configurable response headers for mock endpoints
- Expiration for mock endpoints
- Log requests received by mock endpoints