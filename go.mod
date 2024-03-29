module github.com/cloudwan/goten-sdk

go 1.21

require (
	github.com/alecthomas/participle v0.5.0
	github.com/google/cel-go v0.20.1
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/iancoleman/strcase v0.0.0-20180726023541-3605ed457bf7
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/olekukonko/tablewriter v0.0.4
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.1 // indirect
	github.com/mattn/go-runewidth v0.0.7 // indirect
	github.com/sirupsen/logrus v1.4.2 // indirect
	github.com/stoewer/go-strcase v1.2.0 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.9.0 // indirect
)

replace google.golang.org/protobuf => github.com/cloudwan/goten-protobuf v1.26.0
