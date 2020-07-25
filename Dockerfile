FROM golang:1.8
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
## RUN go get github.com/tealeg/xlsx && go build -o xlsx2csv . 
RUN go get github.com/naototty/xlsx2csv-go && go build -o xlsx2csv-go . 
CMD ["/app/xlsx2csv-go"]
