// Copyright 2017 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package status provides utility functions for google_rpc status objects.
package status

import (
	rpc "github.com/gogo/googleapis/google/rpc"
)

// OK represents a status with a code of rpc.OK
var OK = rpc.Status{Code: int32(rpc.OK)}

// New returns an initialized status with the given error code.
func New(c rpc.Code) rpc.Status {
	return rpc.Status{Code: int32(c)}
}

// WithMessage returns an initialized status with the given error code and message
func WithMessage(c rpc.Code, message string) rpc.Status {
	return rpc.Status{Code: int32(c), Message: message}
}

// WithError returns an initialized status with the rpc.INTERNAL error code and the error's message.
func WithError(err error) rpc.Status {
	return rpc.Status{Code: int32(rpc.INTERNAL), Message: err.Error()}
}

// WithInternal returns an initialized status with the rpc.INTERNAL error code and the error's message.
func WithInternal(message string) rpc.Status {
	return rpc.Status{Code: int32(rpc.INTERNAL), Message: message}
}

// WithCancelled returns an initialized status with the rpc.CANCELLED error code and the error's message.
func WithCancelled(message string) rpc.Status {
	return rpc.Status{Code: int32(rpc.CANCELLED), Message: message}
}

// WithInvalidArgument returns an initialized status with the rpc.INVALID_ARGUMENT code and the given message.
func WithInvalidArgument(message string) rpc.Status {
	return rpc.Status{Code: int32(rpc.INVALID_ARGUMENT), Message: message}
}

// WithPermissionDenied returns an initialized status with the rpc.PERMISSION_DENIED code and the given message.
func WithPermissionDenied(message string) rpc.Status {
	return rpc.Status{Code: int32(rpc.PERMISSION_DENIED), Message: message}
}

// WithResourceExhausted returns an initialized status with the rpc.PERMISSION_DENIED code and the given message.
func WithResourceExhausted(message string) rpc.Status {
	return rpc.Status{Code: int32(rpc.RESOURCE_EXHAUSTED), Message: message}
}

// WithDeadlineExceeded returns an initialized status with the rpc.DEADLINE_EXCEEDED code and the given message.
func WithDeadlineExceeded(message string) rpc.Status {
	return rpc.Status{Code: int32(rpc.DEADLINE_EXCEEDED), Message: message}
}

// WithUnknown returns an initialized status with the rpc.UNKNOWN code and the given message.
func WithUnknown(message string) rpc.Status {
	return rpc.Status{Code: int32(rpc.UNKNOWN), Message: message}
}

// WithNotFound returns an initialized status with the rpc.NOT_FOUND code and the given message.
func WithNotFound(message string) rpc.Status {
	return rpc.Status{Code: int32(rpc.NOT_FOUND), Message: message}
}

// WithAlreadyExists returns an initialized status with the rpc.ALREADY_EXISTS code and the given message.
func WithAlreadyExists(message string) rpc.Status {
	return rpc.Status{Code: int32(rpc.ALREADY_EXISTS), Message: message}
}

// WithFailedPrecondition returns an initialized status with the rpc.FAILED_PRECONDITION code and the given message.
func WithFailedPrecondition(message string) rpc.Status {
	return rpc.Status{Code: int32(rpc.FAILED_PRECONDITION), Message: message}
}

// WithAborted returns an initialized status with the rpc.ABORTED code and the given message.
func WithAborted(message string) rpc.Status {
	return rpc.Status{Code: int32(rpc.ABORTED), Message: message}
}

// WithOutOfRange returns an initialized status with the rpc.OUT_OF_RANGE code and the given message.
func WithOutOfRange(message string) rpc.Status {
	return rpc.Status{Code: int32(rpc.OUT_OF_RANGE), Message: message}
}

// WithUnimplemented returns an initialized status with the rpc.UNIMPLEMENTED code and the given message.
func WithUnimplemented(message string) rpc.Status {
	return rpc.Status{Code: int32(rpc.UNIMPLEMENTED), Message: message}
}

// WithUnavailable returns an initialized status with the rpc.UNAVAILABLE code and the given message.
func WithUnavailable(message string) rpc.Status {
	return rpc.Status{Code: int32(rpc.UNAVAILABLE), Message: message}
}

// WithDataLoss returns an initialized status with the rpc.DATA_LOSS code and the given message.
func WithDataLoss(message string) rpc.Status {
	return rpc.Status{Code: int32(rpc.DATA_LOSS), Message: message}
}

// WithUnauthenticated returns an initialized status with the rpc.UNAUTHENTICATED code and the given message.
func WithUnauthenticated(message string) rpc.Status {
	return rpc.Status{Code: int32(rpc.UNAUTHENTICATED), Message: message}
}

// IsOK returns true is the given status has the code rpc.OK
func IsOK(status rpc.Status) bool {
	return status.Code == int32(rpc.OK)
}

// String produces a string representation of rpc.Status.
func String(status rpc.Status) string {
	result, ok := rpc.Code_name[status.Code]
	if !ok {
		result = "Code" + string(status.Code)
	}

	if status.Message != "" {
		result = result + " (" + status.Message + ")"
	}
	return result
}
