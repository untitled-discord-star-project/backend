package trace

import "github.com/google/uuid"

type Trace struct {
	TraceID uuid.UUID
	RequestID uuid.UUID
}
