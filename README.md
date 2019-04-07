# mojochat

[![pipeline status](https://gitlab.com/le-garff-yoann/mojochat/badges/master/pipeline.svg)](https://gitlab.com/le-garff-yoann/mojochat/pipelines)

> An anonymous chat with Vue, Gorilla and PostgreSQL.

## TL;DR

```bash
# Take a look at vue/mojochat/README.md

go build -o mojochat main.go

MOJOCHAT_PORT=8080 \
MOJOCHAT_PG=postgresql://u:p@pg/mojochat?sslmode=disable \ # Assuming that you have a PostgreSQL running around...
    ./mojochat
```
