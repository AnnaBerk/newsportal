package main

import (
	"newsportal/internal/app"
)

func main() {
	app.Run()
}

//func initDB() (*gorm.DB, error) {
//	dsn := "host=localhost user=postgres password=mypassword dbname=postgres port=5432 sslmode=disable"
//	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	if err != nil {
//		return nil, fmt.Errorf("не удалось подключиться к базе данных: %w", err)
//	}
//	return db, nil
//}
//
//// тестовый вариант для проверки
//func main() {
//	db, err := initDB()
//	db = db.Debug()
//
//	if err != nil {
//		log.Fatalf("Ошибка инициализации базы данных: %v", err)
//	}
//
//	sqlDB, err := db.DB()
//	if err != nil {
//		log.Fatalf("Ошибка получения подключения к базе данных: %v", err)
//	}
//	defer sqlDB.Close()
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//
//	var categoryRepo repo.CategoryRepo = gormdb.NewCategoryRepo(db)
//	var tagRepo repo.TagRepo = gormdb.NewTagRepo(db)
//	var newsRepo repo.NewsRepo = gormdb.NewNewsRepo(db)
//	filter := repo.CategoryFilter{StatusID: 1}
//	categories, err := categoryRepo.GetCategoriesByFilter(ctx, filter)
//	if err != nil {
//		log.Printf("Ошибка: %v", err)
//		return
//	}
//	for _, category := range categories {
//		log.Println(category)
//	}
//	filterTag := repo.TagFilter{StatusID: 1}
//	tags, err := tagRepo.GetTagsByFilter(ctx, filterTag)
//	if err != nil {
//		log.Printf("Ошибка при получении тегов: %v", err)
//		return
//	}
//
//	if len(tags) == 0 {
//		log.Println("Нет тегов для отображения.")
//		return
//	}
//
//	for _, tag := range tags {
//		fmt.Printf("Tag ID: %v, Title: %v\n", tag.TagID, tag.Title)
//	}
//	filters := repo.NewsFilter{
//		CategoryID: 1,
//		TagID:      1,
//		StatusID:   1,
//		Limit:      2,
//		Offset:     1,
//	}
//	news, err := newsRepo.GetNewsByFilters(ctx, filters)
//	if err != nil {
//		log.Printf("Ошибка при получении новостей: %v", err)
//		return
//	}
//	for _, n := range news {
//		fmt.Println(n)
//	}
//	count, err := newsRepo.CountNewsByFilters(ctx, filters)
//	if err != nil {
//		log.Printf("Ошибка при подсчете новостей: %v", err)
//		return
//	}
//	fmt.Println(count)
//	newsID := int32(1)
//	n, err := newsRepo.GetNewsByID(ctx, newsID)
//	if err != nil {
//		log.Fatalf("Ошибка при получении новости: %v", err)
//	}
//
//	if n != nil {
//		log.Printf("Новость: %v", n.Title)
//	} else {
//		log.Printf("Новость с ID %d не найдена", newsID)
//	}
//
//	log.Println("Программа успешно завершена.")
//}
