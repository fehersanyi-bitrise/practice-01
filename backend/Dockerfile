FROM golang AS builder

WORKDIR /go/src/github.com/fehersanyi/practice-01
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o iserver main.go run.go fizzbuzz.go handleCount.go APIHandleCreate.go handleError.go command.go

FROM centos:latest

#RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/fehersanyi/practice-01/iserver .
EXPOSE 3030
CMD ["./iserver"]