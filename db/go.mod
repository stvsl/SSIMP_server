module stvsljl.com/SSIMP/db

require stvsljl.com/SSIMP/utils v0.0.0

require (
	gorm.io/datatypes v1.0.7
	gorm.io/gorm v1.23.8
	gorm.io/driver/mysql v1.3.5
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace stvsljl.com/SSIMP/utils v0.0.0 => ../utils

go 1.18
