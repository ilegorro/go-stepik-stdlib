package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// SQLMap представляет карту, которая хранится в SQL-базе данных
type SQLMap struct {
	db         *sql.DB
	stmtGet    *sql.Stmt
	stmtSet    *sql.Stmt
	stmtDelete *sql.Stmt
}

// NewSQLMap создает новую SQL-карту в указанной базе
func NewSQLMap(db *sql.DB) (*SQLMap, error) {
	query := `
		create table if not exists map(key text primary key, val blob)
	`
	_, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	sGet, err := db.Prepare(`select val from map where key = ?`)
	if err != nil {
		return nil, err
	}
	sSet, err := db.Prepare(`
		insert into map(key, val) values (?, ?)
		on conflict (key) do update set val = excluded.val
	`)
	if err != nil {
		return nil, err
	}
	sDel, err := db.Prepare(`delete from map where key = ?`)
	if err != nil {
		return nil, err
	}
	return &SQLMap{
		db,
		sGet,
		sSet,
		sDel,
	}, nil
}

// Get возвращает значение для указанного ключа.
// Если такого ключа нет - возвращает ошибку sql.ErrNoRows.
func (m *SQLMap) Get(key string) (any, error) {
	row := m.stmtGet.QueryRow(key)
	var val any
	err := row.Scan(&val)

	return val, err
}

// Set устанавливает значение для указанного ключа.
// Если такой ключ уже есть - затирает старое значение (это не считается ошибкой).
func (m *SQLMap) Set(key string, val any) error {
	_, err := m.stmtSet.Exec(key, val)

	return err
}

// SetItems устанавливает значения указанных ключей.
func (m *SQLMap) SetItems(items map[string]any) error {
	tx, err := m.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	txStmt := tx.Stmt(m.stmtSet)

	for k, v := range items {
		_, err = txStmt.Exec(k, v)
		if err != nil {
			return err
		}
	}
	err = tx.Commit()

	return err
}

// Delete удаляет запись карты с указанным ключом.
// Если такого ключа нет - ничего не делает (это не считается ошибкой).
func (m *SQLMap) Delete(key string) error {
	_, err := m.stmtDelete.Exec(key)

	return err
}

// Close освобождает ресурсы, занятые картой в базе.
func (m *SQLMap) Close() error {
	if err := m.stmtGet.Close(); err != nil {
		return err
	}
	if err := m.stmtSet.Close(); err != nil {
		return err
	}
	if err := m.stmtDelete.Close(); err != nil {
		return err
	}

	return nil
}

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	m, err := NewSQLMap(db)
	if err != nil {
		panic(err)
	}
	defer m.Close()

	m.Set("name", "Alice")

	items := map[string]any{
		"name": "Bob",
		"age":  42,
	}
	m.SetItems(items)

	name, err := m.Get("name")
	fmt.Printf("name = %v, err = %v\n", name, err)
	// name = Bob, err = <nil>

	age, err := m.Get("age")
	fmt.Printf("age = %v, err = %v\n", age, err)
	// age = 42, err = <nil>
}
