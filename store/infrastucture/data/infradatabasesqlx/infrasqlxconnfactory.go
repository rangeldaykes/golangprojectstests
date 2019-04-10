package infradatabasesqlx

import (
	"log"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"

	// Firebird connect
	_ "github.com/nakagami/firebirdsql"

	// Postgres connect
	_ "github.com/lib/pq"
)

//var connectionStringFirebird = `SYSDBA:masterkey@192.168.231.208/home/rangelsantos/discod/desenvolvimento/dbsfirebird/forquilinha/avls_20190314.fdb`
var connectionStringFirebird = `SYSDBA:masterkey@127.0.0.1/home/rangelsantos/discod/desenvolvimento/dbsfirebird/store.fdb`

var connectionStringPostgres = "postgres://postgres:postgres@127.0.0.1/store?sslmode=disable"

type IConnDataBase interface {
	GetConn() *sqlx.DB
}

type connDataBase struct {
	conn *sqlx.DB
}

var (
	cdb      *connDataBase
	connOnce sync.Once
)

func ConnDataBaseGetInstance(connectFirebied bool) IConnDataBase {
	if cdb == nil {
		connOnce.Do(func() {
			if connectFirebied {
				cdb = &connDataBase{conn: createConnFirebird()}
			} else {
				cdb = &connDataBase{conn: createConnPostgres()}
			}
		})
	}

	return cdb
}

func (c connDataBase) GetConn() *sqlx.DB {
	return c.conn
}

func createConnFirebird() *sqlx.DB {
	c, err := sqlx.Connect("firebirdsql", connectionStringFirebird)
	c.SetMaxIdleConns(5)
	c.SetMaxOpenConns(100)
	c.SetConnMaxLifetime(time.Minute * 5)

	if err != nil {
		log.Fatalln(err)
	}

	return c
}

func createConnPostgres() *sqlx.DB {
	c, err := sqlx.Connect("postgres", connectionStringPostgres)
	c.SetMaxIdleConns(5)
	c.SetMaxOpenConns(100)
	c.SetConnMaxLifetime(time.Minute * 5)

	if err != nil {
		log.Fatalln(err)
	}

	return c
}

// GetConn get connection with sqlx
// func GetConn() *sqlx.DB {
// 	db, err := sqlx.Connect("firebirdsql",
// 		"SYSDBA:masterkey@192.168.231.208/home/rangelsantos/discod/desenvolvimento/dbsfirebird/forquilinha/avls_20190128.fdb")

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	return db
// }

/* // GetConn get connection with sqlx
func GetConn() *sqlx.DB {
	//if (Cfg.RelationalDB.isFirebird)

	db, err := sqlx.Connect("firebirdsql",
		"SYSDBA:masterkey@192.168.231.208/home/rangelsantos/discod/desenvolvimento/dbsfirebird/forquilinha/avls_20190128.fdb")

	if err != nil {
		log.Fatalln(err)
	}

	return db
	//else
	//return new Npgsql.NpgsqlConnection(Cfg.RelationalDB.StringConexao);
} */
