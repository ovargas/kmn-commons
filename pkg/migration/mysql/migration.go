package mysql

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ovargasmahisoft/kmn-commons/pkg/migration"
)

func MigrateMySql(datasourceName string) {
	datasource, ok := migration.Config()[datasourceName]

	if !ok {
		panic(fmt.Errorf("missing configuration [datasource.%s]", datasourceName))
	}

	MigrateMySqlWithDatasource(datasource)
}

func MigrateMySqlWithDatasource(datasource migration.Datasource) {
	if db, err := sql.Open(datasource.DriverName, datasource.ConnectionString); err == nil {
		defer db.Close()
		driver, err := mysql.WithInstance(db, &mysql.Config{})

		if err != nil {
			panic(err)
		}

		m, err := migrate.NewWithDatabaseInstance(
			fmt.Sprintf("file://%s", datasource.MigrationPath),
			"mysql",
			driver,
		)

		if err != nil {
			panic(err)
		}

		err = m.Up()

		if err != nil && err != migrate.ErrNoChange {
			panic(err)
		}
	} else {
		panic(err)
	}
}
