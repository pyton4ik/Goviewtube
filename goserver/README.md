# Inro
Взято отсюда
https://github.com/ViewTube/viewtube

Хорошо продуманная архитектура сразу поделена на фрон и бек
С фронтом все хорошо в вот бэк заменим на свой, Go-шный.
Ничего особы выпиливать не будем. Папку server со старым беком оставим
и создадим совою.

# Mac os Developer Run
## Запустить фронт локально
pnpm run serve:client
pnpm run serve:server

## Открыть фронт локально
http://localhost:8066/ 

## Открыть IP локально
http://localhost:8067/api/homepage/popular
http://localhost:8067/api/user/subscriptions/videos?limit=4&start=0


