package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/user/quantum-server/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: go run cmd/migrate/main.go <command> [args]\n")
		fmt.Fprintf(os.Stderr, "Commands: up, down, force <version>, reset, make <name>\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	command := args[0]

	if command == "make" || command == "create" {
		if len(args) < 2 {
			log.Fatal("Укажите название миграции. Например: go run cmd/migrate/main.go make create_users_table")
		}

		name := args[1]
		timestamp := time.Now().Format("20060102150405")
		dir := filepath.Join("internal", "database", "migration")
		os.MkdirAll(dir, 0755)

		upFile := filepath.Join(dir, fmt.Sprintf("%s_%s.up.sql", timestamp, name))
		downFile := filepath.Join(dir, fmt.Sprintf("%s_%s.down.sql", timestamp, name))

		os.WriteFile(upFile, []byte("-- UP миграция\n"), 0644)
		os.WriteFile(downFile, []byte("-- DOWN миграция\n"), 0644)

		log.Printf("Файлы миграции созданы:\n%s\n%s\n", upFile, downFile)
		return
	}

	cfg := config.LoadConfig()
	dbCfg := cfg.DB

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Name)

	migrationPath := "file://internal/database/migration"
	m, err := migrate.New(migrationPath, dsn)
	if err != nil {
		log.Fatalf("Ошибка инициализации: %v", err)
	}

	switch command {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Ошибка при выполнении миграций: %v", err)
		} else if err == migrate.ErrNoChange {
			log.Println("Миграции не требуются (нет изменений)")
		} else {
			log.Println("Миграции успешно применены!")
		}
	case "down":
		if err := m.Steps(-1); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Ошибка при откате: %v", err)
		} else {
			log.Println("Последняя миграция успешно откачена!")
		}
	case "force":
		if len(args) < 2 {
			log.Fatal("Укажите версию для force")
		}
		var v int
		fmt.Sscanf(args[1], "%d", &v)
		if err := m.Force(v); err != nil {
			log.Fatalf("Ошибка force: %v", err)
		} else {
			log.Printf("Принудительно установлена версия %d\n", v)
		}
	case "reset":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Printf("Предупреждение при очистке: %v\n", err)
		}
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Ошибка при сбросе: %v", err)
		} else {
			log.Println("Сброс выполнен успешно!")
		}
	default:
		log.Fatalf("Неизвестная команда: %s", command)
	}
}
