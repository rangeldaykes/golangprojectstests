package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/nakagami/firebirdsql"
)

// Linhatabela is a struct to fcv
type Linhatabela struct {
	Codigo             int
	Identificadorlinha string
}

func main() {

	conn, err := sql.Open("firebirdsql",
		"SYSDBA:masterkey@192.168.231.208/home/rangelsantos/discod/desenvolvimento/dbsfirebird/forquilinha/avls_20190128.fdb")

	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	printCountLinhasTabelas(*conn)
	printLinhasTabelasPk(*conn)
	printLinhasTabelasDataExec(*conn)

	fmt.Scanln()
}

func printCountLinhasTabelas(conn sql.DB) {
	var n int
	row := conn.QueryRow("SELECT Count(*) FROM linhas_tabelas")
	row.Scan(&n)

	fmt.Println("Relations count=", n)
}

func printLinhasTabelasPk(conn sql.DB) {
	rows, _ := conn.Query(`
		SELECT LT.CODIGO, LT.IDENTIFICADOR_LINHA
		  FROM LINHAS_TABELAS LT 
		 WHERE LT.codigo = ?`,
		873150)

	defer rows.Close()

	for rows.Next() {
		var u Linhatabela
		rows.Scan(&u.Codigo, &u.Identificadorlinha)
		fmt.Println(u)
	}
}

func printLinhasTabelasPk2(conn sql.DB) {
	rows, _ := conn.Query(`
		SELECT LT.CODIGO, LT.IDENTIFICADOR_LINHA
		  FROM LINHAS_TABELAS LT 
		 WHERE LT.codigo = ?`,
		873150)

	defer rows.Close()

	for rows.Next() {
		var u Linhatabela
		rows.Scan(&u.Codigo, &u.Identificadorlinha)
		fmt.Println(u)
	}
}

func printLinhasTabelasDataExec(conn sql.DB) {
	rows, _ := conn.Query(`
		SELECT LT.CODIGO, LT.IDENTIFICADOR_LINHA
		  FROM LINHAS_TABELAS LT 		
		 WHERE LT.DATA_EXECUCAO = ?`,
		time.Date(2019, 01, 24, 0, 0, 0, 0, time.UTC))

	defer rows.Close()

	for rows.Next() {
		var u Linhatabela
		rows.Scan(&u.Codigo, &u.Identificadorlinha)
		fmt.Println(u)
	}
}
