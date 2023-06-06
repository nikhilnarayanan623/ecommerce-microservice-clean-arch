FROM golang:1.20.4-alpine3.18 AS build-stage
WORKDIR /home/app
COPY ./ /home/app/
RUN cd /home/app && go mod tidy
RUN go build -o ./api ./cmd

# FROM gcr.io/distroless/static-debian11
FROM alpine
COPY --from=build-stage /home/app/api /api
CMD [ "/api" ]