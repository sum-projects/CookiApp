FROM alpine:latest

RUN mkdir /app

COPY ../auth /app

COPY ../auth/app.env /app

CMD [ "/app/authApp"]