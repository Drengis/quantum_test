# Quantum

Telegram Mini App для расчёта ипотеки.

## Архитектура

| Сервис | Технологии | Порт |
|--------|------------|------|
| Frontend | Next.js, TypeScript, Effector, TailwindCSS | 3000 |
| Backend | Go, Gin | 8081 |
| PostgreSQL | postgres:15-alpine | 5432 |
| Redis | redis:7-alpine | 6380 |

## Запуск

```bash
docker compose up -d --build
```

## Доступ

- **TMA:** <http://localhost:3000>
- **API:** <http://localhost:8081>
- **Swagger:** <http://localhost:8081/swagger>

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

| Переменная | По умолчанию |
|-------------|---------------|
| MAIN_APP_PORT | 8081 |
| DB_HOST | postgres |
| DB_PORT | 5432 |
| DB_USER | quantum |
| DB_PASSWORD | quantum |
| DB_DATABASE | quantum |
| REDIS_HOST | redis |
| REDIS_PORT | 6379 |