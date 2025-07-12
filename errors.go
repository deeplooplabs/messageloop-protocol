package protocol

import (
	"fmt"
	sharedv1 "github.com/deeplooplabs/messageloop-protocol/gen/proto/go/shared/v1"
)

type Error struct {
	Code     int                    `json:"code"`
	Reason   string                 `json:"reason"`
	Message  string                 `json:"message"`
	Metadata map[string]interface{} `json:"metadata"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %d", e.Reason, e.Code)
}

var _ error = new(Error)

func NewErrorFromProto(e *sharedv1.Error) *Error {
	md := make(map[string]interface{})
	for k, v := range e.Metadata {
		md[k] = v
	}
	return &Error{
		Code:     int(e.Code),
		Reason:   e.Reason,
		Message:  e.Message,
		Metadata: md,
	}
}
