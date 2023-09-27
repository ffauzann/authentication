DB_SCHEMA=authentication

protocgen:
	cd proto/ && \
	protoc --go_out=. --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false *.proto -I${GOPATH}/src -I. && \
	cd ..

mockgen:
	mockery --config mockery.yaml

migrate-up:
	migrate -path "internal/migration" -database "mysql://${DB_CREDENTIALS_USR}:${DB_CREDENTIALS_PSW}@tcp(${DB_HOST}:3306)/${DB_SCHEMA}?parseTime=true" up

migrate-down:
	migrate -path "internal/migration" -database "mysql://${DB_CREDENTIALS_USR}:${DB_CREDENTIALS_PSW}@tcp(${DB_HOST}:3306)/${DB_SCHEMA}?parseTime=true" down

docker-build:
	docker build -t authentication . && \
	docker tag authentication $(USERNAME)/authentication:$(VERSION) && \
	docker push $(USERNAME)/authentication:$(VERSION)

k8s-destroy: # Clear data async
	kubectl delete deployments authentication & \
	kubectl delete services authentication & \
	kubectl delete pvc mysql-persistent-storage-mysql-statefulset-0 & \
	kubectl delete pvc mysql-pvc & \
	kubectl delete secret mysql-secret &\
	kubectl delete statefulset mysql-statefulset &\
	kubectl delete service mysql &\
	kubectl delete deployments redis &\
	kubectl delete service redis

k8s-deploy:
	kubectl apply -f k8s/

test:
	go test -cover -race ./...

.PHONY: gen test protocgen