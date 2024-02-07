FROM golang:1.18 

RUN mkdir /ACA
ADD . /ACA
WORKDIR /ACA

RUN go build -o ejecutable .

RUN chmod +x /ACA/ejecutable

# Comprobar si el ejecutable se creó correctamente
#RUN ls -l /ACA && /ACA/ejecutable

ENTRYPOINT ["/ACA/ejecutable"]