package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Age       int
	IsActive  bool
	CreatedAt time.Time

	Orders []Order `gorm:"foreignKey:UserID"`
}

type Order struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `gorm:"not null"`
	Product   string  `gorm:"not null"`
	Price     float64 `gorm:"type:numeric(10,2)"`
	CreatedAt time.Time

	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Result struct {
	Name    string
	Age     int
	Product string
	Price   float64
}

func main() {

	dsn := "host=localhost user=postgres password=postgres dbname=testdb port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Создание таблиц
	db.AutoMigrate(&User{}, &Order{})

	// ---------------- Пользователи ----------------

	users := []User{
		{Name: "Иван Иванов", Email: "ivan@example.com", Age: 25, IsActive: true},
		{Name: "Анна Петрова", Email: "anna@example.com", Age: 30, IsActive: true},
		{Name: "Сергей Сидоров", Email: "sergey@example.com", Age: 19, IsActive: false},
		{Name: "Мария Смирнова", Email: "maria@example.com", Age: 27, IsActive: true},
		{Name: "Алексей Кузнецов", Email: "alexey@example.com", Age: 35, IsActive: true},
		{Name: "Елена Орлова", Email: "elena@example.com", Age: 28, IsActive: true},
		{Name: "Дмитрий Волков", Email: "dmitry@example.com", Age: 33, IsActive: false},
		{Name: "Ольга Морозова", Email: "olga@example.com", Age: 24, IsActive: true},
		{Name: "Максим Лебедев", Email: "maxim@example.com", Age: 41, IsActive: true},
		{Name: "Наталья Федорова", Email: "natalia@example.com", Age: 29, IsActive: true},
	}

	db.Create(&users)

	// ---------------- Заказы ----------------

	orders := []Order{
		{ID: 1, UserID: 1, Product: "Ноутбук", Price: 89999.99},
		{ID: 2, UserID: 1, Product: "Мышь", Price: 1999.00},
		{ID: 3, UserID: 2, Product: "Клавиатура", Price: 5499.00},
		{ID: 4, UserID: 2, Product: "Монитор", Price: 24999.99},
		{ID: 5, UserID: 3, Product: "Наушники", Price: 7999.90},
		{ID: 6, UserID: 4, Product: "Смартфон", Price: 65999.00},
		{ID: 7, UserID: 4, Product: "Чехол", Price: 1499.00},
		{ID: 8, UserID: 5, Product: "SSD 1TB", Price: 8999.99},
		{ID: 9, UserID: 6, Product: "Процессор", Price: 31999.00},
		{ID: 10, UserID: 7, Product: "Материнская плата", Price: 18999.00},
		{ID: 11, UserID: 8, Product: "Оперативная память", Price: 9999.99},
		{ID: 12, UserID: 9, Product: "Блок питания", Price: 7499.00},
		{ID: 13, UserID: 10, Product: "Видеокарта", Price: 72999.00},
		{ID: 14, UserID: 10, Product: "Монитор", Price: 29999.99},
		{ID: 15, UserID: 5, Product: "Игровая мышь", Price: 3499.00},
	}

	db.Create(&orders)

	// ---------------- JOIN ----------------

	var result []Result

	db.Table("users").
		Select("users.name, users.age, orders.product, orders.price").
		Joins("JOIN orders ON users.id = orders.user_id").
		Where("orders.price >= ?", 10000).
		Scan(&result)

	fmt.Println("Результат:")

	for _, r := range result {
		fmt.Printf("%s (%d лет) -> %s : %.2f\n",
			r.Name,
			r.Age,
			r.Product,
			r.Price)
	}

	fmt.Println("\n----- Транзакция -----")

	err = db.Transaction(func(tx *gorm.DB) error {

		user := User{
			Name:     "Павел Иванов",
			Email:    "pavel@example.com",
			Age:      26,
			IsActive: true,
		}

		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		order := Order{
			UserID:  user.ID,
			Product: "MacBook Pro",
			Price:   250000,
		}

		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		fmt.Println("Транзакция выполнена успешно")

		return nil
	})

	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}
