must_build:
	@go get -u
	make build

build:
	make build_auth
	make build_product
	make build_promo
	make build_logistic
	make build_order

run:
	@docker-compose up -d
	@echo " >> waiting postgresql to be ready"
	@sleep 3
	@TXAPPNAME=authapp  ./authapp -log_level=debug & > authapp.out
	@TXAPPNAME=productapp ./productapp -appname=productapp -log_level=debug & > productapp.out
	@TXAPPNAME=promoapp ./promoapp -log_level=debug & > promoapp.out
	@TXAPPNAME=logisticapp ./logisticapp -log_level=debug & > logisticapp.out
	@TXAPPNAME=orderapp ./orderapp -log_level=debug & > orderapp.out
	@echo " >> all services is running"

stop:
	@pkill authapp
	@pkill productapp
	@pkill promoapp
	@pkill logisticapp
	@pkill orderapp
	@docker-compose down
	@echo " >> all services stop"

check:
	@pgrep authapp
	@pgrep productapp
	@pgrep promoapp
	@pgrep logisticapp
	@pgrep orderapp

build_auth:
	@echo " >> building auth service binary"
	@go build -o authapp cmd/auth/*.go

start_auth:
	@echo " >> starting auth service"
	@TXAPPNAME=authapp  ./authapp -log_level=debug

build_product:
	@echo " >> building product service binary"
	@go build -o productapp cmd/product/*.go

start_product:
	@echo " >> starting product service"
	@TXAPPNAME=productapp ./productapp -appname=productapp -log_level=debug

build_promo:
	@echo " >> building promo service binary"
	@go build -o promoapp cmd/promo/*.go

start_promo:
	@echo " >> starting promo service"
	@TXAPPNAME=promoapp ./promoapp -log_level=debug

build_logistic:
	@echo " >> building logistic service binary"
	@go build -o logisticapp cmd/logistic/*.go

start_logistic:
	@echo " >> starting logistic service"
	@TXAPPNAME=logisticapp ./logisticapp -log_level=debug

build_order:
	@echo " >> building order service binary"
	@go build -o orderapp cmd/order/*.go

start_order:
	@echo " >> starting order service"
	@TXAPPNAME=orderapp ./orderapp -log_level=debug
