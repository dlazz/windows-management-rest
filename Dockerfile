FROM golang:1.20.4

ARG VERSION

RUN mkdir  ./build \
    && mkdir code 

RUN go install github.com/swaggo/swag/cmd/swag@v1.8.10

WORKDIR code 
COPY . .

RUN swag init -o ./webserver/docs \
    && GOOS=windows go build -o ../build/wmr.exe -ldflags="-X 'main.Version=v.${VERSION}'"  \
    && cp example/config-example.json ../build/config.json