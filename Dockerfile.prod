FROM golang:alpine3.12

COPY . .

RUN go build src/hpa/*.go

ENTRYPOINT [ "./main" ]