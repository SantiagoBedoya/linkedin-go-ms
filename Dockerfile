FROM golang:1.21-alpine as dependencies

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM dependencies as build
COPY . ./
RUN CGO_ENABLED=0 go build -o /main -ldflags="-w -s" ./cmd/main.go

FROM golang:1.21-alpine
COPY --from=build /main /main
CMD [ "/main" ]