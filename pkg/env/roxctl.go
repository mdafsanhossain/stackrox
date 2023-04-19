package env

var (
	// EndpointEnv specifies the central endpoint to use for commandline operations.
	EndpointEnv = RegisterSetting("ROX_ENDPOINT")

	// PasswordEnv specifies the central admin password to use for commandline operations.
	PasswordEnv = RegisterSetting("ROX_ADMIN_PASSWORD")

	// TokenEnv is the variable that clients can source for commandline operations.
	TokenEnv = RegisterSetting("ROX_API_TOKEN")

	// ConfigDirEnv is the variable that clients can use for specifying the config location for commandline operations.
	ConfigDirEnv = RegisterSetting("ROX_CONFIG_DIR")
)
