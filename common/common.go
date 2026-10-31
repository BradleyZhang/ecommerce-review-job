package common

var (
	DebugMode = false
)

func initCommonEnv() {
	DebugMode = GetEnvOrDefaultBool("DEBUG_MODE", DebugMode)
}
