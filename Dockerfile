FROM golang:1.21-alpine as buildStage
WORKDIR /build
RUN go install github.com/swaggo/swag/cmd/swag@latest
COPY ./ ./
RUN swag init
RUN go build -ldflags "-s -w" -o /output .

FROM alpine:latest
WORKDIR /
COPY --from=buildStage /output /output
EXPOSE 8080
CMD [ "/output" ]