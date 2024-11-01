build_flags = 

source = ./main.go

bin_name = template
dest = "~/.local/share/TemplateFilesManager/$(bin_name)"



build:
	$(build_flags) go build -o $(bin_name) $(source)

compile:
	$(build_flags) go build -o $(dest) $(source)
