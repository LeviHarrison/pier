FROM golang:1.15

WORKDIR /go/src/github.com/leviharison/pier
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["pier"]
