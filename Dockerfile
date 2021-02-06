FROM golang
ADD . /app
WORKDIR /app
RUN go build # This will create a binary file named app
ENTRYPOINT /app/app