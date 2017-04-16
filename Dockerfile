FROM golang:1.8

RUN apt-get update \
    && apt-get install -y unzip \
    && go get github.com/golang/lint/golint

COPY . /marathon-schema
WORKDIR /marathon-schema
RUN make

ENTRYPOINT ["/marathon-schema/marathon-schema"]

# ENV CGO_ENABLED 0
# ENV GOPATH /go
