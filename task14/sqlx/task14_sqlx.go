package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Age       int       `db:"age"`
	IsActive  bool      `db:"is_active"`
	CreatedAt time.Time `db:"created_at"`
}

type Order struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	Product   string    `db:"product"`
	Price     float64   `db:"price"`
	CreatedAt time.Time `db:"created_at"`
}

type Result struct {
	Name    string  `db:"name"`
	Age     int     `db:"age"`
	Product string  `db:"product"`
	Price   float64 `db:"price"`
}

func main() {

	dsn := "host=localhost user=postgres password=postgres dbname=testdb port=5432 sslmode=disable"

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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

	for _, u := range users {
		_, err := db.NamedExec(`
			INSERT INTO users(name,email,age,is_active)
			VALUES(:name,:email,:age,:is_active)
		`, u)

		if err != nil {
			log.Println(err)
		}
	}

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

	for _, o := range orders {
		_, err := db.NamedExec(`
			INSERT INTO orders(id,user_id,product,price)
			VALUES(:id,:user_id,:product,:price)
		`, o)

		if err != nil {
			log.Println(err)
		}
	}

	// ---------------- JOIN ----------------

	var result []Result

	err = db.Select(&result, `
		SELECT
			users.name,
			users.age,
			orders.product,
			orders.price
		FROM users
		JOIN orders
			ON users.id = orders.user_id
		WHERE orders.price >= $1
	`, 10000)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Результат:")

	for _, r := range result {
		fmt.Printf("%s (%d лет) -> %s : %.2f\n",
			r.Name,
			r.Age,
			r.Product,
			r.Price)
	}

	// ---------------- Транзакция ----------------

	fmt.Println("\n----- Транзакция -----")

	tx, err := db.Beginx()
	if err != nil {
		log.Fatal(err)
	}

	// если произойдет ошибка, транзакция откатится
	defer tx.Rollback()

	var userID int

	err = tx.QueryRow(`
		INSERT INTO users(name,email,age,is_active)
		VALUES($1,$2,$3,$4)
		RETURNING id
	`,
		"Павел Иванов",
		"pavel@example.com",
		26,
		true,
	).Scan(&userID)

	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.Exec(`
		INSERT INTO orders(user_id,product,price)
		VALUES($1,$2,$3)
	`,
		userID,
		"MacBook Pro",
		250000,
	)

	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Транзакция выполнена успешно")
}
