package es

import (
	"errors"
	"fmt"
	"time"
)

const (
	ESIndex = "nids"
)

type UnsupportedVersion struct{}

func (UnsupportedVersion) Error() string {
	return "Unsupported ElasticSearch Client Version"
}

type elasticWrapper interface {
	IndexExists(indices ...string) (bool, error)
	CreateIndex(name string) (bool, error)
	AddAlias(index string, alias string) (bool, error)
	HasAlias(index string, alias string) (bool, error)
	AddBulkReq(index, typeName string, data interface{}) error
	FlushBulk() error
}

type ElasticConfig struct {
	Url         []string
	User        string
	Secret      string
	MaxRetries  *int
	HealthCheck *bool
	Timeout     *time.Duration
	Sniff       *bool
}

type ElasticSearchService struct {
	EsClient  elasticWrapper
	baseIndex string
}

func (esSvc *ElasticSearchService) Index(date time.Time, namespace string) string {
	dateStr := date.Format("2006.01.02")
	if len(namespace) > 0 {
		return fmt.Sprintf("%s-%s-%s", esSvc.baseIndex, namespace, dateStr)
	}
	return fmt.Sprintf("%s-%s", esSvc.baseIndex, dateStr)
}

func (esSvc *ElasticSearchService) IndexAlias(typeName string) string {
	return fmt.Sprintf("%s-%s", esSvc.baseIndex, typeName)
}

func (esSvc *ElasticSearchService) FlushData() error {
	return esSvc.EsClient.FlushBulk()
}

// SaveDataIntoES save metrics and events to ES by using ES client
func (esSvc *ElasticSearchService) SaveData(date time.Time, typeName string, namespace string, sinkData []interface{}) error {
	if typeName == "" || len(sinkData) == 0 {
		return nil
	}

	indexName := esSvc.Index(date, namespace)

	// Use the IndexExists service to check if a specified index exists.
	exists, err := esSvc.EsClient.IndexExists(indexName)
	if err != nil {
		return err
	}

	if !exists {

		ack, err := esSvc.EsClient.CreateIndex(indexName)
		if err != nil {
			return err
		}
		if !ack {
			return errors.New("Failed to acknoledge index creation")
		}
	}

	aliasName := esSvc.IndexAlias(typeName)

	hasAlias, err := esSvc.EsClient.HasAlias(indexName, aliasName)
	if err != nil {
		return err
	}

	if !hasAlias {
		ack, err := esSvc.EsClient.AddAlias(indexName, esSvc.IndexAlias(typeName))
		if err != nil {
			return err
		}

		if !ack {
			return errors.New("Failed to acknoledge index alias creation")
		}
	}

	for _, data := range sinkData {
		esSvc.EsClient.AddBulkReq(indexName, typeName, data)
	}

	return nil
}

func CreateElasticSearchService(config ElasticConfig, version int) (*ElasticSearchService, error) {

	var esSvc ElasticSearchService
	esSvc.baseIndex = ESIndex

	bulkWorkers := 5
	pipeline := ""

	switch version {
	case 6:
		esSvc.EsClient, err = NewEsClient6(config, bulkWorkers, pipeline)
	case 7:
		esSvc.EsClient, err = NewEsClient7(config, bulkWorkers, pipeline)
	default:
		return nil, UnsupportedVersion{}
	}
	if err != nil {
		return nil, fmt.Errorf("Failed to create ElasticSearch client: %v", err)
	}

	return &esSvc, nil
}
