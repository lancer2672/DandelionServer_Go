FROM golang:1.22.0-alpine3.19 AS build
WORKDIR /app
COPY . . 
RUN go build -o main main.go

#run
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/main . 
COPY app.env .

EXPOSE 8080
CMD [ "/app/main" ]



