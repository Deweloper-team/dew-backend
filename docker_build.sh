docker build -f server/Dockerfile -t tradesignal-backend:latest . --no-cache

docker-compose up -d --build
