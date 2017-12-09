FROM golang:1.8

COPY . /go/src/github.com/HinanawiTenshi/agenda
WORKDIR /go/src/github.com/HinanawiTenshi/agenda/cli

RUN go-wrapper download
RUN go build -o /go/bin/agendalocal .

WORKDIR /go/src/github.com/HinanawiTenshi/agenda/service

RUN go-wrapper download
RUN go build -o /go/bin/agendaserver .

CMD ["agendaserver"]

EXPOSE 8080
