# greenlight

Libs:
- golang-migrate to handle migrations

Generate a migration
```
migrate create -seq -ext=.sql -dir=./migrations create_movies_table
```