# Example Of Order application

This is example of order application

# Design

## Microservice Design

This application formed from 4 services and divided by each domain
- Auth service
- Product service
- Logistic service
- Order service
- Promo service

Simple architecture:
1. `Order` is the central application
2. When user want to submit an order, `Order` service will check if user is logged in through `Auth` service.
3. `Order` service will also look into `Product` service is product is exists.
4. When user want to submit the order, `Order` service will make sure `Product` is available and voucher is correct thorugh `Promo` service.
5. After order is done, admin can confirm the `Order` and `Order` will tell `Logistic` for `Shipment`

Image: `/files/simple_architecture.png`

![Simple Architecture](./files/simple_architecture.png?raw=true "Simple Architecture")

## Repo/Code Design 

The design of repo might be sub-optimal because we are separating application by each domain but all dependencies reside
in one repository. Some of the problem:
- Codes in errors is gonna be bloated because it is used to store all the errors. Investing some time needed to address this problem.
- Config and Log is not abstract enough as dependency.

In my opinion, it is better to split the repo to each service. Eeach service can still have multiple binaries but with less strong dependency as now happened. Deployment should also be easier because not doing download for all big services code. Abstraction layer should keep meintained but generic code maybe can be discarded as `common` package.


# Docker

For easier setup and testing, use Docker and `docker-compose.yml` in this repo.

# How To Run

Make sure your Docker is up and running.

Make sure all dependencies is up-to-date by running `go get -u` first

1. Build all application: `make build`.
2. Start applications by running `make run`.

# How To Check If Services Running

use `make check` and if there are 5 PIDs, then all services is running.

# How To Stop

To stop all services, use `make stop`.

# Testing

Test cases is available in `files/testing`