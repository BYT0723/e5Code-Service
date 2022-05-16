run_user_rpc:
	go run ./service/user/rpc/user.go -f ./service/user/rpc/etc/user.yaml

run_user_api:
	go run ./service/user/api/user.go -f ./service/user/api/etc/user-api.yaml

run_project_rpc:
	go run ./service/project/rpc/project.go -f ./service/project/rpc/etc/project.yaml

run_project_api:
	go run ./service/project/api/project.go -f ./service/project/api/etc/project-api.yaml

grpc:
	protoc --go_out=. --go-grpc_out=. ./api/*/*.proto

grpc-user:
	goctl rpc protoc ./api/user/user.proto --go_out=. --go-grpc_out=. --zrpc_out=./service/user/rpc --style=goZero
	
grpc-project:
	goctl rpc protoc ./api/project/project.proto --go_out=. --go-grpc_out=. --zrpc_out=./service/project/rpc --style=goZero
