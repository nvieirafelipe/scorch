# Scorch

Backend server for burndown charts apps, written in [go](http://golang.org/)

## Development

First you need to define this env vars:

    GOPATH=/path/to/project/scorch
    PATH=$GOPATH/bin:$PATH

Install dependencies:

    go get github.com/nvieirafelipe/scorch

Run the server with:

    PORT=3000 go run src/github.com/nvieirafelipe/scorch/scorch.go

And then access the API with at `http://localhost:3000`

## API

### Repositories `organizations/:organization_name/repositories`
_Lists all repos by organization._

### Work left VS Time `repositories/:repository_name/workleft-vs-time`
#### params:
     organization = organization_name
_Lists all issues by repo._