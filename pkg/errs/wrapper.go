package errs

import (
	"errors"
	"fmt"
)

func Joinf(err error, msg string, args ...any) error {
	if err == nil {
		return nil
	}
	return errors.Join(err, errors.New(fmt.Sprintf(msg, args...)))
}
