# Quantum Server

API для расчёта ипотечных профилей.

## Описание

Go-сервис для расчёта ипотеки с использованием:

- **Go** (Gin framework)
- **PostgreSQL** - хранение данных
- **Redis** - кэширование
- **Docker** - контейнеризация

## Функционал

- Создание ипотечного профиля (POST /mortgage-profiles)
- Получение расчёта по ID (GET /mortgage-profiles/:id)
- Асинхронная обработка расчётов через воркер

## Требования

- Docker
- Docker Compose

## Запуск

```bash
# Создать .env из .env.example (если нет)
cp .env.example .env

# Собрать и запустить
docker-compose up --build
```

## Доступ

- **API:** <http://localhost:8081>
- **Swagger UI:** <http://localhost:8081/swagger>

## Эндпоинты

### POST /mortgage-profiles

Создание ипотечного профиля.

```json
{
  "user_id": "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
  "propertyPrice": 6000000,
  "propertyType": "apartment_in_new_building",
  "downPaymentAmount": 1200000,
  "matCapitalIncluded": false,
  "mortgageTermYears": 20,
  "interestRate": 8.5
}
```

### GET /mortgage-profiles/:id

Получение расчёта по ID.

## Переменные окружения

| Переменная    | По умолчанию |
|---------------|--------------|
| MAIN_APP_PORT | 8081         |
| DB_HOST       | localhost    |
| DB_PORT       | 5432         |
| DB_USER       | quantum      |
| DB_PASSWORD   | quantum      |
| DB_DATABASE   | quantum      |
| REDIS_HOST    | localhost    |
| REDIS_PORT    | 6379         |
