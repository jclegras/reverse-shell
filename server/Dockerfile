FROM golang:1.14.3-alpine AS build

WORKDIR /src/
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/server
CMD ["/bin/server"]
