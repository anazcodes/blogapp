FROM golang:1.24.0-alpine3.21 AS build-stage

WORKDIR /app 

COPY  . . 

RUN go mod download

RUN go build -v -o ./build/bin/ ./cmd/blogapp

FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY  --from=build-stage /app/build/bin/ /

ENTRYPOINT [ "./blogapp" ]
CMD ["--port=3000", "--cache-capacity=50"]
