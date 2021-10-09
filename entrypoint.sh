# !/bin/bash
echo "waiting for mysql server"

while ! nc -z db 3306; do
  sleep 1
done

echo "Connection Successfully"

exec "$@"
exec air -c air.toml