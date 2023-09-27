# authentication
This is a pet project that support multiple environments and cover some of the essential parts of backend development in microservice architecture such as API development, database, in-memory data store, containerization, and container orchestration.

## Guide
### Installation
#### Host Machine:
1. Duplicate file `internal/app/config.example.yaml` as `internal/app/config.yaml`
2. Update `config.yaml` file with your own credentials
3. Create new database schema `authentication`
4. Run `go mod vendor` 
5. Run `go run main.go`

#### Docker:
1. Duplicate file `internal/app/config.example.yaml` as `internal/app/config.yaml`
2. Update `config.yaml` file with your own credentials
3. Update `docker-compose.yaml` as well if needed
4. Run `docker-compose up`

#### Kubernetes:
1. Duplicate file `internal/app/config.example.yaml` as `internal/app/config.yaml`
2. Update `config.yaml` with your own credentials
3. Make sure you already logged in to your docker hub or any other registry using `docker login`
4. Run `make docker-build USERNAME="<YOUR_USERNAME>" VERSION="<DESIRED_IMAGE_VERSION>"` to build your image and push it to the registry
5. Update `spec.template.spec.containers[0].image` value to the same value you used on 4th step in `k8s/auth-deployment.yaml` file
6. [Optional] If you change port mapping, make sure you update `k8s/auth-service.yaml` file as well
7. Run `make k8s-deploy`

To clear all data, you only need to run `make k8s-destroy`

## License
[MIT] (https://github.com/ffauzann/authentication/blob/main/LICENSE) 