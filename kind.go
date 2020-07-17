package errors

import "net/http"

type Kind string

const (
	// KindCanceled indicates the operation was cancelled (typically by the caller).
	KindCanceled Kind = "CANCELED"

	// KindUnknown error. For example when handling errors raised by APIs that do not
	// return enough error information.
	KindUnknown Kind = "UNKNOWN"

	// KindInvalidArgument indicates client specified an invalid argument. It
	// indicates arguments that are problematic regardless of the state of the
	// system (i.e. a malformed file name, required argument, number out of range,
	// etc.).
	KindInvalidArgument Kind = "INVALID_ARGUMENT"

	// KindDeadlineExceeded means operation expired before completion. For operations
	// that change the state of the system, this error may be returned even if the
	// operation has completed successfully (timeout).
	KindDeadlineExceeded Kind = "DEADLINE_EXCEEDED"

	// KindNotFound means some requested entity was not found.
	KindNotFound Kind = "NOT_FOUND"

	// AlreadyExists means an attempt to create an entity failed because one
	// already exists.
	KindAlreadyExists Kind = "ALREADY_EXISTS"

	// KindPermissionDenied indicates the caller does not have permission to execute
	// the specified operation. It must not be used if the caller cannot be
	// identified (Unauthenticated).
	KindPermissionDenied Kind = "PERMISSION_DENIED"

	// KindUnauthenticated indicates the request does not have valid authentication
	// credentials for the operation.
	KindUnauthenticated Kind = "UNAUTHENTICATED"

	// KindResourceExhausted indicates some resource has been exhausted, perhaps a
	// per-user quota, or perhaps the entire file system is out of space.
	KindResourceExhausted Kind = "RESOURCE_EXHAUSTED"

	// KindFailedPrecondition indicates operation was rejected because the system is
	// not in a state required for the operation's execution. For example, doing
	// an rmdir operation on a directory that is non-empty, or on a non-directory
	// object, or when having conflicting read-modify-write on the same resource.
	KindFailedPrecondition Kind = "FAILED_PRECONDITION"

	// KindAborted indicates the operation was aborted, typically due to a concurrency
	// issue like sequencer check failures, transaction aborts, etc.
	KindAborted Kind = "ABORTED"

	// KindOutOfRange means operation was attempted past the valid range. For example,
	// seeking or reading past end of a paginated collection.
	//
	// Unlike InvalidArgument, this error indicates a problem that may be fixed if
	// the system state changes (i.e. adding more items to the collection).
	//
	// There is a fair bit of overlap between FailedPrecondition and OutOfRange.
	// We recommend using OutOfRange (the more specific error) when it applies so
	// that callers who are iterating through a space can easily look for an
	// OutOfRange error to detect when they are done.
	KindOutOfRange Kind = "OUT_OF_RANGE"

	// KindUnimplemented indicates operation is not implemented or not
	// supported/enabled in this service.
	KindUnimplemented Kind = "UNIMPLEMENTED"

	// KindInternal errors. When some invariants expected by the underlying system
	// have been broken. In other words, something bad happened in the library or
	// backend service. Do not confuse with HTTP Internal Server Error; an
	// Internal error could also happen on the client code, i.e. when parsing a
	// server response.
	KindInternal Kind = "INTERNAL"

	// KindUnavailable indicates the service is currently unavailable. This is a most
	// likely a transient condition and may be corrected by retrying with a
	// backoff.
	KindUnavailable Kind = "UNAVAILABLE"

	// DataLoss indicates unrecoverable data loss or corruption.
	KindDataLoss Kind = "DATA_LOSS"

	// KindNone is the zero-value, is considered an empty error and should not be
	// used.
	KindNone Kind = ""
)

func (k Kind) ToStatusCode() int {
	switch k {
	case KindCanceled, KindDeadlineExceeded:
		return http.StatusRequestTimeout
	case KindUnknown:
		return http.StatusInternalServerError
	case KindInvalidArgument:
		return http.StatusBadRequest
	case KindNotFound:
		return http.StatusNotFound
	case KindAlreadyExists, KindAborted:
		return http.StatusConflict
	case KindPermissionDenied:
		return http.StatusForbidden
	case KindUnauthenticated:
		return http.StatusUnauthorized
	case KindResourceExhausted:
		return http.StatusForbidden
	case KindFailedPrecondition:
		return http.StatusPreconditionFailed
	case KindOutOfRange:
		return http.StatusBadRequest
	case KindUnimplemented:
		return http.StatusNotImplemented
	case KindInternal, KindDataLoss, KindNone:
		return http.StatusInternalServerError
	case KindUnavailable:
		return http.StatusServiceUnavailable
	}

	return http.StatusInternalServerError
}
