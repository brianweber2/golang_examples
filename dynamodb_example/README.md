# Repository Pattern Example using DynamoDB

Uses the following packages:

- AWS SDK V2: https://github.com/aws/aws-sdk-go-v2

## Helpful Commands

Run Go program `go run cmd/*.go`

Build Go program `go build cmd/*.go`
Run built Go program `./main`

Initialize a package `go mod init <package_name>`
Add missing and remove unused modules `go mod tidy`

Before runnig tests: `export GITHUB_PROJECT=brianweber2/golang_examples/dynamodb_example`
Run tests `go test -race -timeout=180s -coverpkg=github.com/${GITHUB_PROJECT}/... -coverprofile=coverage.out github.com/${GITHUB_PROJECT}/...`
View test coverage in browser `go tool cover -html=coverage.out`