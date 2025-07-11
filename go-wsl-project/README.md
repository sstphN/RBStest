всем здравствуйте

нужен голанг не ниже 1.21 и докер с wsl2 интеграцией

клонируй этот репозиторий или просто скопируй все файлы в папку которую назови go-wsl-project

перейди в папку проекта через терминал wsl2

-----

чтобы собрать сервер -> go build -o bin/server ./cmd/server 

чтобы запустить сервер -> ./bin/server

чтобы поднять базу данных -> docker-compose up -d

-----

как проверить что всё работает

проверь что сервер отвечает -> curl http://localhost:8080/ping

должен вернуть {"message":"pong"}

проверь соединение с базой -> curl http://localhost:8080/health

если всё ок то выведет -> {"status":"ok", ...}

проверь что можно получить список -> curl http://localhost:8080/list

если в базе есть записи увидишь массив

чтобы добавить новую запись -> curl -X POST -H "Content-Type: application/json" -d '{"name":"test"}' http://localhost:8080/add

-----

как проверить подключение к postgresql

чтобы зайти в базу через терминал -> docker exec -it myapp-postgres psql -U postgres -d myapp

чтобы посмотреть таблицы -> \dt

чтобы посмотреть содержимое -> select * from items;

-----

notes 

если база не создаёт таблицу items автоматически -> проверь папку migrations и перезапусти контейнер с флагом -v

если что-то не работает перезапусти docker desktop и сервер (мне помогало)

все переменные для подключения лежат в .env

проект тестировался на wsl2 ubuntu и windows 10

-----

этот проект сделан для тестового задания

feedback: 
tg @t7i3f 
mail hhhhhhh707@mail.ru

-----

<img width="1680" height="475" alt="image" src="https://github.com/user-attachments/assets/9db63a54-1717-46d2-8242-19e2f8bba98b" />
