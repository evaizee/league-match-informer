FROM golang:1.15.0-alpine3.12

WORKDIR /app
COPY /src /app
RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main