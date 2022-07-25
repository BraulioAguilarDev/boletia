# Boletia

Pull & Save info from [currencyapi](https://currencyapi.com/) built with [fiber](https://docs.gofiber.io/) & [sqlc](https://sqlc.dev/)

## Quick install

```sh
$ cp sample.env .env

$ cp config/sample.config.yml config.yml

# Updates the new files with your information if you want
```



## Requirements
+ [Docker](https://www.docker.com/)
+ [sqlc](https://docs.sqlc.dev/en/latest/index.html)
+ [Docker compose](https://docs.docker.com/compose/)

## Test code

 ```sh
#  Build docker image "golang/boletia"
 $ make docker

# Create the containers
 $ make dc-up

# Run migrations
 $ make seed
 ```

## Endpoints

#### Currencies

`GET: http://localhost:8080/api/v1/currencies/MXN?finit=2022-07-20T15:15:03&fend=2022-07-26T20:15:00`

Successfull

```json
{
    "success": true,
    "data": [
        {
            "code": "MXN",
            "value": 20.548484,
            "updated": "2022-07-24T23:59:59Z"
        }
    ]
}
```

Errors

```json
{
    "success": false,
    "message": "parsing time \"2022-17-20T15:15:03\": month out of range"
}
```

```json
{
    "success": false,
    "message": "current code not valid"
}
```
