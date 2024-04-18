# Yandex Academy Golang - Final Project: Distributed Calculator
Вторая версия калькулятора на Golang. Добавлена авторизация пользователя, обновлена работа приложения. Проект представляет из себя калькулятор с сохранением математического выражения в базу данных. Перед тем, как воспользоваться приложением пользователь должен зарегистрироваться.

Backend - Golang

БД - PostgreSQL

Frontend - HTML, CSS, JS.

+Docker

## Содержание
- Как запустить приложение?
- Работа с приложением
- Схема API
- Полное ТЗ
- Контакты

## Как запустить приложение?

### Этап 0
Проект обернут в Docker, поэтому для установки и запуска нужно скачать [его](https://www.docker.com/products/docker-desktop/) (перейдите по ссылке и установите).  Проверить работает ли Docker Engine можно введя в терминал:  
```
docker
```  
(Должен появится список команд)

### Этап 1
Для запуска проекта необходина всего одна команда:  
```
docker-compose up
```  
Через некоторое время (до 5-7 минут) сборка проекта окочится и должны появится логи от самого приложения. 

![первые логи](https://github.com/roma106/golang_yandex-final/blob/main/frontend/static/imgs/screenshot-logs1.png "Первые логи")  

### Этап 2
Перейдите по ссылке: http://localhost:8080/register

или

Зажмите CTRL и кликните по ссылке на веб-страницу из лога

*Your link for webpage: http://localhost:8080/login* 

в терминале

### Запуск без докера

В случае если Вы не дружите с докером, (или я не успел обернуть в него проект к моменту сдачи)) запуск приложения также не составит много труда, но нужно будет вручную создать БД. Пожалуйста, пишите мне (контакты внизу).

## Работа с приложением
После перехода по ссылке Вы увидите веб-страничку с формой входа(фото 1).

![GUI](https://github.com/roma106/golang_yandex-final/blob/main/frontend/static/imgs/screenshot-register-page.png "Register")  

Для начала Вам нужно зарегистрироваться для этого перейдите на страницу Register нажав на ссылку вверху формы. Введите валидные данные и нажмите на кнопку "Register". Если все прошло успешно, вы получите доступ к калькулятору.

Весь интерфейс интуитивно понятен, если появляются вопросы или проблемы не стесняйтесь писать мне(ссылка на контакты).

## Схема API

![схема работы приложения](https://github.com/roma106/golang_yandex-final/blob/main/dia.drawio.png "схема работы приложения")

## Полное ТЗ 

Продолжаем работу над проектом Распределенный калькулятор
В этой части работы над проектом реализуем персистентность (возможность программы восстанавливать свое состояние после перезагрузки) и многопользовательский режим.
Простыми словами: все, что мы делали до этого теперь будет работать в контексте пользователей, а все данные будут храниться в СУБД

Функционал
Добавляем регистрацию пользователя
Пользователь отправляет запрос
POST /api/v1/register {
"login": ,
"password":
}
В ответ получае 200+OK (в случае успеха). В противном случае - ошибка
Добавляем вход
Пользователь отправляет запрос
POST /api/v1/login {
"login": ,
"password":
}
В ответ получае 200+OK и JWT токен для последующей авторизации.

Весь реализованный ранее функционал работает как раньше, только в контексте конкретного пользователя.
За эту часть можно получить 20 баллов
У кого выражения хранились в памяти - переводим хранение в SQLite. (теперь наша система обязана переживать перезагрузку)
За эту часть можно получить 20 баллов
У кого общение вычислителя и сервера вычислений было реализовано с помощью HTTP - переводим взаимодействие на GRPC
За эту часть можно получить 10 баллов

Дополнительные баллы:
- за покрытие проекта модульными тестами можно получить бонусные 10 баллов
- за покрытие проекта интеграционными тестами можно получить бонусные 10 баллов

Правила оформления:
- проект находится на GitHab - в ЛМС в решении вы сдаёте только ссылку на GitHab
- к проекту прилагается файл с подробным описанием (как заупустить и проверить)
- отдельным блоком идут подробно описанные тестовые сценарии
- постарайтесь автоматизировать поднятие окружения для запуска вашей программы (чтобы можно было это сделать одной командой)

## Контакты

ТГ - @Romanovski228  
Email - roma106ivanovskiy@mail.ru  

Отвечаю в течение часа
