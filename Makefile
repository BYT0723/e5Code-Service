run_user_rpc:
	go run ./service/user/rpc/user.go -f ./service/user/rpc/etc/user.yaml

run_user_api:
	go run ./service/user/api/user.go -f ./service/user/api/etc/user-api.yaml

run_project_rpc:
	go run ./service/project/rpc/project.go -f ./service/project/rpc/etc/project.yaml

run_project_api:
	go run ./service/project/api/project.go -f ./service/project/api/etc/project-api.yaml

build:
	cp ./service/user/rpc/etc/user.yaml ./bin/etc/
	cp ./service/user/api/etc/user-api.yaml ./bin/etc/
	cp ./service/project/rpc/etc/project.yaml ./bin/etc/
	cp ./service/project/api/etc/project-api.yaml ./bin/etc/
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o ./bin/userRpc ./service/user/rpc/user.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o ./bin/userApi ./service/user/api/user.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o ./bin/projectRpc ./service/project/rpc/project.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o ./bin/projectApi ./service/project/api/project.go
