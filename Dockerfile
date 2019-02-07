FROM ubuntu
RUN apt-get install golang-go & \
  export PATH=$PATH:/usr/local/go/bin 
WORKDIR /project/src/
COPY . .
EXPOSE 3030
CMD ["/project/src/practice-01"]
