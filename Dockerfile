FROM golang:alpine

WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go ./
COPY misc misc
COPY search search
COPY track track

RUN go build
ENTRYPOINT ["./MetaGrabAPI"]
