FROM golang:tip-alpine3.22

COPY . /moseleyultimate/

WORKDIR /moseleyultimate/
RUN go build -o moseleyultimate ./cmd/web/  

EXPOSE 3000 
CMD [ "./moseleyultimate" ]



