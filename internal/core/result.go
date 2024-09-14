package core

import "github.com/weissmedia/searchengine/internal/profiler"

type ExecutionResult struct {
	ResultSet []string
	Timings   []profiler.OperationTiming
}
