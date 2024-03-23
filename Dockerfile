FROM golang:1.22.0-alpine3.19 AS build
WORKDIR /app
COPY . . 
RUN go build -o main main.go
RUN apk add curl


#run
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/main . 
COPY app.env .
COPY db/migration .db/migration
COPY start.sh .
COPY wait-for.sh .
RUN chmod +x /app/wait-for.sh

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]


