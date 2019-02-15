FROM golang AS builder

WORKDIR /go/src/github.com/fehersanyi/practice-01
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o \
agent main.go

FROM alpine:latest

#RUN apk --no-cache add ca-certificates
WORKDIR /root/src/agent
COPY --from=builder /go/src/github.com/fehersanyi/practice-01/agent .
EXPOSE 3000
CMD ["./agent"]