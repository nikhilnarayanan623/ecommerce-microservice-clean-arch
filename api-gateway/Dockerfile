# build stage
FROM golang:1.20.4-alpine3.18 AS build-stage

WORKDIR /app
# download the dependencies
COPY go.mod go.sum ./
RUN go mod download
# copy source codes
COPY cmd cmd/
COPY pkg pkg/
# build executable file
RUN go build -v -o  build/api cmd/api/main.go

# final stage distroles light weight linux image
FROM gcr.io/distroless/static-debian11

WORKDIR /app
# copy from file from build stage
COPY --from=build-stage /app/build/api /app/api
CMD [ "/app/api" ]