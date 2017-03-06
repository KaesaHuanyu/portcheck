FROM daocloud.io/golang
maintainer james.xiong@daocloud.io
WORKDIR /gopath/app
ENV GOPATH /gopath/app
RUN apt-get update
ADD . /gopath/app/
RUN go install portcheck
RUN rm -fr /gopath/app/src
ENV TZ Asia/Shanghai
EXPOSE 8989
CMD ["/gopath/app/bin/portcheck"]
