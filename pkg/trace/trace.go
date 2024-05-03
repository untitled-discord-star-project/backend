package trace

import "github.com/oklog/ulid/v2"

type Trace struct {
	TraceID   ulid.ULID
	RequestID ulid.ULID
}
