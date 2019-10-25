FROM golang

WORKDIR /go/src/app

RUN go get github.com/mattn/go-sqlite3

COPY . /go/src/app

RUN go build -o server .

CMD ["./server"]

