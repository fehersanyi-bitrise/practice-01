FROM golang:latest

WORKDIR /go/src/
COPY . /go/src
EXPOSE 3030
RUN go get "github.com/gorilla/mux"
CMD ["go build", "."]
ENTRYPOINT ["/go/src/practice-01"]