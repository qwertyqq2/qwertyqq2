REST-API для платежной системы. 
========================
Описание:
-------------------------

  Из вне приходит запрос, и роутер его обрабатывает. Существует четыре различных обработчика: Создание транзакции, Создание юзера, Получение всех транзакций по Id и по Email. В качестве БД я использовал MYSQL. На своем устройстве подключался к ней как test_user:pass@/TASK, где TASK - бд, содержащая таблицы хранящие юзеров и транзакции.  В папке "DB" содержаться миграция моей бд, то есть код и  все начальные инструкции.
Для создание POST запросов на сервер я использовал интрукцию curl, например:  
curl -d '{"email" : "qwerty@mail.ru"}' -H "Content-Type: application/json" -X POST http://localhost:8070/user  

