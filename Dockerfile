FROM golang:latest
RUN echo "APP_VERSION : ${APP_VERSION}"

WORKDIR /go/src/app
COPY . .

#ENV GOPROXY=https://goproxy.io,direct
ENV TZ=Asia/Jakarta
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o /bin/myapp .

CMD ["/bin/myapp"] 