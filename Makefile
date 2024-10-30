source = ./main.go
dest = ./go_app


dev:
	go run $(source)

compile:
	go build -o $(dest) $(source)
