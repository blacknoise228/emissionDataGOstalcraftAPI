FROM golang:1.22.5

COPY . /

WORKDIR /

RUN go install

CMD ["stalcraftBot startBot"]
