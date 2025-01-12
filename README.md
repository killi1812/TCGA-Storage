# TCGA-Storage
## Dependencies
###  [Taskfile](https://taskfile.dev/) 
Used for building the app. Not necessary see [Running the app]() for more details.

Install task file with go

    go install github.com/go-task/task/v3/cmd/task@latest
or with winget.

    winget install Task.Task

___
### Docker, docker-compose
Used for deployment of containers.
___
### .NET 7 
Used for running the [scrapper](https://github.com/killi1812/PPPK-Scrapper)
## Running the app
Create appsettings.json based on [appsettings.example.json](https://github.com/killi1812/TCGA-Storage/blob/master/appsettings.example.json).

To run the app use `task run` which will pull and start the containers, install dependencies and start the app. Alternativly if you have your own containers just use `go get` to install dependencies and `go run .` to run a program

Exit with ctrl + c and run `task stop` to stop containers