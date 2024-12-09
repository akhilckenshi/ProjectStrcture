FROM golang:1.23.2-alpine AS build-stage
WORKDIR /exclusive
COPY ./ /exclusive
RUN mkdir -p /exclusive/build
RUN go mod download
RUN go build -v -o /exclusive/build/api ./

FROM gcr.io/distroless/static-debian11
COPY --from=build-stage /exclusive/build/api /
COPY --from=build-stage /exclusive/.env /
EXPOSE 3000
CMD ["/api"]