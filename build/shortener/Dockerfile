FROM gcr.io/distroless/base-debian10 as prod
CMD ["/app"]

FROM golang:1.15 as build
WORKDIR /app
RUN mkdir -p bin 

COPY go.mod .
RUN go mod download \
    && go mod verify

COPY . .
RUN go build -o bin/shortener cmd/shortener/main.go

FROM prod 
COPY --from=build /app/bin/shortener /app