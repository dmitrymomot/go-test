FROM dmitrymomot/golang-alpine

RUN echo $GOPATH
RUN mkdir -p $GOPATH/src/go-test/src
WORKDIR $GOPATH/src/go-test/src

COPY ./config /config
COPY ./src .
COPY ./vendor ./vendor

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 80

CMD ["src"]
