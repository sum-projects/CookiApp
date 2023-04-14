FROM alpine:latest

RUN mkdir /app

COPY ../mail-service /app

COPY ../mail-service/app.env /app

CMD [ "/app/mailApp"]