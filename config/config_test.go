package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"runtime"
	"testing"
)

type TestConfiguration struct {
	Dummy Dummy
}

type Dummy struct {
	Value       string
	SecondValue int
}

func init() {
	_, filename, _, _ := runtime.Caller(0)
	os.Setenv(EnvConfigPath, path.Join(path.Dir(filename), "..")+"/test-resources/")
	os.Setenv(EnvSpringCloudConfigUri, "http://localhost:8888/config")
	os.Setenv("DB_HOST", "127.0.0.1")
}

func TestLoadConfiguration(t *testing.T) {
	config := Config()

	if len(config.viper.AllKeys()) == 0 {
		t.Errorf("configuration is empty")
	}
}

func TestUnmarshallConfiguration(t *testing.T) {
	config := Config()

	testConfig := TestConfiguration{}
	err := config.Unmarshal(&testConfig)

	assert.Nil(t, err, "error unmarshalling configuration")
	assert.Equal(t, "hello default profile", testConfig.Dummy.Value)
	assert.Equal(t, 1, testConfig.Dummy.SecondValue)
}

func TestUnmarshallKeyConfiguration(t *testing.T) {
	config := Config()

	dummy := Dummy{}
	err := config.UnmarshalKey("dummy", &dummy)

	assert.Nil(t, err, "error unmarshalling configuration")
	assert.Equal(t, "hello default profile", dummy.Value)
	assert.Equal(t, 1, dummy.SecondValue)
}

func TestUnmarshallConfigurationSeveralProfiles(t *testing.T) {
	os.Setenv(EnvActiveProfiles, "dev,local")

	config := loadApplicationConfig()

	testConfig := TestConfiguration{}
	err := config.Unmarshal(&testConfig)

	assert.Nil(t, err, "error unmarshalling configuration")
	assert.Equal(t, "hello local profile", testConfig.Dummy.Value)
	assert.Equal(t, 1, testConfig.Dummy.SecondValue)
}

func TestUnmarshallConfigurationApplyReplaceEnv(t *testing.T) {
	config := loadApplicationConfig()

	v := config.GetString("datasource.default.connectionString")

	assert.NotNil(t, v)
	assert.Equal(t, "root:password@tcp(127.0.0.1:3306)/localdb?multiStatements=true&parseTime=true", v)
}
