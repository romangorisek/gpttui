package config

type KeymapSetting struct {
	Key      string
	Modifier string
}

type UserConfig struct {
	KeybindingOverrides map[string]KeymapSetting
}

func LoadUserConfig() *UserConfig {
	overrides := map[string]KeymapSetting{ // TODO: get overried in from config
		"global_nextView": {Key: "<arrow_up>"},
	}

	return &UserConfig{
		KeybindingOverrides: overrides,
	}
}
