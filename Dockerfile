FROM 1.12.9-alpine3.9
COPY . /go/src/httpserver/
RUN go build -o /go/bin/httpserver /go/src/httpserver/main.go
ENTRYPOINT ["/go/bin/httpserver"]