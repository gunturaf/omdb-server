package mysqldb

type MysqlDB interface {
	SaveSearchActivity(searchWord string) error
}
