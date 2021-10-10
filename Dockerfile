FROM golang:1.17.1-alpine3.14 as dev

## RootSetting & cgo_enable=0 by multi-stg environment
ENV ROOT=/go/src/app
ENV CGO_ENABLED 0

WORKDIR ${ROOT}

# COPY
COPY go.mod go.sum entrypoint.sh ./

# RUN
RUN apk upgrade --update && \
    apk --no-cache add git

RUN chmod +x entrypoint.sh && \
    go mod download && \
    go get github.com/go-delve/delve/cmd/dlv && \
    go get github.com/cosmtrek/air

COPY . ${ROOT}

RUN go build -o /app/tmp/main ${ROOT}

CMD ["sh", "entrypoint.sh"]