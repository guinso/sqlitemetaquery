package sqlitemetaquery

import (
	"github.com/guinso/rdbmstool"
)

//SQLITEMetaQuery get SQLITE data table meta information
type SQLITEMetaQuery struct {
}

//GetTableNames get list of datatables' name which start with provided search pattern
//search pattern allow '%' as wild card; example 'hub_%'
func (meta *SQLITEMetaQuery) GetTableNames(
	db rdbmstool.DbHandlerProxy, databaseName string, tableNamePattern string) ([]string, error) {
	rows, err := db.Query("SELECT name FROM " + databaseName + ".sqlite_master" +
		" where type='table' AND name LIKE '" + tableNamePattern + "'")

	if err != nil {
		return nil, err
	}

	var result []string
	for rows.Next() {
		var tmp string
		err := rows.Scan(&tmp)

		if err != nil {
			continue
		}

		result = append(result, tmp)
	}

	rows.Close()

	return result, nil
}
