FROM golang:1.23.0-bullseye

WORKDIR /app 

COPY . ./

RUN go build -o ./pricefetcher

EXPOSE 3000

CMD [ "./pricefetcher" ]
