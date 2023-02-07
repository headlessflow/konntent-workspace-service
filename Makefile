compose-restart:
	@sh restart.sh

test-integration:
	go test -tags integration ./internal/.../handler/... -race -coverprofile=coverage_integration.out -coverpkg=./internal/.../handler/... -covermode=atomic -v

test-benchmark:
	go test -tags "benchmark integration" ./internal/.../handler/... -race -coverprofile=coverage_integration.out  -coverpkg=./internal/.../handler/... -covermode=atomic -v

test-unit:
	go test -v ./internal/... -race -covermode=atomic -v

test-all: test-integration test-unit

generate-mocks:
	mockgen -source=./internal/app/workspace/service.go -destination=./internal/app/workspace/mocks/service_mock.go -package=mocks
	mockgen -source=./internal/app/workspace/repository.go -destination=./internal/app/workspace/mocks/repository_mock.go -package=mocks
	mockgen -source=./internal/app/orchestration/workspace.go -destination=./internal/app/orchestration/mocks/orchestration_mock.go -package=mocks
	mockgen -source=./internal/app/handler/workspace.go -destination=./internal/app/handler/mocks/handler_mock.go -package=mocks
	mockgen -source=./pkg/pg/instance.go -destination=./pkg/pg/mocks/instance_mock.go -package=mocks
	mockgen -source=./pkg/nrclient/nrclient.go -destination=./pkg/nrclient/mocks/newrelic_mock.go -package=mocks

lint:
	goimports -w . && gofmt -w .
