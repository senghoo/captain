from golang
MAINTAINER Senghoo Kim "shkdmb@gmail.com"
ENV CAPTAIN_STATIC="/go/src/github.com/senghoo/captain" CAPTAIN_WORKSPACE="/workdir/workspace" LOG_PATH="/workdir/logs"
RUN mkdir -p /go/src/github.com/senghoo/captain
ADD . /go/src/github.com/senghoo/captain
RUN go get github.com/senghoo/captain/...
RUN go install github.com/senghoo/captain
RUN mkdir /workdir
WORKDIR /workdir
CMD ["/go/bin/captain", "web"]
