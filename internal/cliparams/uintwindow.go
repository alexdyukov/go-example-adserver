package cliparams

import (
	"fmt"
	"math"
	"strconv"
)

type UintWindow int64

var errInvalidValue = fmt.Errorf("should be integer in range [0;%v]", math.MaxInt64)

func (window *UintWindow) UnmarshalText(text []byte) error {
	return window.Set(string(text))
}

func (window *UintWindow) String() string {
	return fmt.Sprint(*window)
}

func (window *UintWindow) Int64() int64 {
	return int64(*window)
}

func (window *UintWindow) Set(value string) error {
	intval, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return errInvalidValue
	}

	if intval < 0 {
		return errInvalidValue
	}

	*window = UintWindow(intval)

	return nil
}
