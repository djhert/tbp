package tbp

import (
	"fmt"
	"github.com/hlfstr/ego2d/ego"
)

// Constants that defines self-awareness
// NAME    is the name
// MAJVERSION is the major version number
// MINVERSION is the minor version number
// RELVERSION is the release version number
const (
	NAME       = "Trial By Pong"
	MAJVERSION = 0
	MINVERSION = 1
	RELVERSION = 0
)

// Version returns a string that contains
// the name and version all nice and neat
func Version() string {
	return fmt.Sprintf("%s v%d.%d.%d - Built with: %s", NAME,
		MAJVERSION, MINVERSION, RELVERSION, ego.Version())
}
