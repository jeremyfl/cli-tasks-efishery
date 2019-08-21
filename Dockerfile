FROM golang:alpine

ADD ./ /go/src/app

WORKDIR /go/src/app

ENV PORT=3001

CMD ["go", "run", "app.go"]