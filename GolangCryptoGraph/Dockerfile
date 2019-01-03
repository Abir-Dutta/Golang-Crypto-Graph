FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o proj1 .
CMD ["/app/proj1"]
