# Example Of Order application

This is example of order application

# Design

This application formed from 4 services and divided by each domain
- Auth service
- Product service
- Logistic service
- Order service
- Promo service

The design of repo might be sub-optimal because we are separating application by each domain but all dependencies reside
in one repository. Some of the problem:
- Codes in errors is gonna be bloated because it is used to store all the errors. Investing some time needed to address this problem.
- Config and Log is not abstract enough as dependency.

In my opinion, it is better to split the repo to each service. Eeach service can still have multiple binaries but with less strong dependency as now happened. Deployment should also be easier because not doing download for all big services code. Abstraction layer should keep meintained but generic code maybe can be discarded as `common` package.

# Docker

For easier setup and testing, use Docker and `docker-compose.yml` in this repo.

# Authentication Service

## How To Run

To run the app:

- `go build -o authapp cmd/auth/main.go`
- `docker-compose up -d`
- `./txapp`
