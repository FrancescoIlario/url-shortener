# url Shortener

This repository follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout) and contains a implementation of a simple url shortener.


## ID Format

IDs are generated using the "crypto/rand", the cryptographically secure random number generator from the go's stdlib.
IDs are then encoded in base58 to be human-readable.
Base58 has been used instead of base64 to avoid special chars like '/+=' that would have otherwise produced invalid URLs.

The length of IDs has been set to 7 chars in order to obtain low probability of collision in case of millions of short URLs.
Indeed, with 7 chars we have 58^7 combinations, i.e. like 2,2×10¹².


## Fast URL translation

To translate a short url to the corresponding long url within 10 ms I've used Redis with persistence as a key-value database.
IDs are stored in Redis as key and long urls as value.


## Metrics

To store the metrics I've used a PostgreSQL database.
Metrics are stored in a simple table where each row represents a request to the service.

The postgres database has been seeded with some dummy data.
You can use the UrlID 'abcdefg' to obtain the results generated from the seeded data.

Whenever a user request for a short url translation, a goroutine is fired to update the metrics database.


## Testing Scripts

I've prepared a list of simple scripts that can be used to test the running application.
You can find them in the folder `scripts`.

- `metrics.sh`: invokes the metrics endpoint; it requires an id to be passed
- `shorten.sh`: invokes the shorten endpoint; it requires an url to be passed
- `rount_trip.sh`: invokes the shorten endpoint and then the process one; it requires an url to be passed


## Execution

Please use the Docker Compose file in the `deploy` directory to execute the application.

```console
docker-compose -f deploy/docker-compose.yml up
```

Then use the scripts in the `scripts` folder or `curl` to test the application

> The default configuration is using the port 8080 on the host machine.
> If you need to use a different port on the host please change L:8 and L:14 of the `deploy/docker-compose.yml` file


## Endpoints

- `POST /shorten/anon`: reads the data in the request body (in the format `{ "Url" : "https://sample.io/" }`) and returns a short url
- `GET  /{id}`: reads the id from the path and process the request redirecting to the long url associated to id
- `GET  /metrics/{id}`: retrieves the metrics associated to the ID provided in the path
