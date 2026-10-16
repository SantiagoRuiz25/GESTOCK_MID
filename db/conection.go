package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "tu_contraseña_de_pgadmin" // Cambia por tu contraseña de pgAdmin
	dbname := "gestock_db"

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("No se pudo comprobar la conexión con la base de datos: %v", err)
	}

	fmt.Println("¡Conexión exitosa a la base de datos!")
}