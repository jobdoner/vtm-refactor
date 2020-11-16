FROM golang:1.15 as app
WORKDIR /jobdoner
COPY . .
RUN cd vtm-refactor/ && GO111MODULE=on \
    go install cmd/vtm-refactor/main.go 
EXPOSE 8080

CMD ["main"]
