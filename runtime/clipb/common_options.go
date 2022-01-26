package clipb

import (
	"context"
	"strings"
	"time"

	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc/metadata"

	"github.com/spf13/cobra"
)

const (
	FlagToken                 = "auth-token"
	FlagJwtKey                = "jwt-key"
	FlagEndpoint              = "server-addr"
	FlagTlsInsecureSkipVerify = "tls-insecure-skip-verify"
	FlagTlsCaCertFile         = "tls-ca-cert-file"
	FlagResponseFormat        = "response-format"
	FlagTimeout               = "timeout"
	FlagRawResponse           = "raw-response"
	FlagInputFile             = "input-file"
	FlagHttpHeaders           = "header"
	FlagWithResponseHeaders   = "with-response-headers"
)

type CommonOptions struct {
	ServiceName string

	*DialOptions
	*InputOptions
	*OutputOptions
}

type DialOptions struct {
	ServiceName string

	Token              string        `envconfig:"AUTH_TOKEN"`
	JwtKey             string        `envconfig:"JWT_KEY"`
	Endpoint           string        `envconfig:"SERVER_ADDR"`
	InsecureSkipVerify bool          `envconfig:"TLS_INSECURE_SKIP_VERIFY"`
	CACertFile         string        `envconfig:"TLS_CA_CERT_FILE"`
	Timeout            time.Duration `envconfig:"TIMEOUT" default:"10s"`
}

type InputOptions struct {
	ServiceName string

	InputFile   string `envconfig:"INPUT_FILE"`
	HttpHeaders []string
}

type OutputOptions struct {
	ServiceName string

	RawResponse         bool
	WithResponseHeaders bool
	ResponseFormat      string `envconfig:"RESPONSE_FORMAT" default:"table"`
}

func NewOptionsAttachedToCmd(cmd *cobra.Command, serviceName string) *CommonOptions {
	opts := &CommonOptions{
		ServiceName:   serviceName,
		DialOptions:   NewDialOptionsAttachedToCmd(cmd, serviceName),
		InputOptions:  NewInputOptionsAttachedToCmd(cmd, serviceName),
		OutputOptions: NewOutputOptionsAttachedToCmd(cmd, serviceName),
	}

	return opts
}

func NewDialOptionsAttachedToCmd(cmd *cobra.Command, serviceName string) *DialOptions {
	opts := &DialOptions{}
	if err := envconfig.Process("", opts); err != nil {
		panic(err)
	}

	opts.ServiceName = serviceName

	flags := cmd.PersistentFlags()
	flags.StringVar(&opts.Token, FlagToken, opts.Token, "Authorization token")
	flags.StringVar(&opts.JwtKey, FlagJwtKey, opts.JwtKey, "Jwt key")
	flags.StringVar(&opts.Endpoint, FlagEndpoint, opts.Endpoint, "Server address in form of host:port")
	flags.BoolVar(&opts.InsecureSkipVerify, FlagTlsInsecureSkipVerify, opts.InsecureSkipVerify, "INSECURE: Skip tls checks")
	flags.StringVar(&opts.CACertFile, FlagTlsCaCertFile, opts.CACertFile, "ca certificate file")
	flags.DurationVar(&opts.Timeout, FlagTimeout, opts.Timeout, "Timeout (default is 10s)")

	return opts
}

func NewInputOptionsAttachedToCmd(cmd *cobra.Command, serviceName string) *InputOptions {
	opts := &InputOptions{}
	if err := envconfig.Process("", opts); err != nil {
		panic(err)
	}

	opts.ServiceName = serviceName

	flags := cmd.PersistentFlags()
	flags.StringVarP(&opts.InputFile, FlagInputFile, "f", opts.InputFile, "Specifies input file. To read from stdin, specify '-'. Request property flags are ignored when this option is specified")
	flags.StringArrayVarP(&opts.HttpHeaders, FlagHttpHeaders, "H", opts.HttpHeaders, "List of http headers for request, for example -H \"cache-control: no-cache\" ")
	return opts
}

func NewOutputOptionsAttachedToCmd(cmd *cobra.Command, serviceName string) *OutputOptions {
	opts := &OutputOptions{}
	if err := envconfig.Process("", opts); err != nil {
		panic(err)
	}

	opts.ServiceName = serviceName

	flags := cmd.PersistentFlags()
	flags.StringVarP(&opts.ResponseFormat, FlagResponseFormat, "o", opts.ResponseFormat, "Response format (table or json, table is default)")
	flags.BoolVar(&opts.RawResponse, FlagRawResponse, opts.RawResponse, "Displays raw response instead of simplified one (not working right now)")
	flags.BoolVar(&opts.WithResponseHeaders, FlagWithResponseHeaders, opts.WithResponseHeaders, "Whether should print response headers before actual response")

	return opts
}

func (io *InputOptions) WithRequestHeaders(ctx context.Context) context.Context {
	requestHeaders := metadata.Pairs()
	for _, header := range io.HttpHeaders {
		header = strings.TrimSpace(header)
		if strings.HasSuffix(header, ";") {
			key := strings.TrimSuffix(header, ";")
			requestHeaders.Append(key)
		} else {
			items := strings.SplitN(header, ":", 2)
			if len(items) == 2 {
				key := strings.TrimSpace(items[0])
				values := strings.Split(items[1], ",")
				for i, val := range values {
					values[i] = strings.TrimSpace(val)
				}
				requestHeaders.Append(key, values...)
			}
		}
	}
	return metadata.NewOutgoingContext(ctx, requestHeaders)
}
