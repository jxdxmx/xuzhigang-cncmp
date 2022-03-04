FROM golang:1.17-alpine AS build
#COPY module3 /httpserver
WORKDIR /opt/
COPY . .
#ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN go build -o httpserver homework.go
#RUN GOOS=linux GOARCH=amd64 go build -o httpserver homework.go
#RUN go mod vendor && go build -mod=vendor -o /opt/httpserver homework.go
#RUN GOOS=linux go build -o httpserver homework.go
#RUN GOOS=linux go build -installsuffix cgo -o httpserver homework.go

FROM alpine
COPY --from=build /opt/httpserver /opt/httpserver
EXPOSE 80
ENTRYPOINT ["/opt/httpserver"]

