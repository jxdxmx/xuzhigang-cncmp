FROM golang:1.16-alpine AS build
#COPY module3 /httpserver
ADD . .
ENV GOMODULE111=ON
RUN GOOS=linux GOARCH=amd64 go build -o /httpserver homework.go

FROM alpine
COPY --from=build /httpserver httpserver
EXPOSE 80
ENTRYPOINT ["/httpserver"]
