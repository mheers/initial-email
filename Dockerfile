FROM golang:1.19-alpine as builder

# Create the working directory.
RUN mkdir /app
WORKDIR /app

# Copy the source code and build the application.
COPY . .
RUN go build -o initial-email

FROM alpine:3.17.1

# Create the working directory.
RUN mkdir /app
WORKDIR /app

# Copy the compiled binary from the builder image.
COPY --from=builder /app/initial-email .

# Run the application.
CMD ["./initial-email"]
