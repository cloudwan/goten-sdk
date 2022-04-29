package client

import (
	"google.golang.org/protobuf/proto"

	"github.com/cloudwan/goten-sdk/runtime/resource"
)

// standard verbs
var (
	StdVerbGet      = "get"
	StdVerbBatchGet = "batchGet"
	StdVerbList     = "list"
	StdVerbWatch    = "watch"
	StdVerbSearch   = "search"
	StdVerbCreate   = "create"
	StdVerbUpdate   = "update"
	StdVerbDelete   = "delete"
)

// MethodDescriptor allows writing code operating on Method without knowing exact input/output
// type or name. It also allows accessing common properties of a method.
type MethodDescriptor interface {
	// NewEmptyClientMsg returns empty, freshly allocated client message for this method
	NewEmptyClientMsg() proto.Message

	// NewEmptyServerMsg returns empty, freshly allocated server message for this method
	NewEmptyServerMsg() proto.Message

	// IsUnary returns true if method is unary (IsClientStream = false and IsServerStream = false)
	IsUnary() bool

	// IsClientStream returns true if method has client streaming
	IsClientStream() bool

	// IsServerStream returns true if method has server streaming
	IsServerStream() bool

	// IsCollectionSubject returns true if method operates on resource collection
	IsCollectionSubject() bool

	// IsPluralSubject return true if method operates on multiple items in resource collection
	IsPluralSubject() bool

	// HasSubjectResource returns true if method has subject resource defined
	HasSubjectResource() bool

	// RequestHasResourceBody returns true if method contains subject resource body
	RequestHasResourceBody() bool

	// GetVerb returns verb of method, for example "list", "search"...
	GetVerb() string

	// GetMethodName returns RPC method name (as defied in proto file), for example ListBooks
	GetMethodName() string

	// GetFullMethodName returns fully qualified name that is same as used in grpc.UnaryStreamInfo
	// (or stream equivalent), for example /goten.example.library.v1beta.BookService/ListBooks
	GetFullMethodName() string

	// GetProtoPkgName returns proto package where method and corresponding RPC service is defined,
	// for example "goten.example.library.v1beta"
	GetProtoPkgName() string

	// GetApiName returns Goten API name (GRPC service as defined in proto file), for example BookService
	GetApiName() string

	// GetServiceDomain returns domain of Goten service, as defined by field "name" in api-skeleton yaml file,
	// for example library.edgelq.com
	GetServiceDomain() string

	// GetServiceVersion returns version of Goten service, as defined by field "proto.package.currentVersion"
	// field in api-skeleton yaml file, for example "v1"
	GetServiceVersion() string

	// GetApiDescriptor returns descriptor of service owning this method
	GetApiDescriptor() ApiDescriptor

	// GetSubjectResourceDescriptor returns descriptor of resource which is subject to this method (may be nil)
	GetSubjectResourceDescriptor() resource.Descriptor

	// GetClientMsgReflectHandle returns handle accessing common properties from client message
	GetClientMsgReflectHandle() MethodMsgHandle

	// GetServerMsgReflectHandle returns handle accessing common properties from server message
	GetServerMsgReflectHandle() MethodMsgHandle
}

// MethodMsgHandle allows accessing common properties of a request/response object without knowing
// exact type. Goten provides basic code-gen for standard methods (and custom ones, as long as request
// has standard "name" or "parent" field in a root).
type MethodMsgHandle interface {
	// ExtractSubjectResourceName returns subject resource name from request/response objects. For example,
	// handle for Get<Resource>/Get<Resource>Request will return value "name". For
	// Create<Resource>/Create<Resource>Request, it will return value of "<resource>.name" field. For
	// Create<Resource>/<Resource> (response!) it will return value of "name" field.
	ExtractSubjectResourceName(msg proto.Message) resource.Name

	// ExtractSubjectResourceNames returns list of subject resource names from request/response objects. It will
	// be used for batchGet requests (field "names"!) or responses for List/Watch/Search queries.
	ExtractSubjectResourceNames(msg proto.Message) resource.NameList

	// ExtractSubjectCollectionName returns subject resource collection name from request/response objects (However,
	// at this moment we dont code-gen anything for responses - although if user provides "Override", then it can).
	ExtractSubjectCollectionName(msg proto.Message) resource.Name

	// Methods below can be implemented by developer for generated MethodMsgHandle in separate, non-pb go file.
	// If they are defined, methods above will use below versions:
	// OverrideExtractSubjectResourceName(msg *<concrete_msg>) *<concrete_name>
	// OverrideExtractSubjectResourceNames(msg *<concrete_msg>) []*<concrete_name>
	// OverrideExtractSubjectCollectionName(msg *<concrete_msg>) *<concrete_name>
}
