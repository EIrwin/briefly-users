FROM golang
 
WORKDIR /go/src/briefly-users
ADD . /go/src/briefly-users
RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux
RUN go install location-tracking
EXPOSE 8181
ENTRYPOINT /go/bin/location-tracking
 
