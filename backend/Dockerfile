FROM golang AS builder

WORKDIR /go/src/github.com/fehersanyi/practice-01
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o \
iserver main.go run.go fizzbuzz.go handleCount.go APIHandleCreate.go handleError.go command.go APIHandleList.go APIHanadleGetLog.go

FROM alpine:latest

#RUN apk --no-cache add ca-certificates
WORKDIR /root/src/backend
COPY --from=builder /go/src/github.com/fehersanyi/practice-01/iserver .
EXPOSE 3030
CMD ["./iserver"]