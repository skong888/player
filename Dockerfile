FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./
COPY tool ./tool
COPY *.csv ./

RUN go build -o /player

EXPOSE 8080

CMD [ "/player" ]