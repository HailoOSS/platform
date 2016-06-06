package schema

import (
	"github.com/hailocab/protobuf/proto"

	"github.com/HailoOSS/platform/errors"
	"github.com/HailoOSS/platform/server"

	"github.com/hailocab/go-hailo-lib/schema"
	schemaProto "github.com/HailoOSS/platform/proto/schema"
	service "github.com/HailoOSS/platform/server"
)

func Endpoint(name string, configStruct interface{}) *service.Endpoint {
	handler := func(req *server.Request) (proto.Message, errors.Error) {
		return &schemaProto.Response{
			Schema: proto.String(schema.Of(configStruct).String()),
		}, nil
	}

	return &server.Endpoint{
		Name:       name,
		Mean:       200,
		Upper95:    400,
		Handler:    handler,
		Authoriser: service.OpenToTheWorldAuthoriser(),
	}
}
