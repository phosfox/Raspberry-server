# Raspberry Pi Server?

## Running
Clone the repo and run `docker-compose up` this will start the server.
A `GET` request to `localhost:8090` will be responded with a json-encoded list of added ip addresses. The list is ordered by newly added ip first.
A `POST` request with urlencoded data or json data to `localhost:8090/ip` will add a new ip to the db.

## Example
GET: `curl localhost:8090`

POST: `curl -d "ip=1.1.1.1&name=test123" -X POST localhost:8090/ip`
POST: `curl -d '{"ip":"4.4.4.4", "name":"test321"}' -H "Content-Type: application/json" -X POST http://localhost:8090/ip`
