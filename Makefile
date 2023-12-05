protoc:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative ./pkg/proto/user/user.proto
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative ./pkg/proto/class/class.proto

all: class auth user

class:
	go build -o bin/class github.com/alibekabdrakhman1/gradeHarbor/cmd/app/class

auth:
	go build -o bin/auth github.com/alibekabdrakhman1/gradeHarbor/cmd/app/auth

user:
	go build -o bin/user github.com/alibekabdrakhman1/gradeHarbor/cmd/app/user

