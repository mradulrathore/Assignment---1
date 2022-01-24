build:
	go generate ./service/enum/item.go
	go generate ./service/db.go
	go build -o bin/main main.go

run:
	go generate ./service/enum/item.go
	go generate ./service/db.go
	go run main.go