docker build -f server/Dockerfile -t dew-backend:latest . --no-cache

docker-compose up -d --build
