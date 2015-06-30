package multierror

import (
	"fmt"
	"strings"
)

type Error []error

func (me Error) Error() string {
	if me == nil {
		return ""
	}

	strs := make([]string, len(me))
	for i, err := range me {
		strs[i] = fmt.Sprintf("[%d] %v", i, err)
	}
	return strings.Join(strs, " ")
}

func (me Error) AsError() error {
	if len([]error(me)) <= 0 {
		return nil
	}

	return me
}
