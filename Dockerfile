FROM golang:latest

ENV GOPATH=/

COPY ./ ./

COPY go.mod .
COPY go.sum .

RUN go mod download


RUN go build -o main ./cmd/main.go


EXPOSE 3002 3001

CMD ["./main", "", "", ""]