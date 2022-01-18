build:
	go generate ./service/course/enum/course.go
	go generate ./service/repository
	go build -o bin/main main.go

run:
	go generate ./service/course/enum/course.go
	go generate ./service/repository
	go run main.go
