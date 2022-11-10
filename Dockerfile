
FROM golang:1.16-alpine
RUN mkdir /main
ADD . /main

WORKDIR /main

ADD go.mod ./main
ADD go.sum ./main
RUN go mod download
RUN go get github.com/gofiber/fiber/v2
RUN go get github.com/joho/godotenv
RUN go get gorm.io/gorm
RUN go get -u github.com/lib/pq
RUN go get -t

ADD *.go ./main

RUN go build -o main .
EXPOSE 3000 5432
CMD [ "/main/main" ]