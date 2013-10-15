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

####  example:

    {
        milestones: [
            {
                number: 1,
                title: "1.0",
                url: "https://api.github.com/repos/organization/repository/milestones/1",
                created_at: "2013-09-06T14:08:15Z",
                due_on: "2013-10-03T07:00:00Z",
                work_left_vs_time: [
                    11, 11, 11, 10, 10, 8, 8, 8, 8, 8, 7, 5, 5, 4, 4
                ]
            }
        ]
    }

_Lists work left grouped by day._