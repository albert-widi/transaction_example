build_auth:
	@echo " >> building auth service binary"
	@go build -o authapp cmd/auth/*.go

start_auth:
	@echo " >> starting auth service"
	@./authapp -appname=authapp -log_level=debug

build_product:
	@echo " >> building product service binary"
	@go build -o productapp cmd/product/*.go

start_product:
	@echo " >> starting product service"
	@./productapp -appname=productapp -log_level=debug

build_promo:
	@echo " >> building promo service binary"
	@go build -o promoapp cmd/promo/*.go

start_promo:
	@echo " >> starting promo service"
	@./promoapp -appname=promoapp -log_level=debug