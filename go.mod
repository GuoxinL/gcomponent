module github.com/GuoxinL/gcomponent

go 1.13

require (
	github.com/fasthttp/router v1.0.1
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gobike/envflag v0.0.0-20160830095501-ae3268980a29
	github.com/gomodule/redigo v1.8.1
	github.com/jinzhu/gorm v1.9.12
	github.com/mitchellh/mapstructure v1.2.2 // indirect
	github.com/mna/redisc v1.1.7
	github.com/pelletier/go-toml v1.7.0 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.6.3
	github.com/valyala/fasthttp v1.9.0
	go.uber.org/atomic v1.4.0
	golang.org/x/crypto v0.0.0-20200221231518-2aa609cf4a9d // indirect
	golang.org/x/sys v0.0.0-20200409092240-59c9f1ba88fa // indirect
	golang.org/x/text v0.3.2 // indirect
	gopkg.in/ini.v1 v1.55.0 // indirect
)
replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20200406173513-056763e48d71
	golang.org/x/net => github.com/golang/net v0.0.0-20200324143707-d3edc9973b7e
	golang.org/x/sync => github.com/golang/sync v0.0.0-20200317015054-43a5402ce75a
	golang.org/x/sys => github.com/golang/sys v0.0.0-20200409092240-59c9f1ba88fa
	golang.org/x/text => github.com/golang/text v0.3.2
	google.golang.org/protobuf => github.com/golang/protobuf v1.4.2
)
