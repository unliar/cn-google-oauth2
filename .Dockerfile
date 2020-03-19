FROM golang:latest as go-build

RUN mkdir /app

WORKDIR /app

COPY . /app

RUN GO111MODULE=on GOPROXY=https://goproxy.io CGO_ENABLED=0 GOOS=linux go build -a -innpm install -g now-clistallsuffix cgo -o main app.go

RUN ls

FROM orvice/go-runtime:lite

WORKDIR /


COPY --from=go-build /app/main /

ENTRYPOINT ["/main"]