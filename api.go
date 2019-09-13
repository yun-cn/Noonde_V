package api

import (
	"context"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"time"
)

//--------------------------------------------------------------------------------
// conf

// ConfService ..
type ConfService interface {

	// Bool ..
	Bool(key string) bool

	// Int ..
	Int(key string) int

	// Int64 ..
	Int64(key string) int64

	// String ..
	String(key string) string

	// StringSlice ..
	StringSlice(key string) []string
}

//--------------------------------------------------------------------------------
// elastic

// ElasticService ..
type ElasticService interface {
	// AddTargets ..
	AddTargets(targets []*ElasticTarget, typeName string, ids []int64) []*ElasticTarget

	// BulkUpdate ..
	BulkUpdate(ctx context.Context, tx *sqlx.Tx, targets []*ElasticTarget) error

	// Clear ..
	Clear(ctx context.Context, tname string) error

	// Search ..
	Search(ctx context.Context, query *ElasticQuery) ([]int64, int64, error)

	// StoClient ..
	StopClient()
}

// ElasticTarget ..
type ElasticTarget struct {
	Type string
	ID   int64
}

// ElasticQuery ..
type ElasticQuery struct {
	Type        string
	IDsQuery    string
	SearchQuery string
	TagsQuery   string
	Filter      string
	Start       interface{}
	End         interface{}
	From        int
	Size        int
	Sort        string
	Asc         bool
}

//--------------------------------------------------------------------------------
// log

// LogService ..
type LogService interface {

	// Debug ..
	Debug(args ...interface{})

	// Info ..
	Info(args ...interface{})

	// Warn ..
	Warn(args ...interface{})

	// Error ..
	Error(args ...interface{})

	// Fatal ..
	Fatal(args ...interface{})

	// Panic ..
	Panic(args ...interface{})

	// ErrorWithStacktrace ..
	ErrorWithStacktrace(info interface{})
}

//--------------------------------------------------------------------------------
// mail

// MailService ..
type MailService interface {

	// Send ..
	Send(
		name string,
		data interface{},
		locale string,
		dests []string,
		from *MailPerson,
		to *MailPerson,
	) error
}

// MailPerson ..
type MailPerson struct {
	Name string
	Addr string
}

//--------------------------------------------------------------------------------
// mysql

// MySQLService ..
type MySQLService interface {

	// Begin ..
	Begin(ctx context.Context) (*sqlx.Tx, error)

	// Deadlocked ..
	Deadlocked(err error) bool

	// Format ..
	Format(t time.Time) string

	// MaxTime ..
	MaxTime() *time.Time

	// MinTime ..
	MinTime() *time.Time

	// Reader ..
	Reader() *sqlx.DB

	// SafeTime ..
	SafeTime(t time.Time) time.Time

	// SumUp ..
	SumUp(
		ctx context.Context,
		tx *sqlx.Tx,
		ofType string,
		forType string,
		forIDs []int64,
	) error

	// Writer ..
	Writer() *sqlx.DB
}

//--------------------------------------------------------------------------------
// s3

// S3Service ..
type S3Service interface {

	// Delete ..
	Delete(path string) error

	// Presign ..
	Presign(path string) (string, error)

	// Upload ..
	Upload(path string, body io.Reader) error
}

//--------------------------------------------------------------------------------
// http

// HTTPService ..
type HTTPService interface {

	// NewContext ..
	NewContext(ctx context.Context) Context

	// NewParamError ..
	NewParamError() ParamError

	// Router ..
	Router() *httprouter.Router

	// Upgrader ..
	//Upgrader() *websocket.Upgrader

	// WriteError ..
	WriteError(w http.ResponseWriter, code string, err error)
}

// ParamError ..
type ParamError interface {

	// Error ..
	Error() string

	// IsNotEmpty ..
	IsNotEmpty() bool

	// PushIfNotExists ..
	PushIfNotExists(key string, code string)
}

// Context ..
type Context interface {

	// Context ..
	Context() context.Context

	// CurUser ..
	CurUser() *User

	// ElasticQuery ..
	ElasticQuery() *ElasticQuery

	// Locale ..
	Locale() string

	// Params ..
	Params() httprouter.Params

	// Request ..
	Request() *http.Request

	// Tx ..
	Tx() *sqlx.Tx

	// SetContext ..
	SetContext(ctx context.Context)

	// SetCurUser ..
	SetCurUser(curUser *User)

	// SetElasticQuery ..
	SetElasticQuery(q *ElasticQuery)

	// SetLocale ..
	SetLocale(locale string)

	// SetParams ..
	SetParams(p httprouter.Params)

	// SetRequest ..
	SetRequest(r *http.Request)

	// SetTx ..
	SetTx(tx *sqlx.Tx)
}

//--------------------------------------------------------------------------------
// job

// JobService ..
type JobService interface {

	// Add ..
	Add(job Job)

	// Count ..
	Count() int32

	// SetInShutdown ..
	SetInShutdown()

	// StartWorkers ..
	StartWorkers()
}

// Job ..
type Job interface {

	// Perform ..
	Perform()
}

//--------------------------------------------------------------------------------
// www

// WWWService ..
type WWWService interface {

	// SetRoutes ..
	SetRoutes()
}

//--------------------------------------------------------------------------------
// mqtt

// MQTTService ..
type MQTTService interface {

	// Connect ..
	Connect(handler func(c MQTTClient)) (MQTTClient, error)
}

// MQTTClient ..
type MQTTClient interface {

	// Disconnect ..
	Disconnect()

	// IsConnected ..
	IsConnected()

	// Publish ..
	Publish(topic string, payload interface{})

	// Subscribe ..
	Subscribe(topic string, handler mqtt.MessageHandler) error
}

//--------------------------------------------------------------------------------
// ws

// WSService ..
type WSService interface {

	// ChangedCols ..
	ChangedCols(typeName string, oldData interface{}, newData interface{}, moreIDs [][]int64) ([]string, error)

	// PublishByHub ..
	PublishByHub(hubKey string, subKey string, payload interface{}) error

	// PublishNotif ..
	PublishNotif(
		typeName string,
		actionName string,
		key string,
		user interface{},
		meta interface{},
	) error

	// Read ..
	Read(
		hubKey string,
		subKey string,
		client *WSClient,
		handler func(
			client *WSClient,
			message []byte,
			args ...interface{},
		) error,
		args ...interface{},
	)

	// Register ..
	Register(hubKey string, subKey string, client *WSClient)

	// StopHubs ..
	StopHubs()

	// SubscribeIfNotConnected ..
	SubscribeIfNotConnected(hubKey string, subKey string)

	// Write ..
	Write(
		client *WSClient,
		handler func(
			client *WSClient,
			message []byte,
			args ...interface{},
		) error,
		args ...interface{},
	)
}

// WSClient ..
type WSClient struct {
	Closed int32
	Conn   *websocket.Conn
	Send   chan []byte
}

// Hub ..
const (
	HubNotif   = "notif"
	HubViewers = "viewers"
)

// Action ..
const (
	ActionCreated = "created"
	ActionDeleted = "deleted"
	ActionUpdated = "updated"
)

// NotifKeyAny ..
const NotifKeyAny = "any"

//--------------------------------------------------------------------------------
// scraper

// ScraperService ..
type ScraperService interface {
	// NewSpaceMarketClient ..
	NewSpaceMarketClient() (SpaceMarketClient, error)

	// NewInstabaseClient ..
	NewInstabaseClient() (InstabaseClient, error)
}

// SpaceMarketClient ..
type SpaceMarketClient interface {

	// RefreshProxyClient ..
	RefreshProxyClient() error

	// SearchRoomList ..
	SearchRoomListHourly(int, int, string, string) (string, error)

	// GetRoomDetails ..
	GetRoomDetails(string, string) (string, error)
}

// InstabaseClient ..
type InstabaseClient interface {

	// RefreshProxyClient ..
	RefreshProxyClient() error

	// SearchSpaceList ..
	SearchSpaceList(int, int) (string, error)

	// Get space details
	GetSpaceDetails(int64) (string, error)
}
