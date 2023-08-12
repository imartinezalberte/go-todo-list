FROM golang:alpine3.18

RUN apk add --no-cache --update ca-certificates tzdata git && update-ca-certificates

WORKDIR /go/src/github.com/imartinezalberte/go-todo-list/backend

COPY . .

# I'm using low bandwidth
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags="-w -s" -o /go/bin ./...

CMD ["/go/bin/backend"]
