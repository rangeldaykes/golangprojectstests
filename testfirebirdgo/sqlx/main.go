package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/nakagami/firebirdsql"
)

type usuario struct {
	Id   int
	Nome string
}

type linhatabela struct {
	Codigo             int    `db:"CODIGO"`
	Identificadorlinha string `db:"IDENTIFICADOR_LINHA"`
}

func main() {

	//TestMySql()
	TestFirebird()

	//fmt.Scanln()
}

func TestFirebird() {
	//db, err := sql.Open("firebirdsql",
	//"SYSDBA:masterkey@192.168.231.208/home/rangelsantos/discod/desenvolvimento/dbsfirebird/forquilinha/avls_20190128.fdb")

	db, err := sqlx.Connect("firebirdsql",
		"SYSDBA:masterkey@192.168.231.208/home/rangelsantos/discod/desenvolvimento/dbsfirebird/forquilinha/avls_20190314.fdb")

	if err != nil {
		log.Fatalln(err)
	}

	s := struct {
		Codigo int
	}{
		1028894,
	}

	var people2 string
	fcvsSqlxNamedGetOne(&people2, db, s)
	fmt.Println(people2)

	defer db.Close()
}

func fcvsSqlxNamedSelect(dest interface{}, db *sqlx.DB, param interface{}) error {
	nstmt, err := db.PrepareNamed(`
	SELECT LT.CODIGO, LT.IDENTIFICADOR_LINHA
	  FROM LINHAS_TABELAS LT 
	 WHERE LT.codigo = :codigo`)

	err = nstmt.Select(dest, param)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func fcvsSqlxNamedGetOne(dest interface{}, db *sqlx.DB, param interface{}) error {
	//SELECT LT.CODIGO, LT.IDENTIFICADOR_LINHA
	nstmt, err := db.PrepareNamed(`
	SELECT LT.IDENTIFICADOR_LINHA	
	  FROM LINHAS_TABELAS LT 
	 WHERE LT.codigo = :codigo`)

	err = nstmt.Get(dest, param)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func fcvsSqlx(db *sqlx.DB) {

	people := []linhatabela{}
	db.Select(&people, `
	SELECT LT.CODIGO, LT.IDENTIFICADOR_LINHA
	  FROM LINHAS_TABELAS LT 
	 WHERE LT.codigo = ?`, 1028894)

	fmt.Println(people)
}

func fcvsSql(conn *sql.DB) {
	rows, _ := conn.Query(`
		SELECT LT.CODIGO, LT.IDENTIFICADOR_LINHA
		  FROM LINHAS_TABELAS LT 
		 WHERE LT.codigo = ?`,
		1028894)

	defer rows.Close()

	for rows.Next() {
		var u linhatabela
		rows.Scan(&u.Codigo)
		//rows.Scan(&u.Codigo, &u.Identificadorlinha)
		fmt.Println(u)
	}
}

func TestMySql() {
	//db, err := sql.Open("mysql", "root:usuubu@tcp(192.168.231.208)/cursogo")
	db, err := sqlx.Connect("mysql", "root:usuubu@tcp(192.168.231.208)/cursogo")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//getUsuariosSql(db.DB)
	getUsuariosSqlx(db)
}

func getUsuariosSqlx(db *sqlx.DB) {
	people := []usuario{}
	db.Select(&people, "select id, nome from usuarios where id = ?", 1)

	fmt.Println(people)
}

func getUsuariosSql(db *sql.DB) {
	rows, _ := db.Query("select id, nome from usuarios where id = ?", 1)

	defer rows.Close()

	for rows.Next() {
		var u usuario
		rows.Scan(&u.Id, &u.Nome)
		fmt.Println(u)
	}
}
