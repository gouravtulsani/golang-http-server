FROM golang:1.14.2-alpine3.11 as builder
ENV GO111MODULE=on
WORKDIR /app
ADD go.mod .
ADD go.sum .
RUN go mod download
ADD main.go .
RUN go build -o test-backend -tags musl main.go

FROM scratch
COPY --from=builder /app/test-backend /usr/bin/
ENTRYPOINT ["test-backend"]
