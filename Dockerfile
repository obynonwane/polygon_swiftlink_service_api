
# base go image - for production 
FROM --platform=linux/amd64 alpine:latest

RUN mkdir /app

COPY serviceApp /app

CMD [ "/app/serviceApp" ]



