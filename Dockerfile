FROM golang:1.14.2-alpine3.11 as builder
run apk add git
ENV GO111MODULE=on
WORKDIR /app
ADD go.mod .
ADD go.sum .
RUN go mod download
ADD cmd/api_server/start.go .
RUN go build -o test-backend -tags musl start.go

FROM alpine
COPY --from=builder /app/test-backend /usr/bin/
ENTRYPOINT ["test-backend"]
