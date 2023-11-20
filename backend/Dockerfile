FROM golang:alpine
WORKDIR /app
COPY . .
RUN go build -o main .
EXPOSE 9988
CMD ["./main"]