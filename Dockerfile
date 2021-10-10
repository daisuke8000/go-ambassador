FROM golang:1.17.1-alpine3.14 as dev

## RootSetting & cgo_enable=0 by multi-stg environment
<<<<<<< HEAD
#ENV ROOT=/go/src/app
#ENV CGO_ENABLED 0

WORKDIR /app/
=======
ENV ROOT=/go/src/app
ENV CGO_ENABLED 0

WORKDIR ${ROOT}
>>>>>>> master

# RUN
RUN apk upgrade --update && \
    apk --no-cache add git
# COPY
COPY go.mod go.sum entrypoint.sh ./

<<<<<<< HEAD
=======
# RUN
RUN apk upgrade --update && \
    apk --no-cache add git

>>>>>>> master
RUN chmod +x entrypoint.sh && \
    go mod download && \
    go get github.com/go-delve/delve/cmd/dlv && \
    go get github.com/cosmtrek/air

COPY . ${ROOT}

RUN go build -o /app/tmp/main ${ROOT}

CMD ["sh", "entrypoint.sh"]