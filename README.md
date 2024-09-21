Получение информации о выбросах в STALCRAFT:X с помощью stalcraftAPI "https://eapi.stalcraft.net/".
Бот присылает сообщение о начале выброса. По запросу дает информацию о времени последнего выброса и времени прошедшем с его окончания.


### TODO

v1.2.0

- [x] Подключить sptf13/cobra для запуска приложения через cli
- [x] Подключить sptf13/viper для работы с конфигами и переменными среды(должен быть дефолтный конфиг, переопределяемый переменными среды)
- [x] Перевести логирование на zerolog, установить уровни логирования(info, debug, ...) устанавливаемые через конфиг
- [x] Разделить приложение на bot(работает постоянно), crawler(опрашивает сайт через заданные промежутки времени), admin-api(приложение для управление списком пользователей, удаление, бан, добавление). Приложения должны запускаться в одного файла, каждый через свою ключ, пример: ```./app --crawler```
- [x] Подготовить сервис api, который будет отдавать данные боту, сервис читает данные из файла, подготовленного crawler
- [x] Все апи должны отдавать swagger-page по префиксу /swagger, заполненную рабочими запросами 
- [x] Организовать общение приложений через FS-файлы(JSON или YAML)
- [x] Собрать приложение используя multi-stage Docker-image
- [x] Запускать все приложения через единый docker-compose.yml

v2.0.0

 - [x] Перевести хранение данных о выбросах с FS-storage на Redis-кэш, там же указать время хранения записей. При отсутствии записи в кэше сервис должен получить свежие данные и закэшировать их
- [x] Организовать хранение данных пользователей в SQL DB Postgres.
- [x] Подготовить миграции схемы данных на старте приложения(если схемы нет, оно должно их создать при старте). Использовать [goose](https://github.com/pressly/goose)
- [x] Postgres запускается через docker-compose до старта основных сервисов
- [x] Организовать ci/cd процесс на github-action для сборки и запуска приложения на проде при создании тэга. Переменные среды должны передаваться через github-repo-secrets

