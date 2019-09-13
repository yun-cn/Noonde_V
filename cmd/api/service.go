package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/yun313350095/Noonde/api/conf"
	"github.com/yun313350095/Noonde/api/elastic"
	"github.com/yun313350095/Noonde/api/http"
	"github.com/yun313350095/Noonde/api/job"
	"github.com/yun313350095/Noonde/api/log"
	"github.com/yun313350095/Noonde/api/mail"
	"github.com/yun313350095/Noonde/api/mysql"
	"github.com/yun313350095/Noonde/api/s3"
	"github.com/yun313350095/Noonde/api/www"
)

var (
	confsv    = &conf.Service{}
	elasticsv = &elastic.Service{}
	httpsv    = &http.Service{}
	jobsv     = &job.Service{}
	logsv     = &log.Service{}
	mailsv    = &mail.Service{}
	mysqlsv   = &mysql.Service{}
	s3sv      = &s3.Service{}
	wwwsv     = &www.Service{}
)

type service struct {
	Conf    *conf.Service
	Elastic *elastic.Service
	HTTP    *http.Service
	Job     *job.Service
	Log     *log.Service
	Mail    *mail.Service
	MySQL   *mysql.Service
	S3      *s3.Service
	WWW     *www.Service
}

func newService() *service {
	var err error

	elasticsv.Conf = confsv
	elasticsv.Log = logsv

	httpsv.Conf = confsv
	httpsv.Log = logsv
	httpsv.Elastic = elasticsv
	httpsv.Job = jobsv
	httpsv.WWW = wwwsv

	jobsv.Conf = confsv
	jobsv.Log = logsv

	logsv.Conf = confsv

	mailsv.Conf = confsv
	mailsv.Log = logsv

	mysqlsv.Conf = confsv
	mysqlsv.Log = logsv

	wwwsv.HTTP = httpsv
	wwwsv.Conf = confsv
	wwwsv.Log = logsv
	wwwsv.Mail = mailsv
	wwwsv.S3 = s3sv
	wwwsv.MySQL = mysqlsv
	wwwsv.Job = jobsv

	s3sv.Conf = confsv
	s3sv.Log = logsv

	confsv.Log = logsv

	// conf
	if err = confsv.Load(); err != nil {
		panic(err)
	}

	// elastic
	if err = elasticsv.Load(); err != nil {
		panic(err)
	}

	// http
	httpsv.Load()
	httpsv.SetMiddlewares()
	httpsv.SetRoutes()

	// job
	jobsv.Load()

	// log
	logsv.Load()

	// mail
	if err := mailsv.Load(); err != nil {
		panic(err)
	}

	// mysql
	if err := mysqlsv.Load(); err != nil {
		panic(err)
	}

	// S3
	if err := s3sv.Load(); err != nil {
		panic(err)
	}

	return &service{
		Conf:    confsv,
		Elastic: elasticsv,
		HTTP:    httpsv,
		Job:     jobsv,
		Log:     logsv,
		Mail:    mailsv,
		MySQL:   mysqlsv,
		S3:      s3sv,
		WWW:     wwwsv,
	}
}
