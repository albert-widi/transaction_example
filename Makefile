build:
	make build_auth
	make build_product
	make build_promo
	make build_logistic

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