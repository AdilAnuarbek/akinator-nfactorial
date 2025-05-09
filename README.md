
# Akinator clone

Данный сайт это клон популярного веб-сайта Акинатор, джина который может угадать любого персонажа которого вы загадаете!

**Данный проект был создан исключительно в образовательных целях и никак не пытается выдать себя за оригинального Акинатора!**

Оригинальный веб-сайт [Акинатора](https://en.akinator.com/).

Проект был написано на Golang и использовал Node.js модуль который делал запросы на веб-сайт Акинатора.
## Установка

Для того чтобы использовать данный проект на своём ПК, вам нужен [golang](https://go.dev/doc/install) и [npm](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm) скачанный на вашем ПК. Затем, скачайте репозиторий или склонируйте его и в нужную папку:
```bash
git clone https://github.com/AdilAnuarbek/akinator-nfactorial
```
Откройте папку проекта в коммандной строке:
```bash
cd my-project
```
Установите нужные библиотеки npm:
```bash
npm install cors express node_akinator
```

## Запуск

Откройте две коммандной строки и введите каждый из следующих комманд:

```bash
node api/node_api/server.js
```
```bash
go run api/go_api/server.go
```

## Процесс проектирования и разработки
При проектирований, проект был разделен на следующие этапы:

- Нахождение способа получать запросы от Акинатора.
- Написание базовой логики для отображения простой html страницы.
- Написание и связание логики получения данных от пользователя, запроса данных в Акинатора, получения ответа и отображения ответа на странице для пользователя.
- Дизайн страницы
- Верстка приложения на публичный доступ

В начале я написал базовый код на го который просто показывал пустой html файл, затем писал серверную часть, а затем написал супер простой дизайн. Но когда дошло до верстки сайта на vercel.com, у меня были проблемы. Также, из-за того что время дедлайна уже подходило, после многочисленных попыток, я решил не верстать сайт. Весь опыт от начала разработки, до тестинга и попытки верстки, показала мне что разработка сайта на скорости (у меня финальные экзамены семестра на момент написания данного текста) очень отличается от более медленной разработки. 

## Уникальные подходы

Один из уникальных подходов которые были использованы при работе было использование express.js сервера для отправки и получения запросов на сервера Акинатора и отправки и получения запросов с основного Golang сервера на раннего сервера express.js. Использование двух серверов для одной общей цели было то что я еще не делал и скорее всего не часто применяется в настоящих серверах.

## Известные ошибки

Это более недочет из-за нехватки времени и опыта, но дизайн оказался почти что никаким. Всё что я успел сделать это header и footer с минимальным css. Откровенных ошибок не было, кроме как того что я потратил много времени на попытку верстки сайта.

## Почему этот стак

Я уже писал на Golang веб-сайт который был для меня первым на данном стэке и был довольно комплексный. Из-за этого я решил не придумывать ничего нового и решил снова использовать Golang с Chi для роутинга. Express для части джаваскрипта решил использовать потому-что он был довольно простой и быстрый для усвоения. Также было несколько других API запросчиков на сайт Акинатора, но я выбрал node_akinator потому что он выглядел наиболее простым для использования.  
