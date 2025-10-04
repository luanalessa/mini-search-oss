#build

# build
FROM golang:1.22 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/minisearch ./cmd/api

# run
FROM gcr.io/distroless/base-debian12:nonroot
WORKDIR /
ENV HTTP_ADDR=":8080"
EXPOSE 8080
COPY --from=builder /bin/minisearch /bin/minisearch
USER nonroot:nonroot
ENTRYPOINT ["/bin/minisearch"]