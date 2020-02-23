FROM golang
EXPOSE 9500
EXPOSE 9501

COPY . /app
WORKDIR /app

RUN go build -o rzreversescheme .

CMD ["/app/rzreversescheme", "--proto=http"]
