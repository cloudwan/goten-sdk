module github.com/cloudwan/goten-sdk

go 1.16

require (
	github.com/alecthomas/participle v0.5.0
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/iancoleman/strcase v0.0.0-20180726023541-3605ed457bf7
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/olekukonko/tablewriter v0.0.4
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.26.0
)

replace google.golang.org/protobuf => github.com/cloudwan/goten-protobuf v1.26.0
