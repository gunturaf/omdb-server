package mysqldb

import "context"

type MysqlDB interface {
	SaveSearchActivity(ctx context.Context, searchWord string) error
}
