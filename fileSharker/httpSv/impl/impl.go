package impl

import (
	"context"

	"lyh/ddkava/fileSharker/Park"
)

const pakgeName = "httpImpl"

var ctx context.Context

func init() {
	Park.RegisterStart(pakgeName, func(ctx context.Context) []error {
		ctx = ctx
		return []error{}
	})
}
