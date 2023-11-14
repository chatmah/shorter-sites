BINARY_NAME=short-site

migration-up:
	migrate -path database/migration/ -database "postgresql://postgres:mysecretpassword@localhost:5433/postgres?sslmode=disable" -verbose up

migration-down:
	migrate -path database/migration/ -database "postgresql://postgres:mysecretpassword@localhost:5433/postgres?sslmode=disable" -verbose down

migration-fix: 
	migrate -path database/migration/ -database "postgresql://postgres:mysecretpassword@localhost:5433/postgres?sslmode=disable" force VERSION

build:
	GOARCH=amd64 GOOS=darwin go build -o ./bin/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ./bin/${BINARY_NAME}-windows main.go

run: build
	./bin/${BINARY_NAME}-darwin

clean:
	go clean
	rm ./bin/${BINARY_NAME}-darwin
	rm ./bin/${BINARY_NAME}-linux
	rm ./bin/${BINARY_NAME}-windows
