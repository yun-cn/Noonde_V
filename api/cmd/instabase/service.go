package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/yun313350095/Noonde/api/conf"
	"github.com/yun313350095/Noonde/api/elastic"
	"github.com/yun313350095/Noonde/api/log"
	"github.com/yun313350095/Noonde/api/mysql"
	"github.com/yun313350095/Noonde/api/s3"
	"github.com/yun313350095/Noonde/api/scraper"
)

var (
	confsv    = &conf.Service{}
	elasticsv = &elastic.Service{}
	logsv     = &log.Service{}
	mysqlsv   = &mysql.Service{}
	scrapersv = &scraper.Service{}
	s3sv      = &s3.Service{}
)

type service struct {
	Conf    *conf.Service
	Elastic *elastic.Service
	Log     *log.Service
	MySQL   *mysql.Service
	Scraper *scraper.Service
	S3      *s3.Service
}

func newService() *service {
	var err error

	elasticsv.Conf = confsv
	elasticsv.Log = logsv
	elasticsv.S3 = s3sv

	logsv.Conf = confsv

	mysqlsv.Conf = confsv
	mysqlsv.Log = logsv

	scrapersv.Conf = confsv
	scrapersv.Log = logsv
	scrapersv.MySQL = mysqlsv

	confsv.Log = logsv

	// conf ..
	if err = confsv.Load(); err != nil {
		panic(err)
	}

	// elastic ..
	if err := elasticsv.Load(); err != nil {
		panic(err)
	}

	// log ..
	logsv.Load()

	// mysql ..
	if err := mysqlsv.Load(); err != nil {
		panic(err)
	}

	return &service{
		Conf:    confsv,
		Elastic: elasticsv,
		Log:     logsv,
		MySQL:   mysqlsv,
		Scraper: scrapersv,
	}
}
