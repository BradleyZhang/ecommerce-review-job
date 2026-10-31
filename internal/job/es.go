package job

import (
	"context"
	"encoding/json"

	"github.com/BradleyZhang/ecommerce-review-job/common"
	es "github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/core/update"
)

type ESClient struct {
	*es.TypedClient
	index string
}

type ESCfg struct {
	Addresses []string
	Index     string
	Apikey    string
}

func NewESClient(cfg ESCfg) (*ESClient, error) {
	var opts []es.Option
	opts = append(opts,
		es.WithAddresses(cfg.Addresses...),
		es.WithAPIKey(cfg.Apikey),
	)

	client, err := es.NewTyped(opts...)
	if err != nil {
		return nil, err
	}
	common.SysLogf("connect to es with address: %v", cfg.Addresses)
	common.DebugLogf("[job] NewESClient cfg:%v", cfg)
	return &ESClient{
		TypedClient: client,
		index:       cfg.Index,
	}, nil
}
func NewESDefaultCfg() ESCfg {
	return ESCfg{
		Addresses: common.ESAddresses,
		Index:     common.Index,
		Apikey:    common.ESAPIKey,
	}
}

//TODO 下面理解

func (jw JobWorker) indexDocument(d map[string]any) {
	reviewID := d["review_id"].(string)
	resp, err := jw.esClient.Index(jw.esClient.index).
		Id(reviewID).
		Document(d).
		Do(context.Background())
	if err != nil {
		common.ErrorLogf("indexing document failed, err:%v data:%v\n", err, d)
		return
	}
	common.DebugLogf("result:%#v\n", resp.Result)
}

func (jw JobWorker) updateDocument(d map[string]interface{}) {
	reviewID := d["review_id"].(string)
	body, err := json.Marshal(d)
	if err != nil {
		common.ErrorLogf("failed to marshal data:%v  , err:%v\n", d, err)
		return
	}
	resp, err := jw.esClient.Update(jw.esClient.index, reviewID).Request(&update.Request{
		Doc: json.RawMessage(body),
	}).Do(context.Background())
	if err != nil {
		common.ErrorLogf("update document failed, err:%v\n", err)
		return
	}
	common.DebugLogf("result:%v\n", resp.Result)
}
