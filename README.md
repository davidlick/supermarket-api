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

A Dockerfile is included to allow running in ECS or GKE. Make commands are included for building and running the containers. `make docker-run` sets local development variables and should not be used for production.

## API Spec
**Method**|**Endpoint**|**Description**|**Request Body**|**Response**
:-----:|:-----|:-----|:-----|:-----
GET|/v1/produce|Return all catalogued produce.| `null`| 200 OK<br>400 Bad Request<br>500 Internal Server Error
POST|/v1/produce|Add produce items to the catalogue.|`[{"code":"string","name":"string","price":{"amount":123,"currency":"USD"}}]`|201 Created<br>400 Bad Request<br>500 Internal Server Error
GET|/v1/produce/{produceCode}|Get the produce item with the given produceCode.|`null`|200 OK<br>400 Bad Request<br>500 Internal Server Error
DELETE|/v1/produce/{produceCode}|Delete the produce item with the given produceCode.|`null`|204 No Content<br>400 Bad Request<br>500 Internal Server Error

## Load Test

This application was load tested using K6. To run the load test follow the installation documentation for K6 [here](https://k6.io/docs/getting-started/installation/). Once installed, make sure Supermarket-API is running and initiate the load test by running:

```
~$ k6 run loadtests/load-test.js
```

The load test is configured to ramp up to 50 iterations per second over a 4 minute period and then maintain that load for a period of 1 minute. Each iteration adds a new unique produce item, and then queries all items catalogued in the service.

Load test results:
```

          /\      |‾‾| /‾‾/   /‾‾/
     /\  /  \     |  |/  /   /  /
    /  \/    \    |     (   /   ‾‾\
   /          \   |  |\  \ |  (‾)  |
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: loadtests/load-test.js
     output: -

  scenarios: (100.00%) 2 scenarios, 500 max VUs, 5m30s max duration (incl. graceful stop):
           * ramp_to_load: Up to 50.00 iterations/s for 4m0s over 1 stages (maxVUs: 200-250, gracefulStop: 30s)
           * maintain_load: 50.00 iterations/s for 1m0s (maxVUs: 200-250, startTime: 4m0s, gracefulStop: 30s)


running (5m00.1s), 000/400 VUs, 9000 complete and 0 interrupted iterations
ramp_to_load  ✓ [======================================] 200/200 VUs  4m0s  50 iters/s
maintain_load ✓ [======================================] 200/200 VUs  1m0s  50 iters/s

   ✓ AddProduceOK...................: 100.00% ✓ 9000  ✗ 0
   ✓ AddProduceTiming...............: avg=0.706569 min=0.195315 med=0.465269  max=15.900239 p(90)=1.499495  p(95)=1.846937
   ✓ GetProduceOK...................: 100.00% ✓ 9000  ✗ 0
   ✓ GetProduceTiming...............: avg=30.18456 min=0.302081 med=27.981297 max=99.590601 p(90)=57.708492 p(95)=63.63907
     data_received..................: 3.3 GB  11 MB/s
     data_sent......................: 2.6 MB  8.5 kB/s
     http_req_blocked...............: avg=16.31µs  min=1.15µs   med=4.26µs    max=7.21ms    p(90)=15.23µs   p(95)=21.52µs
     http_req_connecting............: avg=8.18µs   min=0s       med=0s        max=7.13ms    p(90)=0s        p(95)=0s
     http_req_duration..............: avg=15.44ms  min=195.31µs med=2.28ms    max=99.59ms   p(90)=49.71ms   p(95)=57.7ms
       { expected_response:true }...: avg=15.44ms  min=195.31µs med=2.28ms    max=99.59ms   p(90)=49.71ms   p(95)=57.7ms
     http_req_failed................: 0.00%   ✓ 0     ✗ 18000
     http_req_receiving.............: avg=260.07µs min=16.67µs  med=129.53µs  max=8.89ms    p(90)=594.14µs  p(95)=883.46µs
     http_req_sending...............: avg=34.61µs  min=5.6µs    med=24.85µs   max=583.87µs  p(90)=81.81µs   p(95)=115.66µs
     http_req_tls_handshaking.......: avg=0s       min=0s       med=0s        max=0s        p(90)=0s        p(95)=0s
     http_req_waiting...............: avg=15.15ms  min=160.47µs med=2.04ms    max=98.86ms   p(90)=48.99ms   p(95)=56.93ms
     http_reqs......................: 18000   59.988732/s
     iteration_duration.............: avg=31.42ms  min=1.51ms   med=29.3ms    max=102.12ms  p(90)=59.07ms   p(95)=65.16ms
     iterations.....................: 9000    29.994366/s
     vus............................: 200     min=200 max=200
     vus_max........................: 400     min=400 max=400
```
