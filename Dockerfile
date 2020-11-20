FROM golang:1.15-buster as builder

# RUN mkdir -p /opt/code/

WORKDIR /opt/code/

ADD ./ /opt/code/

RUN go get
# build for alpine
RUN GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -o bin/hello main.go

FROM alpine

WORKDIR /app

COPY --from=builder /opt/code/bin/hello /app/
RUN ls

CMD ["./hello"]