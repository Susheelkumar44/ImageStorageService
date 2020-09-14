FROM golang:1.14

LABEL AuthorName="Susheelkumar S" 

WORKDIR $GOPATH/src/ImageStorageService

COPY . . 

RUN go get -d -v ./...

RUN go install -d -v ./...

EXPOSE 8051

CMD ["ImageStorageService"]
