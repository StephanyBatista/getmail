FROM golang:latest

WORKDIR /app

COPY . /app
# We specify that we now wish to execute 
# any further commands inside our /app
# directory

# we run go build to compile the binary
# executable of our Go program
RUN go mod download
# Our start command which kicks off
# our newly created binary executable


ENTRYPOINT go run main.go
 
EXPOSE 80