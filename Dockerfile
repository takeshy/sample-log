FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ADD . /work
WORKDIR /work
RUN go build -o sample-log

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /work/sample-log /sample-log
EXPOSE 80
ENTRYPOINT ["/sample-log"]
