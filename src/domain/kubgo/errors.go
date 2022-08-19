package kubgo

import "fmt"

// Create domain custom errors

type ExceedsCostKubgoError struct {
	Cost int32
}

func (e *ExceedsCostKubgoError) Error() string {
	return fmt.Sprintf("cost %d exceeds the budget.", e.Cost)
}
