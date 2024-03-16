# Профильное задание Go разработчик (03/2024)

### Что реализовано:

1. POST запрос на добавление Фильма и Актеров:

```json
POST http://localhost:8080/api/v1/films
```
Тестовый запрос:
```json
{
    "film": {
        "title": "The Matrix New Revolutions part 4",
        "description": "Picking up immediately where Reloaded ended, Neo and Bane still lie unconscious in the medical bay of the ship Hammer. Inside the Matrix, Neo is trapped in a subway station named Mobil Ave, a transition zone between the Matrix and the machine world.",
        "year": 2005,
        "rating": 7
    },
    "actors": [
        {
            "name": "Keanu Reeves5",
            "gender": "Male",
            "birth_date": "1964"
        },
        {
            "name": "Hugo Weaving5",
            "gender": "Male",
            "birth_date": "1960"
        }
    ]
}
```
2. Получение всех фильмов с сортировкой по умолчанию по полю Рейтинг:

```json
GET http://localhost:8080/api/v1/films
```
Ответ:
```json
[
    {
        "id": 2,
        "title": "The Matrix New Revolutions part 1",
        "description": "Picking up immediately where Reloaded ended, Neo and Bane still lie unconscious in the medical bay of the ship Hammer. Inside the Matrix, Neo is trapped in a subway station named Mobil Ave, a transition zone between the Matrix and the machine world.",
        "year": 2003,
        "rating": 10
    },
    {
        "id": 3,
        "title": "The Matrix New Revolutions part 2",
        "description": "Picking up immediately where Reloaded ended, Neo and Bane still lie unconscious in the medical bay of the ship Hammer. Inside the Matrix, Neo is trapped in a subway station named Mobil Ave, a transition zone between the Matrix and the machine world.",
        "year": 2004,
        "rating": 9
    },
    {
        "id": 4,
        "title": "The Matrix New Revolutions part 3",
        "description": "Picking up immediately where Reloaded ended, Neo and Bane still lie unconscious in the medical bay of the ship Hammer. Inside the Matrix, Neo is trapped in a subway station named Mobil Ave, a transition zone between the Matrix and the machine world.",
        "year": 2004,
        "rating": 9
    },
    {
        "id": 1,
        "title": "The Matrix New Revolutions",
        "description": "Picking up immediately where Reloaded ended, Neo and Bane still lie unconscious in the medical bay of the ship Hammer. Inside the Matrix, Neo is trapped in a subway station named Mobil Ave, a transition zone between the Matrix and the machine world.",
        "year": 2025,
        "rating": 6
    }
]
```
3. Получение фильма по ID:

```json
GET http://localhost:8080/api/v1/film/1
```
Ответ:
```json
{
    "id": 1,
    "title": "The Matrix New Revolutions",
    "description": "Picking up immediately where Reloaded ended, Neo and Bane still lie unconscious in the medical bay of the ship Hammer. Inside the Matrix, Neo is trapped in a subway station named Mobil Ave, a transition zone between the Matrix and the machine world.",
    "year": 2025,
    "rating": 6
}
```
Реализовано с получением токена для юзера
>login: admin
>password : admin123

Запросом:
```json
POST http://localhost:8080/api/v1/user/auth
```
```json
{
    "login": "admin",
    "password": "admin123"
}
```


4. Получение всех актеров:

```json
GET http://localhost:8080/api/v1/actors
```
Ответ:
```json
[
    {
        "id": 1,
        "name": "Keanu Reeves1",
        "gender": "Male",
        "birth_date": "1964"
    },
    {
        "id": 2,
        "name": "Hugo Weaving",
        "gender": "Male",
        "birth_date": "1960"
    },
    {
        "id": 3,
        "name": "Keanu Reeves2",
        "gender": "Male",
        "birth_date": "1964"
    },
    {
        "id": 4,
        "name": "Hugo Weaving1",
        "gender": "Male",
        "birth_date": "1960"
    },
    {
        "id": 5,
        "name": "Keanu Reeves3",
        "gender": "Male",
        "birth_date": "1964"
    },
    {
        "id": 6,
        "name": "Hugo Weaving3",
        "gender": "Male",
        "birth_date": "1960"
    },
    {
        "id": 7,
        "name": "Keanu Reeves4",
        "gender": "Male",
        "birth_date": "1964"
    },
    {
        "id": 8,
        "name": "Hugo Weaving4",
        "gender": "Male",
        "birth_date": "1960"
    }
]
```
5. Создание и удаление актера:

```json
POST http://localhost:8080/api/v1/actors
```
```json
DELETE http://localhost:8080/api/v1/actors/{id}
```

Тестовый запрос:
```json
{
    "name": "Ian McKellens1",
    "gender": "Male2",
    "birth_date": "1939"
}
```
6. Создание юзера:
```json
POST http://localhost:8080/api/v1/user/register
```

Тестовый запрос:
```json
{
    "login": "bigadmin",
    "password": "adminos"
}
```
7. Прочие:

Реализована проверка на создание дубликата фильма, проверяется название и год

Таблицы базы данных реализованы как 3 нормальная форма

>films, actors, filmactor (содержит ключи actorID и filmID)

БД postgresql с настроенной миграцией внутри докера
Docker и docker compose сгенерированы
Коллекция запросов сохранена в корне

Из минусов:
тестами не покрыто
документация в свагер не описана

Спасибо за ваше время
Всего доброго!

