FROM golang
 
WORKDIR /go/src/location-tracking
ADD . /go/src/location-tracking
RUN go get gopkg.in/mgo.v2
RUN go get github.com/streadway/amqp
RUN go get github.com/gorilla/mux
RUN go install location-tracking
EXPOSE 8181
ENTRYPOINT /go/bin/location-tracking
 
