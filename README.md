# How to run docker compose ###
## Step 1: Build docker images
### Frontend
```
cd next_start_over_fe-main
docker build -t frontend:0.0.1 .
```
### Backend
```
cd go_viet_tran1
docker build -t backend:0.0.1 .
```
## Step 2: Run docker compose
```
cd deployment
docker compose up -d
```

## Check
Access to `http://localhost/` to check your app.
