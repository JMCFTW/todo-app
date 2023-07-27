# Start by building the application.
FROM golang:1.18 as build

WORKDIR /go/src/app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build cmd/main.go

# Now copy it into our base image.
FROM gcr.io/distroless/static-debian11
COPY --from=build /go/src/app/main /

ENV GIN_MODE=release
ENV PORT=8080

ENTRYPOINT ["/main"]
