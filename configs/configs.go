package configs

import _ "embed"

//go:embed configs.yaml
var DefaultConfigBytes []byte
