# Тестовый проект с использованием Memcached
По материалам видео https://www.youtube.com/watch?v=0aSaQVMzANg

## Клиент для go
```bash
go get github.com/bradfitz/gomemcache/memcache
```

## Установка memchached в docker
```bash
docker pull memcached
```

## Запуск memcached в docker
```bash
docker run -p 11211:11211 --name my-memcache -d memcached
```

## Тестовые данные
```json
{
  "ID": 1,
  "CreatedAt": "2024-11-06T13:00:00Z",
  "UpdatedAt": "2024-11-06T13:00:00Z",
  "DeletedAt": null,
  "Title": "Master & Commander of the Seas",
  "Content": "Фильм 2003 года. В главной роли Рассел Кроу. Один из самых реалистичных фильмов о быте моряков парусного флота начала 19 века"
}
```
