package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

// NewMariaDBConnection crée une connexion MariaDB à une base spécifique
func NewMariaDBConnection(user, password, host, port, database string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("impossible d'ouvrir la connexion à MariaDB : %w", err)
	}
	return db, nil
}

//// NewMariaDBAdminConnection crée une connexion administrative à MariaDB
//func NewMariaDBAdminConnection(user, password, host, port string) (*sql.DB, error) {
//	return NewMariaDBConnection(user, password, host, port, "mysql") // Base par défaut pour MariaDB
//}

// NewMariaDBAdminConnection crée une connexion administrative à MariaDB
func NewMariaDBAdminConnection(user, password, host, port, database string) (*sql.DB, error) {
	return NewMariaDBConnection(user, password, host, port, database)
}
