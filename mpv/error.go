package mpv

//#include <mpv/client.h>
import "C"

import (
	"fmt"
)

type Error int

//const char *mpv_error_string(int error);
func NewError(errcode C.int) Error {
	return Error(errcode)
}

func (m Error) Error() string {
	return fmt.Sprintln("MPV_ERROR", int(m), C.GoString(C.mpv_error_string(C.int(m))))
}