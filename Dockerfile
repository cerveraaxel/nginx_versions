FROM golang:1.19-alpine as base

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY main.go ./
COPY scrapper scrapper

RUN go build -o /nginx-version-checker
RUN apk add binutils
RUN strip /nginx-version-checker

FROM alpine:3.16

COPY --from=base /nginx-version-checker /nginx-version-checker

EXPOSE 8080

ENV GIN_MODE=release
CMD [ "/nginx-version-checker" ]