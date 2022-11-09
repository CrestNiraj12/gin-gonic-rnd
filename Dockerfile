FROM golang:latest

LABEL maintainer="Niraj <crestniraj@gmail.com>"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 3000

RUN go build

# Remove source files after building app
RUN find . -name "*.go" -type f -delete

# Make port available outside container
EXPOSE $PORT

CMD ["./golab-gin-poc"]
