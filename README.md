# Go REST API

Simple REST API with Go with minimal dependencies. Perfect starting point for a new project.

## Project

The API is divided into smaller services that are responsible for a specific part of the application. 
Each service is the main package to not be opinionated about the project structure.

Example: 
```
.
users.go      // Service for users
users_test.go // Tests for the users service
```

## Run

To run the project, you need to have Go installed. Then, you can run the following command:

```bash
make run 
// or Docker
```

Also make sure to have the environment variables set from the `config.go` file.
I recommend injecting the variables in runtime. I personally use [direnv](https://direnv.net/) for that.

## Test

```bash
make test
```

## Deploy

Just Docker build and run into your favorite cloud provider.