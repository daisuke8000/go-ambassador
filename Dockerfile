FROM golang:1.17.1

## Allow Go to retreive the dependencies for the build step

WORKDIR /app/
ADD . /app/

# COPY
COPY go.mod .
COPY go.sum .
COPY entrypoint.sh .

# RUN
RUN chmod +x entrypoint.sh && \
    go mod download && \
    go get github.com/go-delve/delve/cmd/dlv && \
    apt-get update && \
    apt-get -y install netcat && \
    curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

COPY . .

RUN go build -o /app/tmp/main .

CMD ["sh", "entrypoint.sh"]