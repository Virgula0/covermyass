package analysis

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sundowndev/covermyass/v2/lib/check"
	"github.com/sundowndev/covermyass/v2/utils"
)

type Summary struct {
	TotalFiles   int
	TotalRWFiles int
}

type Result struct {
	Check    check.Check
	Path     string
	Size     int64
	Mode     os.FileMode
	ReadOnly bool
}

type Analysis struct {
	Date     time.Time
	summary  Summary
	patterns []string
	results  []Result
}

func NewAnalysis() *Analysis {
	return &Analysis{
		Date:     time.Now(),
		summary:  Summary{},
		patterns: []string{},
		results:  []Result{},
	}
}

func (a *Analysis) AddResult(result Result) {
	a.results = append(a.results, result)
	a.summary.TotalFiles += 1
	if !result.ReadOnly {
		a.summary.TotalRWFiles += 1
	}
}

func (a *Analysis) Results() []Result {
	return a.results
}

func (a *Analysis) Write(w io.Writer) {
	if len(a.results) > 0 {
		_, _ = fmt.Fprintf(w, "Found the following files\n")

		for _, res := range a.results {
			_, _ = fmt.Fprintf(w, "%s (%s, %s)\n", res.Path, utils.ByteCountSI(res.Size), res.Mode.String())
		}
		_, _ = fmt.Fprintf(w, "\n")
	}

	_, _ = fmt.Fprintf(
		w,
		"Summary\nFound %d files (%d read-write, %d read-only) in %s\n",
		a.summary.TotalFiles,
		a.summary.TotalRWFiles,
		a.summary.TotalFiles-a.summary.TotalRWFiles,
		time.Since(a.Date).Round(time.Millisecond).String(),
	)
}
