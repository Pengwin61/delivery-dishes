# delivery-dishes












# CockroachDB

## Инициализация БД(кластера)
в режиме _insecure_
```
docker compose exec -it cockroachdb sh -c 'cockroach init --insecure --host=roach1'
```

в режиме сертификатов
```
cockroach sql --certs-dir=/certs --host=cockroachdb
```
