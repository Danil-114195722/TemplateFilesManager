FROM golang:1.22.4-alpine3.20

# set up workdir
RUN cd /go/src
RUN mkdir -p ./github.com/Danil-114195722/TemplateFilesManager
WORKDIR /go/src/github.com/Danil-114195722/TemplateFilesManager
# set up build dir
RUN mkdir /build

# install dependences
COPY ./go.mod .
COPY ./go.sum .
RUN go mod tidy
RUN go mod download

# copy project files to container
COPY . .

# disable CGO
ENV CGO_ENABLED=0

# compile
CMD ["go", "build", "-o", "/build/template", "/go/src/github.com/Danil-114195722/TemplateFilesManager/main.go"]
