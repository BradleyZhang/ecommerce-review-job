package job

import (
	"os"

	"github.com/BradleyZhang/ecommerce-review-job/common"
	"github.com/twmb/franz-go/pkg/kgo"
)

type KgoConfig struct {
	Seeds          []string
	ConsumerGroup  string
	ConsumerTopics []string
}

func NewKgoCl(cfg KgoConfig) (*kgo.Client, error) {
	var opts []kgo.Opt
	common.SysLogf("New Kafka Client with seeds:%v ,Consumer Group:%v ,Consume Topics:%v", cfg.Seeds, cfg.ConsumerGroup, cfg.ConsumerTopics)
	opts = append(opts,
		kgo.SeedBrokers(cfg.Seeds...),
		kgo.ConsumerGroup(cfg.ConsumerGroup),
		kgo.ConsumeTopics(cfg.ConsumerTopics...),
	)
	if common.DebugMode {
		opts = append(opts, kgo.WithLogger(kgo.BasicLogger(os.Stderr, kgo.LogLevelDebug, func() string { return "[FRANZ_GO] " })))
	}
	return kgo.NewClient(opts...)
}
func NewDefaultKgoCfg() KgoConfig {
	return KgoConfig{
		Seeds:          common.KafkaSeeds,
		ConsumerGroup:  common.ConsumerGroup,
		ConsumerTopics: common.ConsumerTopics,
	}
}
