build_auth:
	@echo " >> building auth service binary"
	@go build -o authapp cmd/auth/main.go

start_auth:
	@echo " >> starting auth service"
	@./authapp -appname=authapp