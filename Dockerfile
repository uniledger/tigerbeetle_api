FROM golang:latest

RUN apt install git

WORKDIR /app

# Step 1: Copy only dependency files
COPY go.mod go.sum ./
RUN go mod download

# Step 2: Now copy source code
COPY . .
RUN go build -o tigerbeetle_api .

ENTRYPOINT ./tigerbeetle_api
