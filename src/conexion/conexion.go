package conex

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func obtenerBaseDeDatos() (db *sql.DB, e error) {

	// Debe tener la forma usuario:contraseña@host/nombreBaseDeDatos
	db, err := sql.Open("mysql", "Johan:america1927@tcp(localhost:3306)/prueba?charset=utf8")
	if err != nil {
		return nil, err
	}

	return db, err

}

func Conex() *sql.DB {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		fmt.Printf("Error obteniendo base de datos: %v", err)
		return nil
	}
	// Terminar conexión al terminar función

	// Ahora vemos si tenemos conexión
	err = db.Ping()
	if err != nil {
		panic(err)

	}
	// Listo, aquí ya podemos usar a db!
	//fmt.Println("Conectado correctamente")

	return db
}
