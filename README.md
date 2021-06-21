# supermarket-api

Supermarket-API is a service providing functionality for cataloging produce and their prices.

## Getting Started

To start `supermarket-api` you will need to create an `.env` file at `cmd/api/.env`. A convenience example with sane defaults is located at `cmd/api/example.env` and you can create your environment file by running:

```
~$ cp cmd/api/example.env cmd/api/.env
```

### Usage

Supermarket-API has a Makefile with commonly needed commands. To use the Makefile append the command to `make` in your terminal:

```
~$ make run
```

### Testing

To run unit tests for Supermarket-API use the included `make` command:

```
~$ make unit-test
```

To see test coverage:

```
~$ make test-coverage
```

To see a visual report of test coverage:
```
~$ make coverage-report
```

### Docker

A Dockerfile is included to allow running in ECS or GKS. Make commands are included for building and running the containers. `make docker-run` sets local development variables and should not be used for production.

## API Spec
**Method**|**Endpoint**|**Description**|**Request Body**|**Response**
:-----:|:-----|:-----|:-----|:-----
GET|/v1/produce|Return all catalogued produce.| `null`| 200 OK<br>400 Bad Request<br>500 Internal Server Error
POST|/v1/produce|Add produce items to the catalogue.|`[{"code":"string","name":"string","price":{"amount":123,"currency":"USD"}}]`|201 Created<br>400 Bad Request<br>500 Internal Server Error
GET|/v1/produce/{produceCode}|Get the produce item with the given produceCode.|`null`|200 OK<br>400 Bad Request<br>500 Internal Server Error
DELETE|/v1/produce/{produceCode}|Delete the produce item with the given produceCode.|`null`|204 No Content<br>400 Bad Request<br>500 Internal Server Error

