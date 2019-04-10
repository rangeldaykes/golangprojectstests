package transaction

import "github.com/gomodule/redigo/redis"

// Transaction is structure
type Transaction struct {
	err     error
	success func(interface{})
	reply   interface{}
}

var c redis.Conn

// NewTransaction return a new Transaction
func NewTransaction(redis redis.Conn) *Transaction {
	c = redis
	return &Transaction{}
}

// Do Execute a new command redis
func (t *Transaction) Do(cb func(conn redis.Conn)) *Transaction {
	//pool is a global object that has been setup in my app
	//c := pool.Get()
	defer c.Close()
	c.Send("MULTI")
	cb(c)
	reply, err := c.Do("EXEC")
	t.reply = reply
	t.err = err
	return t
}

// OnFail return when erros
func (t *Transaction) OnFail(cb func(err error)) *Transaction {
	if t.err != nil {
		cb(t.err)
	} else {
		t.success(t.reply)
	}
	return t
}

// OnSuccess return when success
func (t *Transaction) OnSuccess(cb func(reply interface{})) *Transaction {
	t.success = cb
	return t
}
