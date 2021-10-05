FROM golang:1.16-alpine AS build

WORKDIR /src/src
COPY *.go .
RUN CGO_ENABLED=0 go build -o /bin/app *.go

FROM alpine:3.14
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]