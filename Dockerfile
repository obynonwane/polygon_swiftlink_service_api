
# base go image - for production 
FROM alpine:latest

RUN mkdir /app

COPY serviceApp /app

CMD [ "/app/serviceApp" ]



