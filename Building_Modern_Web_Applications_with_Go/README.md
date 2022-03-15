# Building Modern Web Applications with Go (Golang)

A collection of examples I worked on in Golang based on the course `Building Modern Web Applications with Go (Golang)` on Udemy.

## Helpful Commands

Run Go program `go run main.go [extra filenames]`

Build Go program `go build main.go`
Run built Go program `./main`

Initialize a package `go mod init <package_name>`
Add missing and remove unused modules `go mod tidy`

Run tests `go test -v`
View test coverage in terminal `go test -cover`
View test coverage `go test -coverprofile=coverage.out && go tool cover -html=coverage.out`