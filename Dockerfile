FROM golang:1.23

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod tidy

COPY . .

RUN go build -o go-account-revocery cmd/main.go

RUN chmod +x go-account-revocery

EXPOSE 9999

CMD [ "./go-account-revocery" ]

