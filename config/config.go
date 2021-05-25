package config

import (
	"bytes"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

//default env variables
const (
	EnvActiveProfiles       = "ACTIVE_PROFILES"
	EnvSpringCloudConfigUri = "SPRING_CLOUD_CONFIG_URI"
	EnvConfigPath           = "CONFIG_PATH"
	EnvAppName              = "APP_NAME"
)

var (
	once              sync.Once
	applicationConfig ApplicationConfig
	re                = regexp.MustCompile(`\$\{([A-Za-z0-9_-]+)((:)(.+?))?\}`)
)

type DecoderConfigOption func(*mapstructure.DecoderConfig)

type ApplicationConfig struct {
	viper *viper.Viper
}

func loadApplicationConfig() ApplicationConfig {
	v := viper.New()
	v.SetEnvPrefix("")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetDefault(EnvActiveProfiles, "")
	v.AutomaticEnv()
	v.SetConfigType("yaml")
	loadFile(v)
	loadSpringCloudConfig(v)
	v.AutomaticEnv()
	replaceConfigurationVariables(v)
	return ApplicationConfig{
		viper: v,
	}
}

func replaceConfigurationVariables(v *viper.Viper) {
	for _, k := range v.AllKeys() {
		value := v.GetString(k)
		groups := re.FindAllStringSubmatch(value, -1)
		for _, group := range groups {
			match := group[0]
			variable := group[1]
			defaultValue := group[4]
			variableValue := v.GetString(variable)
			if variableValue == "" {
				variableValue = defaultValue
			}

			value = strings.Replace(value, match, variableValue, -1)
		}
		v.Set(k, value)

	}
}

//loadFile loads the configuration from a file
func loadFile(v *viper.Viper) {
	v.SetDefault(EnvConfigPath, "./resources/")
	configPath := filepath.Dir(v.GetString(EnvConfigPath))
	v.SetConfigName("application")
	v.AddConfigPath(configPath)
	v.AddConfigPath("./config")
	v.AddConfigPath(".")
	err := v.MergeInConfig()
	if err != nil {
		fmt.Printf("unable to load config file application.yaml, error: %s\n", err)
	}
	profiles := v.GetString(EnvActiveProfiles)
	for _, p := range strings.Split(profiles, ",") {
		if p == "" {
			continue
		}
		v.SetConfigName(fmt.Sprintf("application-%s", p))
		err = v.MergeInConfig()
		if err != nil {
			fmt.Printf("unable to load config file application-%s.yaml, error: %s\n", p, err)
		}
	}
}

//loadSpringCloudConfig loads the configuration from a spring cloud configuration server
func loadSpringCloudConfig(v *viper.Viper) {
	if !v.IsSet(EnvSpringCloudConfigUri) {
		return
	}

	v.SetDefault(EnvAppName, "app")

	serverURL := v.GetString(EnvSpringCloudConfigUri)
	profiles := v.GetString(EnvActiveProfiles)
	appName := v.GetString(EnvAppName)
	if profiles != "" {
		profiles = "-" + profiles
	}
	url := fmt.Sprintf("%s/%s%s.yaml", serverURL, appName, profiles)
	v.SetConfigName("remote")
	resp := fetchConfiguration(url)
	err := v.MergeConfig(bytes.NewBuffer(resp))
	if err != nil {
		fmt.Printf("could not read config at :%s, viper error: %v.\n", url, err)
	}
}

func toViperOpts(opts []DecoderConfigOption) []viper.DecoderConfigOption {
	var viperOpts []viper.DecoderConfigOption
	for _, vo := range opts {
		viperOpts = append(viperOpts, viper.DecoderConfigOption(vo))
	}
	return viperOpts
}

func (c ApplicationConfig) Unmarshal(rawVal interface{}, opts ...DecoderConfigOption) error {
	return c.viper.Unmarshal(rawVal, toViperOpts(opts)...)
}

func (c ApplicationConfig) UnmarshalKey(key string, rawVal interface{}, opts ...DecoderConfigOption) error {
	return c.viper.UnmarshalKey(key, rawVal, toViperOpts(opts)...)
}

func (c ApplicationConfig) Settings() map[string]interface{} {
	return c.viper.AllSettings()
}

func (c ApplicationConfig) GetString(key string) string  {
	return c.viper.GetString(key)
}

func (c ApplicationConfig) Get(key string) interface{}  {
	return c.viper.Get(key)
}

func Config() ApplicationConfig {
	once.Do(func() {
		applicationConfig = loadApplicationConfig()
	})
	return applicationConfig
}
