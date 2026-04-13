# Quantum Server

Server implementation in Go.

## Structure

.
├── cmd/
│   ├── api-server/        # Основной сервис
│   │   └── main.go
│   └── worker/            # Фоновый обработчик
│       └── main.go
├── internal/
│   ├── auth/              # Логика аутентификации
│   ├── database/          # Работа с БД
│   └── service/           # Бизнес-логика
├── api/
│   └── swagger.yaml
├── go.mod                 # Управление зависимостями
├── go.sum
└── README.md
