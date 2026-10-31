package common

var (
	KafkaSeeds     []string
	ConsumerGroup  = "review-job"
	ConsumerTopics = []string{"review"}
)

func initKafkaEnv() {
	KafkaSeeds = GetEnvOrDefaultStrings("KAFKA_SEEDS", KafkaSeeds)
}
