# !/bin/bash

echo "waiting for redies server"
while ! nc -z redis 6379; do
  sleep 1
done
echo "redis Connection Successfully"

echo "waiting for mysql server"
while ! nc -z db 3306; do
  sleep 1
done
echo "mysql Connection Successfully"

exec "$@"
exec air -c air.toml