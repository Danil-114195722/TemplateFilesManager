build_flags = CGO_ENABLED=1 CGO_CFLAGS="-D_LARGEFILE64_SOURCE" CGO_LDFLAGS="-D_LARGEFILE64_SOURCE"

source = ./main.go

bin_name = template
dest = "~/.local/share/TemplateFilesManager/$(bin_name)"



build:
	$(build_flags) go build -o $(bin_name) $(source)

compile:
	$(build_flags) go build -o $(dest) $(source)
