module stvsljl.com/SSIMP/db

require stvsljl.com/SSIMP/utils v0.0.0

require (
	gorm.io/datatypes v1.1.1
	gorm.io/driver/mysql v1.4.7
	gorm.io/gorm v1.24.6
)

require (
	github.com/antonfisher/nested-logrus-formatter v1.3.1 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace stvsljl.com/SSIMP/utils v0.0.0 => ../utils

go 1.20
