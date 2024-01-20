### Go api with MySQL

`docker compose run --rm db--migration`

`docker compose up api`

or (if you want to pass env variable)

`docker compose run --rm -p 4000:4000 -e REQUEST_TIMEOUT_ENABLED=0 api`

`curl -s -X POST -H 'Content-Type: application/json' -d '{"name": "Bryan Cranston"}' http://localhost:4000/v1/actors | jq`

`curl -s http://localhost:4000/v1/actors/1 | jq`

`curl -s http://localhost:4000/v1/actors | jq`
