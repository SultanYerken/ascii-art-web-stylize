FROM golang:alpine AS builder

WORKDIR /app

COPY . ./

RUN go build -o dockermain cmd/web/*.go


FROM alpine

RUN apk add bash 

WORKDIR  /app2

COPY --from=builder /app/dockermain            /app2/
COPY --from=builder /app/templates             /app2/templates/
COPY --from=builder /app/cmd/ascii-art/*.txt   /app2/cmd/ascii-art/

EXPOSE 1800

CMD ["./dockermain"]


