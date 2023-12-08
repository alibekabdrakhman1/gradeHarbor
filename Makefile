protoc:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative ./pkg/proto/user/user.proto
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative ./pkg/proto/class/class.proto

swagger-auth:
	cd internal/auth && swag init -g ../../cmd/app/auth/main.go

swagger-class:
	cd internal/class && swag init -g ../../cmd/app/class/main.go

swagger-user:
	cd internal/user && swag init -g ../../cmd/app/user/main.go

fix-lint:
	golangci-lint run --fix

compose-up:
	docker-compose up --build -d postgres && docker-compose logs -f
.PHONY: compose-up

compose-down:
	docker-compose down
.PHONY: compose-up