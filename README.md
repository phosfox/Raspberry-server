# Raspberry Pi Server?

## Running
Clone the repo and run `docker-compose up` this will start the server.
A `GET` request to `localhost:8090` will be responded with a json-encoded list of added ip addresses. The list is ordered by recently added ip first.
A `POST` request with urlencoded data to `localhost:8090/ip` will add a new ip to the db.

## Example
GET: `curl localhost:8090`
POST: `curl -d "ip=1.1.1.1" -X POST localhost:8090/ip`