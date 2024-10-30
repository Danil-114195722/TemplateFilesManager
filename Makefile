source = ./main.go

bin_name = "template"
dest = "~/.local/share/TemplateFilesManager/$(bin_name)"


build:
	go build -o $(bin_name) $(source)

compile:
	go build -o $(dest) $(source)
