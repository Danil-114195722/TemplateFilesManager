source = ./main.go
dest = ./build/template


compile:
	go build -o $(dest) $(source)
