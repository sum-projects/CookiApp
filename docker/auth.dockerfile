FROM alpine:latest

RUN mkdir /app

COPY ../auth-service /app

COPY ../auth-service/app.env /app

CMD [ "/app/authApp"]