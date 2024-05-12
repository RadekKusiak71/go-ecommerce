build:
	@go build -o ./bin/goEcom ./cmd/main.go

run:build
	@./bin/goEcom

migrateUP:
	migrate -path cmd/migrations -database "postgresql://postgres:secretpsw@localhost:5433/gocommerce?sslmode=disable" -verbose up
migrateDOWN:
	migrate -path cmd/migrations -database "postgresql://postgres:secretpsw@localhost:5433/gocommerce?sslmode=disable" -verbose down