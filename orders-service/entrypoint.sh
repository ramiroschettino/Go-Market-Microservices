#!/bin/sh
set -e

echo "⏳ Esperando a PostgreSQL..."
while ! nc -z $DB_HOST 5432; do
  sleep 1
done

echo "⏳ Esperando a products-service..."
while ! nc -z products-service 50051; do
  sleep 1
done

echo "🚀 Iniciando servidores..."
./grpcserver &
sleep 2  

exec ./httpserver