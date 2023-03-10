# Clean Architecture with Go

This is a sample project to demonstrate how to
apply [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
) in Go.

> Note: This is my first Go project, I'm just learning Go so don't expect a good code
> quality.

## Getting Started

### Server

```bash
# start the server
make run
# create and get an user
curl --location --request POST 'http://localhost:8080/api/v1/users/hello-world'
curl --location --request GET 'http://localhost:8080/api/v1/users/hello-world'
curl --location --request GET 'http://localhost:8080/api/v1/users'
```

### CLI

```
# build an use cli
make build-cli
./build/gca-cli user add --username hello-world
./build/gca-cli user update --username hello-world --status active
./build/gca-cli user get --username hello-world
```

## Architecture

The dependency goes towards the
center - [The Dependency Rule](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html#the-dependency-rule).

![Architecture Overview](docs/assets/clean-architecture.png)

## License

Copyright 2023 © CSalih. All rights reserved.

Licensed under MIT