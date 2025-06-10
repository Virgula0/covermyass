//go:build !windows

package analysis

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sundowndev/covermyass/v2/mocks"
)

func TestAnalysis_Write(t *testing.T) {
	fakeCheck := &mocks.Check{}

	testcases := []struct {
		name     string
		filename string
		results  []Result
	}{
		{
			name:     "test valid analysis",
			filename: "testdata/valid_analysis.txt",
			results: []Result{
				{
					Path:     "/var/log/dummy.log",
					Size:     int64(32),
					Mode:     os.ModePerm,
					ReadOnly: false,
					Check:    fakeCheck,
				},
				{
					Path:     "/var/log/dummy2.log",
					Size:     int64(8),
					Mode:     600,
					ReadOnly: true,
					Check:    fakeCheck,
				},
			},
		},
		{
			name:     "test empty analysis",
			filename: "testdata/empty_analysis.txt",
			results:  []Result{},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			testdata, err := os.ReadFile(tt.filename)
			if err != nil {
				t.Fatal(err)
			}

			a := NewAnalysis()
			for _, r := range tt.results {
				a.AddResult(r)
			}
			assert.Len(t, a.Results(), len(tt.results))

			buffer := &bytes.Buffer{}
			a.Write(buffer)

			assert.Equal(t, string(testdata), buffer.String())
		})
	}

	fakeCheck.AssertExpectations(t)
}

func Test_byteCountSI(t *testing.T) {
	testcases := map[int64]string{
		0:                   "0 B",
		1:                   "1 B",
		10000:               "10.0 kB",
		22736526:            "22.7 MB",
		2123652683:          "2.1 GB",
		2123652683000:       "2.1 TB",
		2123652683000000:    "2.1 PB",
		2123652683000000000: "2.1 EB",
	}

	for v, expected := range testcases {
		t.Run(fmt.Sprintf("test with input %d", v), func(t *testing.T) {
			assert.Equal(t, expected, byteCountSI(v))
		})
	}
}
