# !/bin/bash
echo "waiting for mysql server"

while ! nc -z db 3306; do
  sleep 1
done

echo "Connection Successfully"

exec "$@"
#debug毎にリモートが切れるため使用不可
#exec air -c air.toml
exec dlv exec ./tmp/main --listen=:40000 --headless=true --api-version=2