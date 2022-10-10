run:
	go run cmd/main.go
sqlc-generate:
	sqlc generate -f ./config/sqlc.yaml
migrate-down:
	- migrate -database cockroachdb://root@localhost:26257/authz_test?sslmode=disable -path internal/constants/query/schemas -verbose down
migrate-up:
	- migrate -database cockroachdb://root@localhost:26257/authz_test?sslmode=disable -path internal/constants/query/schemas -verbose up
migrate-create:
	- migrate create -ext sql -dir internal/constants/query/schemas -tz "UTC" $(ARGS)