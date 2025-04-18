# ใช้ Go image เป็น base
FROM golang:latest AS builder

# สร้าง working directory
WORKDIR /app

# Copy go.mod and go.sum แล้ว download dependency ก่อน (cache)
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy source code ทั้งหมด
COPY . .

# สร้าง binary
RUN go build -o main .

# ระบุ port ที่ Fiber จะรัน
EXPOSE 5001

# สั่งให้รันแอป
CMD ["/app/main"]
