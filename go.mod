module github.com/hrdkgmz/dbWrapper

go 1.12

require (
	dm.com v0.0.0
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gomodule/redigo v2.0.0+incompatible
)

replace dm.com v0.0.0 => ./dm
