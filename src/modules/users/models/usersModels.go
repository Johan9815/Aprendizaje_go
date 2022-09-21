package models

import (
	"fmt"
	"log"

	conex "github.com/johan9815/aprendizaje_go/src/conexion"
	"github.com/johan9815/aprendizaje_go/src/modules/users/structures"
)

var u structures.User

func GetUser(User, Pass string) (*structures.User, error) {
	db := conex.Conex()
	q := `SELECT * FROM users WHERE user = ? AND pass = ?`

	rows, err := db.Query(q, User, Pass)
	if err != nil {
		return &u, err
	}

	defer rows.Close()

	for rows.Next() {

		rows.Scan(
			&u.Id,
			&u.User,
			&u.Pass,
		)

	}
	return &u, nil

}

func GetAll() ([]structures.User, error) {

	db := conex.Conex()

	q := `SELECT * FROM users`
	// Ejecutamos la query
	rows, err := db.Query(q)
	if err != nil {
		return []structures.User{}, err
	}
	// Cerramos el recurso
	//defer rows.Close()

	usuarios := []structures.User{}
	// Declaramos un slice de notas para que almacene las
	// notas que retorna la petición.

	// El método Next retorna un bool, mientras sea true indicará
	// que existe un valor siguiente para leer.
	defer rows.Close()
	for rows.Next() {
		// Escaneamos el valor actual de la fila e insertamos el
		// retorno en los correspondientes campos de la nota.
		rows.Scan(
			&u.Id,
			&u.User,
			&u.Pass,
		)
		usuarios = append(usuarios, u)
		// Añadimos cada nueva nota al slice de notas que
		// declaramos antes.

	}
	return usuarios, nil

}

func Add(user structures.User) (int64, error) {

	db := conex.Conex()

	rows, err := db.Query("SELECT * FROM users WHERE user = ?", user.User)

	if err != nil {
		log.Fatal(err)
	}

	data := []structures.User{}

	for rows.Next() {
		rows.Scan(&u.Id, &u.User, &u.Pass)
		data = append(data, u)
	}
	defer rows.Close()

	if len(data) > 0 {
		return 0, nil
	}

	rows2, err := db.Exec("INSERT INTO users (user,pass) VALUES (?,?)", user.User, user.Pass)

	if err != nil {
		return 0, err
	}

	id, err := rows2.LastInsertId()

	if err != nil {
		return 0, err
	}

	db.Close()
	return id, nil

}

func UpdateUser(user structures.User) (structures.User, error) {

	db := conex.Conex()

	rows, err := db.Query("Update users set user = ?,pass = ? WHERE id = ?", user.User, user.Pass, user.Id)

	defer rows.Close()

	if err != nil {
		return structures.User{}, err
	}

	return user, nil

}

func DeleteUser(user structures.User) (int64, error) {

	db := conex.Conex()

	rows, err := db.Query("DELETE users from users Where id = ?", user.Id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	defer rows.Close()

	return int64(user.Id), nil

}
