package job

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

// 评价数据流处理
type JobWorker struct {
	kgoCl    *kgo.Client
	esClient *ESClient
}

func NewJobWorker(kgoCl *kgo.Client, esClient *ESClient) *JobWorker {
	return &JobWorker{
		kgoCl:    kgoCl,
		esClient: esClient,
	}
}

type Msg struct {
	Data     []map[string]any `json:"data"`
	Database string           `json:"database"`
	Table    string           `json:"table"`
	IsDdl    bool             `json:"isDdl"`
	Type     string           `json:"type"`
}

func (j *JobWorker) Start(ctx context.Context) error {
	for {
		// 从 kafka 获取 mysql 的数据变更
		fetches := j.kgoCl.PollFetches(ctx)
		if errs := fetches.Errors(); len(errs) > 0 {
			return fmt.Errorf("fetching message from kafka failed,errs: %v", errs)
			// continue
		}
		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			msg := &Msg{}
			if err := json.Unmarshal(record.Value, msg); err != nil {
				record.Ack(kgo.AckRelease) // TODO 这样错误就会一直流窜，怎么处理好
				return fmt.Errorf("unmarshal msg from kafka failed ,err:%v msg: %v", err, record)
			}

			if msg.Type == "INSERT" {
				// es 中新增 document
				for idx := range msg.Data {
					j.indexDocument(msg.Data[idx])
				}
			} else {
				// es 中更新 document
				for idx := range msg.Data {
					j.updateDocument(msg.Data[idx])
				}

			}

			record.Ack(kgo.AckAccept)
		}
	}
}
func (j *JobWorker) Stop() {
	j.kgoCl.Close()
}
