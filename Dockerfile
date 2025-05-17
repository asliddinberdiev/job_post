ARG APP=job_post

FROM golang:alpine3.16 AS builder
RUN apk --no-cache add make git ca-certificates

WORKDIR /app
COPY . .

ARG APP
RUN make build APP=${APP}

FROM alpine
RUN apk --no-cache add ca-certificates

ARG APP
COPY --from=builder /app/bin/${APP} /${APP}
ENTRYPOINT ["/${APP}"]
