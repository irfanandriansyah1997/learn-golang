package learndatabase

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestInsertSQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// Delete all data before insert
	deleteScript := "DELETE FROM customer"
	_, errDelete := db.ExecContext(ctx, deleteScript)

	if errDelete != nil {
		panic(errDelete)
	}

	// Testing insert data
	script := `INSERT INTO customer (id, name, email, balance, rating, birth_date, married)
	VALUES
	  ('10912910', 'John Doe', 'john.doe1993@gmail.com', 100000, 5.0, '1993-5-12', true),
	  ('10912911', 'Arnold S.', 'arnold.the_terminator@gmail.com', 400000, 1.0, '1980-7-9', true);`
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert data")
}

type Customer struct {
	id, name, email        string
	balance                int64
	rating                 float64
	created_at, birth_date time.Time
	married                bool
}

func TestGetSQLAdvance(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	rows, err := db.QueryContext(ctx, "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer")

	if err != nil {
		panic(err)
	}

	var datum = make([]Customer, 0)
	cols, _ := rows.Columns()

	for rows.Next() {
		row := make([]any, len(cols))
		rowPtr := make([]any, len(cols))
		temp := Customer{}

		for i := range row {
			rowPtr[i] = &row[i]
		}

		if err := rows.Scan(rowPtr...); err != nil {
			panic(err)
		}

		for i, value := range cols {
			switch value {
			case "id":
				temp.id = string(row[i].([]uint8))

			case "name":
				temp.name = string(row[i].([]uint8))

			case "email":
				temp.email = string(row[i].([]uint8))

			case "balance":
				balance, _ := strconv.ParseInt(string(row[i].([]uint8)), 10, 32)
				temp.balance = balance

			case "rating":
				rating, _ := strconv.ParseFloat(string(row[i].([]uint8)), 10)
				temp.rating = rating

			case "birth_date":
				temp.birth_date = row[i].(time.Time)

			case "created_at":
				temp.created_at = row[i].(time.Time)
			}
		}

		datum = append(datum, temp)
	}

	defer rows.Close()
	for _, value := range datum {
		fmt.Println("===============================")
		fmt.Println("Id:", value.id)
		fmt.Println("Name:", value.name)
		fmt.Println("Email:", value.email)
		fmt.Println("Balance:", value.balance)
		fmt.Println("Rating:", value.rating)
		fmt.Println("Birth Date:", value.birth_date)
		fmt.Println("Married:", value.married)
		fmt.Println("Created At:", value.created_at)
	}
}
