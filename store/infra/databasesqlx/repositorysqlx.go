package databasesqlx

type IInfraRepositorySqlx interface {
	PrepareNamedSelect(dest interface{}, sql string, param interface{}) error
	PrepareNamedGetOne(dest interface{}, sql string, param interface{}) error
}

type InfraRepositorySqlx struct {
	IConnDataBase
}

func NewInfraRepositorySqlx(conn IConnDataBase) *InfraRepositorySqlx {
	return &InfraRepositorySqlx{conn}
}

// PrepareNamedSelect get data from database with params named
// use dest param: as slice of []struct to return a data,
// sql param: your sql query and
// param:  as struct with fields that you want to pass to the query (anonymous struc is great)
func (rs *InfraRepositorySqlx) PrepareNamedSelect(dest interface{}, sql string, param interface{}) error {
	c := rs.GetConn()

	nstmt, err := c.PrepareNamed(sql)
	if err != nil {
		return err
	}

	err = nstmt.Select(dest, param)
	if err != nil {
		return err
	}

	return nil
}

// PrepareNamedGetOne
func (rs *InfraRepositorySqlx) PrepareNamedGetOne(dest interface{}, sql string, param interface{}) error {
	c := rs.GetConn()

	nstmt, err := c.PrepareNamed(sql)
	if err != nil {
		return err
	}

	err = nstmt.Get(dest, param)
	if err != nil {
		return err
	}

	return nil
}

// Select return one instace from query
// func (rs *InfraRepositorySqlx) Select(dest interface{}, query string, args ...interface{}) error {
// 	c := rs.conn()
// 	defer c.Close()

// 	return c.Select(&dest, query, args)
// }
