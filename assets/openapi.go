package assets

import (
	_ "embed"
)

var (
	//go:embed openapi.yaml
	openapiData []byte
)

func Data() []byte {
	return openapiData
}
