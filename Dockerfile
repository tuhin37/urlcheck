# # STAGE1---------------Build the dedicated q3a server in this stage
FROM golang:1.22.4 as builder


# copy the build q3a dedicated server to /app
WORKDIR /app
COPY go.mod  ./ 


RUN go mod download

COPY . . 
## Build the Go app statically
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o urlcheck main.go

# Stage 2: Create a minimal image to run the app
FROM scratch

# Copy the pre-built binary from the previous stage
COPY --from=builder /app/urlcheck /

# Command to run the executable
ENTRYPOINT ["/urlcheck"]