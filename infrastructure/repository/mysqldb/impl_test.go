package mysqldb_test

import (
	"database/sql"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/gunturaf/omdb-server/infrastructure/repository"
	"github.com/gunturaf/omdb-server/infrastructure/repository/mysqldb"
)

var _ = Describe("Impl", func() {

	Describe("MysqlDBDSL", func() {
		Context("GetDSN", func() {
			It("correct", func() {
				dsn := mysqldb.MysqlDBDSL{
					Username: "sukab",
					Password: "world",
					Host:     "localhost",
					Port:     "3306",
					DBName:   "sukab",
				}

				str := dsn.GetDSN()
				Expect(str).To(Equal("sukab:world@tcp(localhost:3306)/sukab"))
			})
		})
	})

	Describe("MysqlDBImpl", func() {

		db := repository.NewMockDB()

		Context("no error", func() {
			It("ok", func() {
				db.MockQuery = func(query string, args ...interface{}) (*sql.Rows, error) {
					return &sql.Rows{}, nil
				}
				op := mysqldb.NewMysqlDB(db)

				err := op.SaveSearchActivity("Batman")
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("any error", func() {
			It("error", func() {
				db.MockQuery = func(query string, args ...interface{}) (*sql.Rows, error) {
					return nil, sql.ErrNoRows
				}
				op := mysqldb.NewMysqlDB(db)

				err := op.SaveSearchActivity("Batman")
				Expect(err).To(HaveOccurred())
			})
		})
	})

})
