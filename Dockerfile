FROM golang:1.14-alpine as builder
RUN mkdir -p /backend
RUN CGO_ENABLE=0
RUN GOOS=linux

ENV GOPATH /go
WORKDIR /backend
ADD go.mod /backend
ADD go.sum /backend
RUN go mod download
ADD . /backend

#RUN GO GET
RUN go build

FROM alpine
COPY --from=builder /backend/backend /app/
RUN chmod +x /app/backend
WORKDIR /app
ENTRYPOINT ["/app/backend"]