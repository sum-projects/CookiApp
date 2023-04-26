FROM alpine:latest

RUN mkdir /app

COPY ../gateway-service /app

COPY ../gateway-service/app.env /app

CMD [ "/app/gatewayApp"]