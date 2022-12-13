# Spocker

A tool that makes creating mock API endpoints as simple as filling out a small form on a dashboard.

Built with Svelte and Go. Work in progress. 

## Run with Docker Compose

Working on it!

## Run with Docker

#### Downloads
- [Docker](https://docs.docker.com/get-docker/)

Build the frontend image by running the following command in the `web` folder:

```
docker build -t spocker-frontend .
```

Run a container from the image by running the following command:

```
docker run -p 5000:5000 spocker-frontend
```

Build the backend image by running the following command in the root folder:

```
docker build -t spocker-api .
```

Run a container from the image by running the following command:

```
docker run -p 5001:5001 spocker-api
```

## Run without Docker

#### Downloads

- [npm](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm)
- [go](https://go.dev/doc/install)
- [gin](https://github.com/gin-gonic/gin#installation)
- [MongoDB](https://docs.mongodb.com/manual/installation/)

Make sure your local instance of MongoDB is running.

Run the following command in the root folder of this repository:

```
gin --appPort 5001 --port 5005 --immediate --build . --path main.go --bin spocker
```

Run `npm run dev` in the `web` folder of this repository.

## Backlog

Priority

- docker compose to run all components together

Future

- Update
- Expected authentication for mock endpoints
- Logging