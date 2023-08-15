# build stage
FROM golang:1.20.4-alpine3.18 AS build-stage

WORKDIR /app
# download dependecnies
COPY go.mod go.sum ./
RUN go mod download

# copy the sorce files
COPY pkg pkg/
COPY cmd cmd/

# build executable file
RUN go build -v -o build/api cmd/api/main.go

# final stage
FROM gcr.io/distroless/static-debian11

WORKDIR /app
# copy build file from build stage
COPY --from=build-stage /app/build/api api

CMD ["/app/api"]
