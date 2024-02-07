FROM golang:1.18 

RUN mkdir /ACA
ADD . /ACA
WORKDIR /ACA

RUN go build -o ejecutable .

RUN chmod +x /ACA/ejecutable

ENTRYPOINT ["/ACA/ejecutable"]