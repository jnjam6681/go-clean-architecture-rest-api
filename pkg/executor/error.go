package executor

import (
	"fmt"
)

type Error struct {
	Err     error
	Execute string
	Stderr  string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %q => %s", e.Err, e.Execute, e.Stderr)
}
