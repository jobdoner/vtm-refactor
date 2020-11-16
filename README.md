# vtm (Video Target Manager)
Project contains REST API to manipulate folders of campains & adgroups (more in swagger.yaml) 
and document generation based on campaings and adgroups.
## TODO:
- Цикл обновления ридми
- перенести со свагер движка на стд либу, сделать фронтенд  на шаблонах
- перенести генерацию темлейтов выгружаемых документов для интеграций на git движек  с интеграцией с системой
- Вынести авторизацию в отдельный контур и сервис генерации в отдельный сервис по рпц
- Вынести миграции в отдельный сервис миграций и переделать структуру
## Deployment
Just
`docker-compose up`
more info about deployment can be found in docker-compose.yaml file
## Technologies Used
- Go
- Redis 
- Postgres
- Swagger