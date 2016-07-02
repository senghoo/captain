from golang
MAINTAINER Senghoo Kim "shkdmb@gmail.com"
RUN go get github.com/tools/godep
ENV CAPTAIN_STATIC="/go/src/github.com/senghoo/captain" CAPTAIN_WORKSPACE="/workdir/workspace" LOG_PATH="/workdir/logs"
RUN mkdir -p /go/src/github.com/senghoo/captain
ADD . /go/src/github.com/senghoo/captain
RUN cd /go/src/github.com/senghoo/captain && godep go install
RUN mkdir /workdir
WORKDIR /workdir
EXPOSE 80
CMD ["/go/bin/captain", "web", "-l", "0.0.0.0:80"]
