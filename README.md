# Hello World

## Dependencies

If `go version` outputs something that is `<1.13`, set the `GO111MODULE` env variable as `on`

```sh
go build
```

## Running

```sh
go run main.go
```

### Dockerized

```sh
docker build . -t hello-world-go
```

```sh
docker run -it -p 8080:8080 --rm --name my-running-hellow-world-app hello-world-go
```
