FROM golang

WORKDIR /go/src/app

COPY . /go/src/app

RUN cd /go/src/app
RUN go get
RUN go build -o server
RUN pwd && ls

CMD ["sh", "/go/src/app/server"]
