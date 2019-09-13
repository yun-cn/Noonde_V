package http

import (
	cont "context"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/yun313350095/Noonde/api"
	"net/http"
)

type context struct {
	context      cont.Context
	request      *http.Request
	params       httprouter.Params
	tx           *sqlx.Tx
	elasticQuery *api.ElasticQuery
	locale       string
	curUser      *api.User
}

// Context ..
func (c *context) Context() cont.Context {
	return c.context
}

// CurUser ..
func (c *context) CurUser() *api.User {
	return c.curUser
}

// ElasticQuery ..
func (c *context) ElasticQuery() *api.ElasticQuery {
	return c.ElasticQuery()
}

// Locale ..
func (c *context) Locale() string {
	return c.locale
}

// Params ..
func (c *context) Params() httprouter.Params {
	return c.params
}

// Request ..
func (c *context) Request() *http.Request {
	return c.request
}

// Tx ..
func (c *context) Tx() *sqlx.Tx {
	return c.tx
}

// SetContext ..
func (c *context) SetContext(ctx cont.Context) {
	c.context = ctx
}

// SetCurUser ..
func (c *context) SetCurUser(curUser *api.User) {
	c.curUser = curUser
}

// SetElasticQuery ..
func (c *context) SetElasticQuery(q *api.ElasticQuery) {
	c.elasticQuery = q
}

// SetLocale ..
func (c *context) SetLocale(locale string) {
	c.locale = locale
}

// SetRequest ..
func (c *context) SetRequest(r *http.Request) {
	c.request = r
}

// SetParams ..
func (c *context) SetParams(p httprouter.Params) {
	c.params = p
}

// SetTx ..
func (c *context) SetTx(tx *sqlx.Tx) {
	c.tx = tx
}
