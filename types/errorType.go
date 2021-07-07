package types

import "fmt"

type ConfigErr struct {
	s string
}

func (e *ConfigErr) Error() string {
	return fmt.Sprintf("can not find config: %s", e.s)
}
