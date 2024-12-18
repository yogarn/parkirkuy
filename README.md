# ParkirKuy
> ParkirKuy is an Android application that can be used to make parking reservations, so you can avoid parking lots that are full, especially in areas with heavy traffic. This application was created as a learning assignment regarding system analysis and design.

## Requirements
1. Golang
2. PostgreSQL
3. Docker (optional, preferred)

## Installation
1. Clone this repository
```
git clone https://github.com/yogarn/parkirkuy.git
```
2. Setup `.env`
```
mv .env.example .env && nano .env
```
3. Run Docker (if available)
```
docker compose up --build -d
```
or do it manually using golang
```
go mod tidy
go run ./cmd/app
```
## Reference
Front-End Repository: https://github.com/ahmadnafi30/parkirkuy
