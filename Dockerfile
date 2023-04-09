FROM golang:latest

ENV GOPATH=/

COPY ./ go.mod go.sum ./

RUN go mod download && go build -o main ./cmd/main.go

EXPOSE 3002 3001

CMD ["./main", "", "", ""]