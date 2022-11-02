# Avito-Tasks
___

### Avito-internship-backend

Это проект на GO созданный для Авито как тестовое задание.

Репозиторий с заданием: https://github.com/avito-tech/internship_backend_2022

В проекте надо было реализовать Микросервис для работы с балансом пользователей.

Из необходимых заданий было реализовано:

>Основное задание<br/>
>Доп.задание 1<br/>
>Доп.задание 2<br/>

Что не было сделано

>Докер <br/>
>Swagger

___

### Запуск.

Для запуска проекта необходимо прописать следующие команды:

​
go mod download<br/>
go run .\cmd\api\main.go


Так-же необходимо подключить бд при помощи Postgresql, скрипт с таблицами содержится в ./cmd/SqlSqript.sql.

ПОРТ:5432

Название самой бд: AVITO-TASKS

___
​
## Основные запросы

Для запросов я использовал Postman.

>Запрос который возвращает пользователя по id
```yaml
http://localhost:4000/user/getId/:id
```

>Запрос который добавляет денег пользователю на счет или создает пользователя если он отсуствует
```yaml
http://localhost:4000/user/add/:id/:user_cash
```

>Запрос который переводит деньги со счета одного пользователя на счет друго-го пользователя
```yaml
http://localhost:4000/user/transfer/:id/:ToId/:user_cash
```

>Запрос который снимает сумму денег со счета пользователя
```yaml
http://localhost:4000/user/withdrawal/:id/:user_cash
```

>Запрос который возвращает сервис по его id
```yaml
http://localhost:4000/service/get/:id
```

>Запрос создает сервис
```yaml
http://localhost:4000/service/create/:name/:price
```

>Запрос который резервирует услугу за пользователем
```yaml
http://localhost:4000/user/reserv/:id_user/:id_service
```

>Запрос создает создает платеж который еще не завершился окончательно
```yaml
http://localhost:4000/reporttemp/get/:id
```

>Запрос который подтверждает платеж 
```yaml
http://localhost:4000/report/create/:id
```

>Запрос который не подтверждает платеж 
```yaml
http://localhost:4000/report/notcreate/:id
```

>Запрос который возвращает платеж
```yaml
http://localhost:4000/report/getID/:id
```

>Запрос который создает историю платежей за определенный месяц в формате csv
```yaml
http://localhost:4000/report/get/:year/:month
```

>Запрос который открывает отчет в формате csv
```yaml
http://localhost:4000/file/:filename
```

>Запрос который возвращает транзакции определенного пользователя
```yaml
http://localhost:4000/transaction/get/:id
```

>Запрос который возвращает операции со счетом определенного пользователя
```yaml
http://localhost:4000/history/user/:id
```
