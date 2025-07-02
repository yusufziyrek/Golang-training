#!/bin/bash

# Config
CONTAINER_NAME="postgres-test"
POSTGRES_USER="postgres"
POSTGRES_PASSWORD="1234"
POSTGRES_DB="productapp"
HOST_PORT="6432"
CONTAINER_PORT="5432"
IMAGE_NAME="postgres:latest"  # latest olarak sabit

# Önceki konteyner varsa sil
if [ "$(docker ps -aq -f name=$CONTAINER_NAME)" ]; then
  echo "Önceki konteyner tespit edildi. Durduruluyor ve siliniyor..."
  docker rm -f $CONTAINER_NAME
fi

# PostgreSQL konteyneri başlat
echo "Yeni PostgreSQL konteyneri başlatılıyor..."
docker run --name $CONTAINER_NAME \
  -e POSTGRES_USER=$POSTGRES_USER \
  -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD \
  -e POSTGRES_DB=$POSTGRES_DB \
  -p $HOST_PORT:$CONTAINER_PORT \
  -d $IMAGE_NAME

# Veritabanı hazır olana kadar bekle
echo "Veritabanı başlatılıyor, lütfen bekleyin..."
sleep 5

# Veritabanı oluştur (gerekli değilse bile güvenlik amaçlı)
docker exec -i $CONTAINER_NAME psql -U $POSTGRES_USER -d postgres -c "CREATE DATABASE $POSTGRES_DB;"
echo "Veritabanı '$POSTGRES_DB' oluşturuldu."

# Tabloları oluştur
docker exec -i $CONTAINER_NAME psql -U $POSTGRES_USER -d $POSTGRES_DB <<EOF
CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DOUBLE PRECISION NOT NULL,
    discount DOUBLE PRECISION,
    store VARCHAR(255) NOT NULL
);
EOF

echo "Tablo 'products' başarıyla oluşturuldu ✅"
