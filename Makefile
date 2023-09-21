DB_SCHEMA=authentication

protocgen:
	cd proto/ && \
	protoc --go_out=. --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false *.proto -I${GOPATH}/src -I. && \
	cd ..
migrate-up:
	migrate -path "internal/migration" -database "mysql://${DB_CREDENTIALS_USR}:${DB_CREDENTIALS_PSW}@tcp(${DB_HOST}:3306)/${DB_SCHEMA}?parseTime=true" up
migrate-down:
	migrate -path "internal/migration" -database "mysql://${DB_CREDENTIALS_USR}:${DB_CREDENTIALS_PSW}@tcp(${DB_HOST}:3306)/${DB_SCHEMA}?parseTime=true" down

test:
	go test -cover -race ./...

.PHONY: gen test protocgen