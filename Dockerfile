FROM golang:1.17

WORKDIR /app

# COPY
COPY go.mod .
COPY go.sum .
COPY entrypoint.sh .

# RUN
RUN chmod +x /app/entrypoint.sh && \
    go mod download && \
    apt-get update && \
    apt-get -y install netcat && \
    curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

COPY . .

CMD ["sh", "entrypoint.sh"]