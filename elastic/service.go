package elastic

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/jmoiron/sqlx"
	"github.com/olivere/elastic"
	aws "github.com/olivere/elastic/aws/v4"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"github.com/yun313350095/Noonde/api"
	"strings"
)

// Service ..
type Service struct {
	Log    api.LogService
	Conf   api.ConfService
	S3     api.S3Service
	client *elastic.Client
}

// AddTargets ..
func (s *Service) AddTargets(targets []*api.ElasticTarget, typeName string, ids []int64) []*api.ElasticTarget {
	var tt []*api.ElasticTarget

	for _, id := range ids {
		tt = append(tt, &api.ElasticTarget{
			Type: typeName,
			ID:   id,
		})
	}

	return append(targets, tt...)
}

// BulkUpdate ..
func (s *Service) BulkUpdate(ctx context.Context, tx *sqlx.Tx, targets []*api.ElasticTarget) error {
	var err error

	req := s.client.Bulk()

	var spaceIDs []int64

	for _, target := range targets {
		switch target.Type {
		case api.TypeSpace:
			spaceIDs = append(spaceIDs, target.ID)

		default:
			err = errors.New("No case for " + target.Type)
			s.Log.Error(err)
			return err
		}
	}

	req, err = s.updateSpaces(ctx, tx, req, spaceIDs)
	if err != nil {
		s.Log.Error(err)
		return err
	}

	if req.NumberOfActions() == 0 {
		return nil
	}

	res, err := req.Do(ctx)
	if err != nil {
		s.Log.Error(err)
		return err
	}

	if req.NumberOfActions() != 0 {
		err = errors.New("xxx")
		s.Log.Error(err)
		return err
	}

	failed := res.Failed()
	if len(failed) > 0 {
		var results []string
		for _, f := range failed {
			results = append(results, f.Id+": "+f.Result)
		}
		err = errors.New(strings.Join(results, " | "))
		s.Log.Error(err)
		return err
	}

	return nil
}

// Clear ..
func (s *Service) Clear(ctx context.Context, tname string) error {
	bquery := elastic.NewBoolQuery()

	bquery = bquery.Must(
		elastic.NewSimpleQueryStringQuery(tname).Field("type").DefaultOperator("AND"))

	_, err := elastic.NewDeleteByQueryService(s.client).
		Index(s.Conf.String("elastic.index")).
		Type(s.Conf.String("elastic.type")).
		Query(bquery).
		Do(ctx)
	if err != nil {
		s.Log.Error(err)
		return err
	}

	return nil
}

// Search ..
func (s *Service) Search(ctx context.Context, query *api.ElasticQuery) ([]int64, int64, error) {
	bquery := elastic.NewBoolQuery()

	if query.Filter != "" && query.Start != nil && query.End != nil {
		bquery.Filter(
			elastic.NewRangeQuery(query.Filter).
				From(query.Start).
				To(query.End))
	}

	if query.Type != "" {
		bquery = bquery.Must(
			elastic.NewSimpleQueryStringQuery(query.Type).Field("type").DefaultOperator("AND"))
	}

	if query.IDsQuery != "" {
		bquery = bquery.Must(
			elastic.NewSimpleQueryStringQuery(query.IDsQuery).Field("id").DefaultOperator("OR"))
	}

	if query.SearchQuery != "" {
		bquery = bquery.Must(
			elastic.NewSimpleQueryStringQuery(query.SearchQuery).Field("search").DefaultOperator("AND"))
	}

	if query.TagsQuery != "" {
		bquery = bquery.Must(
			elastic.NewSimpleQueryStringQuery(query.TagsQuery).Field("tags").DefaultOperator("AND"))
	}

	src := elastic.NewSearchSource().
		Query(bquery).
		From(query.From).
		Size(query.Size).
		Sort(query.Sort, query.Asc)

	// Debug
	sc, _ := src.Source()
	js, _ := json.Marshal(sc)
	s.Log.Debug(string(js))

	res, err := s.client.Search(s.Conf.String("elastic.index")).SearchSource(src).Do(ctx)
	if err != nil {
		s.Log.Error(err)
		return nil, 0, err
	}

	var ids []int64
	if res.Hits.TotalHits > 0 {
		for _, h := range res.Hits.Hits {
			bb, err := h.Source.MarshalJSON()
			if err != nil {
				s.Log.Error(err)
				return nil, 0, err
			}
			id := gjson.Get(string(bb), "id").Int()
			ids = append(ids, id)
		}
	}
	return ids, res.Hits.TotalHits, nil
}

// StopClient
func (s *Service) StopClient() {
	s.client.Stop()
}

//--------------------------------------------------------------------------------

// Load ..
func (s *Service) Load() error {
	var err error

	if s.Conf.String("goenv") != "local" {
		signingClient := aws.NewV4SigningClient(credentials.NewStaticCredentials(
			s.Conf.String("aws.access_key_id"),
			s.Conf.String("aws.secret_access_key"),
			"",
		), "ap-northeast-1")

		s.client, err = elastic.NewClient(
			elastic.SetURL(s.Conf.String("elastic.url")),
			elastic.SetSniff(false),
			elastic.SetHealthcheck(false),
			elastic.SetHttpClient(signingClient),
		)
		if err != nil {
			s.Log.Error(err)
			return err
		}
		return nil
	}

	s.client, err = elastic.NewClient(
		elastic.SetURL(s.Conf.String("elastic.url")),
	)
	if err != nil {
		s.Log.Error(err)
		return err
	}

	return nil
}
