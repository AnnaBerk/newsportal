package main

import (
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"newsportal/internal/repo/gormdb"
	"time"
)

func initDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=mypassword dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к базе данных: %w", err)
	}
	return db, nil
}

// тестовый вариант для проверки
func main() {
	db, err := initDB()
	db = db.Debug()

	if err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Ошибка получения подключения к базе данных: %v", err)
	}
	defer sqlDB.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tagRepo := gormdb.NewTagRepo(db)

	tags, err := tagRepo.GetTags(ctx)
	if err != nil {
		log.Printf("Ошибка при получении тегов: %v", err)
		return
	}

	if len(tags) == 0 {
		log.Println("Нет тегов для отображения.")
		return
	}

	for _, tag := range tags {
		fmt.Printf("Tag ID: %v, Title: %v\n", tag.TagID, tag.Title)
	}

	log.Println("Программа успешно завершена.")
}
