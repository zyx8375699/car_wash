FROM golang:1.9.1

MAINTAINER 630272339@qq.com

ADD . /go/src/yuxuan/car-wash

RUN go install yuxuan/car-wash/cmd/car_wash_manager

ENTRYPOINT ["entrypoint.sh"]