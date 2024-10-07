FROM golang:1.16

WORKDIR /app
COPY . .
RUN ls -al
RUN ls -al server/

RUN CGO_ENABLED=0 GOOS=linux go build -o /tcp_chat
CMD ["/tcp_chat"]
