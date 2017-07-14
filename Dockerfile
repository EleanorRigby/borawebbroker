FROM golang:1.8
COPY . /go/src/github.com/eleanorrigby/borawebbroker
WORKDIR /go/src/github.com/eleanorrigby/borawebbroker
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /borawebbroker .

FROM ubuntu:16.04
COPY --from=0 /borawebbroker /borawebbroker

#Add all sorts of 
ADD charts /charts

CMD ["/borawebbroker", "-logtostderr"]