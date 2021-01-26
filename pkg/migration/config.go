package migration

import (
	"github.com/ovargasmahisoft/kmn-commons/pkg/config"
	"sync"
)

type DatasourceConfig map[string]Datasource

type Datasource struct {
	ConnectionString string
	DriverName       string
	MigrationPath    string
}

var (
	once             sync.Once
	datasourceConfig DatasourceConfig
)

func Config() DatasourceConfig {
	once.Do(func() {
		ds := &map[string]Datasource{}
		config.Config().UnmarshalKey("datasource", ds)
		datasourceConfig = *ds
	})
	return datasourceConfig
}
