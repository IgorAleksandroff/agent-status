# Сервис управления статусами специалистов поддержки
## Общие требования

Service agent status представляет собой сервис, позволяющий управлять статусами специалистов службы поддержки.
Сервис должен реализовывать следующую бизнес-логику:
* регистрация, аутентификация и авторизация специалистов;
* хранение возможных состояний и переходов специалиста поддержки (начало смены, в работе чат, в работе письмо, перерыв, обед, наставничество, завершение смены, низкий поток, форс-мажор итд)
* получение текущего статуса специалиста и/или возможных переходов из текущего статуса;
* конкурентное изменение статуса специалиста из нескольких клиентов (авторизованным специалистом или из других сервисов);
* сервер должен давать специалисту возможность получить информацию о версии и дате сборки бинарного файла сервиса.

Типы клиентов и протоколы взаимодействия с ними:
* REST API – для интерфейса взаимодействия специалиста с сервисом;
* gRPC – для взаимодействия с другими сервисами (автоназначение диалогов с пользователями, автоназначение писем пользователей итп).

## Тестирование и документация

Код сервиса должен быть покрыт юнит-тестами не менее чем на 80%. Каждая экспортированная функция, тип, переменная, а также пакет системы должны содержать исчерпывающую документацию, в том числе описание REST API протокола взаимодействия клиента и сервера в формате Swagger. Наличие функциональных и/или интеграционных тестов приветствуется. 

## [Граф](https://miro.com/app/board/uXjVPjwYSlw=/) возможных состояний и переходов специалиста поддержки
![Граф](/images/AgentStatus.jpg)

