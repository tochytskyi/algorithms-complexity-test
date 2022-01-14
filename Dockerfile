# --- Step 1
FROM golang:1.17-alpine AS local

WORKDIR /go/src/app

RUN apk update \
    && apk upgrade --force \
    && apk add --no-cache bash make

COPY ./ ./

RUN go get -d -v ./...
#RUN go install -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
      -ldflags='-w -s -extldflags "-static"' -a \
      -o /artifacts/app .

RUN cp .env.example /artifacts/.env


# --- Step 2
FROM scratch AS bin-unix

WORKDIR /
COPY --from=local /artifacts/ /

CMD ["/app"]

