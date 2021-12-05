run_user_rpc:
	go run ./service/user/rpc/user.go -f ./service/user/rpc/etc/user.yaml

run_user_api:
	go run ./service/user/api/user.go -f ./service/user/api/etc/user-api.yaml

run_project_rpc:
	go run ./service/project/rpc/project.go -f ./service/project/rpc/etc/project.yaml

run_project_api:
	go run ./service/project/api/project.go -f ./service/project/api/etc/project-api.yaml
