FROM golang:1.18 as build
ARG build_version
ARG build_date


RUN apt-get update && apt-get install -y \
    vim \
    procps \
    iputils-ping \
    netcat


#COPY logs/session.log /

COPY . /go/src

WORKDIR /go/src/

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go app

RUN CGO_ENABLED=0 go build -o main -ldflags "-X main.BuildTag=$build_version -X main.BuildDate=$build_date" .

#
# Deploy
#
FROM golang:1.18

RUN apt-get update && apt-get install -y \
    vim \
    procps \
    iputils-ping \
    netcat


WORKDIR /
#
COPY --from=build /go/src .

# Expose port 8083 to the outside world
EXPOSE 8080


ENTRYPOINT ["./main"]