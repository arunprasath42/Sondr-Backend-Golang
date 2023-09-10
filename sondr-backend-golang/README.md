# Sondr-Golang Backend API's
API's for the Sondr-Golang Backend

![logo](https://github.com/slack-go/slack/blob/master/logo.png?raw=true)
# Table of Contents
1. [Project Structure](#project-structure)
2. [Structure overview](#structure-overview)


# Project structure
This project follows MVC pattern, except the View. The source code related the web application is present inside `src` folder and follows the folder structure of a Java Spring based application.

# Structure-
        .
        |──config/
        |  ├──config.go
        ├──migration/
        |  ├──migration.go
        ├──route.go
        |  ├──routes.go
        ├──src/
        |  ├──controllers
        |  ├──models
        |  ├──repository
        |  ├──service
        ├──test/
        |  ├──service_test.go
        ├──utils/
        |  ├──constant
        |  ├──database
        |  ├──logging
        |  ├──middleware
        |  ├──response
        |  ├──validator
        └──app.yaml
        └──main.go


# Build and deploy

## Download dependencies

To download dependencies, run
```
go mod download
go build
go run main.go
```
```