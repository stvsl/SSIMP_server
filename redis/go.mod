module stvsljl.com/SSIMP/redis

require github.com/go-redis/redis/v8 v8.11.5

require stvsljl.com/SSIMP/utils v0.0.0

replace stvsljl.com/SSIMP/utils v0.0.0 => ../utils

require (
	github.com/antonfisher/nested-logrus-formatter v1.3.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

go 1.20
