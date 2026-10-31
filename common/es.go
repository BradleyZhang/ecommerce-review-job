package common

var (
	ESAddresses []string
	ESAPIKey    string
	Index       = "review"
)

func initESEnv() {
	ESAddresses = GetEnvOrDefaultStrings("ES_ADDRESSES", ESAddresses)
	ESAPIKey = GetEnvOrDefaultString("ES_API_KEY", ESAPIKey)
}
