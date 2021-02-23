FROM alpine:latest

RUN sed -i s/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g /etc/apk/repositories
RUN apk add --no-cache go mongodb-tools

ENV GOPROXY="https://goproxy.io"
# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN chmod +x /app/tools/update.sh
# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 6868

# Run the executable
CMD ["./main"]