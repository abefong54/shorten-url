# Project Title

URL Shortner

## Description

A simple URL shortning service similar to TinyURL.

## Getting Started

### Dependencies

- GoLang - go version go1.19.3
- Docker
- github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2
- github.com/go-redis/redis/v8 v8.11.5
- github.com/gofiber/fiber/v2 v2.51.0

### Installing

- Clone repo to your local machine.
- CD into `./shorten-url` directory and run the following command to install dependencies.

```
go mod tidy
```

- Any modifications needed to be made to files/folders

### Executing program

- from the `./shorten-url` directory level run docker containers using the following command

```
docker-compose up -d
```

## Authors

Contributors names and contact info

ex. Abraham Fong (abefong54@gmail.com)

## Version History

- 0.1
  - Initial Release
