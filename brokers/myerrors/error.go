package myerrors

import "errors"

var ErrDetailsExist = errors.New("DETAILS_EXIST")
var ErrNIL = errors.New("NIL")
var ErrNoRecordsDeleted = errors.New("no records were deleted")
