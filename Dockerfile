FROM golang:1.19

RUN mkdir /app

ADD . /app

WORKDIR /app
COPY . .
RUN go get -d -v ./...

RUN go build -o products .

EXPOSE 9003

CMD ["./products"]