package db

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"time"

	"github.com/RomeroGabriel/go-api-standards/internal/entity"
)

type ProductDB struct {
	DB *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS Product 
		(Id VARCHAR(255) NOT NULL PRIMARY KEY, Name VARCHAR(255), Price INTEGER, CreatedAt VARCHAR(255));
	`
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	conn, err := db.Conn(ctx)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil
	}
	defer conn.Close()
	_, err = conn.ExecContext(ctx, sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil
	}
	return &ProductDB{DB: db}
}

func (p *ProductDB) Create(product *entity.Product) error {
	ctx := context.Background()
	conn, err := acquireConnection(ctx, p.DB)
	if err != nil {
		return err
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, "INSERT INTO Product (Id, Name, Price, CreatedAt) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price, product.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductDB) FindByID(id string) (*entity.Product, error) {
	ctx := context.Background()
	conn, err := acquireConnection(ctx, p.DB)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, "SELECT Id, Name, Price, CreatedAt FROM Product WHERE Id = ?")
	if err != nil {
		return nil, err
	}
	var result entity.Product
	var date string
	err = stmt.QueryRowContext(ctx, id).Scan(&result.ID, &result.Name, &result.Price, &date)
	if err != nil {
		return nil, err
	}
	dateTime, err := time.Parse("2006-01-02 15:04:05.999999999-03:00", date)
	if err != nil {
		return nil, err
	}
	result.CreatedAt = dateTime
	return &result, nil
}

func (p *ProductDB) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		return nil
	}
	ctx := context.Background()
	conn, err := acquireConnection(ctx, p.DB)
	if err != nil {
		return err
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, "UPDATE Product SET Name = ?, Price = ? WHERE Id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductDB) Delete(product *entity.Product) error {
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		return nil
	}
	ctx := context.Background()
	conn, err := acquireConnection(ctx, p.DB)
	if err != nil {
		return err
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, "DELETE FROM Product WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(product.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductDB) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	stringQuery := strings.Replace(
		"SELECT Id, Name, Price, CreatedAt FROM Product ORDER BY CreatedAt SORT LIMIT ? OFFSET ?", "SORT", sort, 1)
	ctx := context.Background()
	conn, err := acquireConnection(ctx, p.DB)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	offset := (page - 1) * limit
	rows, err := conn.QueryContext(ctx, stringQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []entity.Product
	for rows.Next() {
		var pr entity.Product
		var date string
		if err := rows.Scan(&pr.ID, &pr.Name, &pr.Price, &date); err != nil {
			return nil, err
		}
		dateTime, err := time.Parse("2006-01-02 15:04:05.999999999-03:00", date)
		if err != nil {
			return nil, err
		}
		pr.CreatedAt = dateTime
		products = append(products, pr)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}
