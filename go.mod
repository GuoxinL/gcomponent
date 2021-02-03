module github.com/GuoxinL/gcomponent

go 1.13

require (
	github.com/gin-contrib/zap v0.0.1
	github.com/gin-gonic/gin v1.6.3
	github.com/gobike/envflag v0.0.0-20160830095501-ae3268980a29
	github.com/google/uuid v1.1.2
	github.com/mitchellh/mapstructure v1.2.2 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/onsi/ginkgo v1.14.2 // indirect
	github.com/onsi/gomega v1.10.3 // indirect
	github.com/pelletier/go-toml v1.7.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.5.1 // indirect
	github.com/valyala/fasthttp v1.9.0
	go.uber.org/atomic v1.7.0
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.15.0
	gopkg.in/ini.v1 v1.55.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/redis.v5 v5.2.9
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.8
	moul.io/zapgorm2 v1.0.1
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20200406173513-056763e48d71
	golang.org/x/net => github.com/golang/net v0.0.0-20200324143707-d3edc9973b7e
	//golang.org/x/sync => github.com/golang/sync v0.0.0-20200317015054-43a5402ce75a
	golang.org/x/sys => github.com/golang/sys v0.0.0-20200409092240-59c9f1ba88fa
	golang.org/x/text => github.com/golang/text v0.3.2
//google.golang.org/protobuf => github.com/golang/protobuf v1.4.2
)
