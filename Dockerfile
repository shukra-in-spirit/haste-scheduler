# Use the latest golang image as the base
FROM golang:latest as builder

# Set the working directory
WORKDIR /app

# Clone the repository
RUN git clone https://github.com/shukra-in-spirit/haste-scheduler.git .

# Build the Go API server
RUN make build

# Create a new stage for the final image
FROM scratch

# Copy the binary from the builder stage
COPY --from=builder /app/bin/haste /

# Expose the necessary ports
EXPOSE 8080

# Set the command to execute the binary
CMD ["/haste"]