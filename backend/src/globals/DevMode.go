package globals

import "os"

var (
	DevMode bool = os.Getenv("GO_ENV") == "development"
)
