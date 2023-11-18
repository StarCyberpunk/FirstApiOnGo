goose golang

 goose -dir ./migration  create init sql
goose -dir ./migration postgres "user=postgres password=12345 port=5432 host=localhost dbname=postgres sslmode=disable" up

https://thisis-blog.ru/chto-takoe-context-v-golang/
https://ru.hexlet.io/courses/go-web-development/lessons/auth/theory_unit