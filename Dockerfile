
FROM golang:1.24.4-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Replace with correct path if your file is not at root
RUN GOOS=linux GOARCH=arm64 go build -o migrate ./main.go

RUN chmod +x ./migrate

ENTRYPOINT ["./migrate"]

