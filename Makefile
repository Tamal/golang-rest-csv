run:
	nodemon --exec go run ./cmd/emp/ --signal SIGTERM --ext go

build:
	go build -o emp_service ./cmd/emp