package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf("New version: Hello, Go. You said: %s", string(req))
}
