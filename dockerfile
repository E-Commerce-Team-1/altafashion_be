FROM golang:1.19

##buat folder APP
RUN mkdir /app

##set direktori utama
WORKDIR /app

##copy seluruh file ke app
ADD . .

##buat executeable
RUN go build -o main .

##jalankan executeable
CMD ["./main"]