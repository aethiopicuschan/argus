package argus_test

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/aethiopicuschan/argus"
	"github.com/stretchr/testify/require"
)

func ExampleBuilder_Print() {
	l := argus.NewLogger(os.Stdout)
	l.Info().Remove("time").Add("key", "value").Print()
	// Output: {"level":"INFO","key":"value"}
}

func TestWithMinLevelCases(t *testing.T) {
	testCases := []struct {
		name           string
		minLevel       argus.Level
		logFunc        func(logger *argus.Logger) *argus.Builder
		expectedOutput bool
		expectedLevel  argus.Level
	}{
		{
			name:           "MinLevel Warn, logging Info",
			minLevel:       argus.Warn,
			logFunc:        func(logger *argus.Logger) *argus.Builder { return logger.Info().Add("message", "should be ignored") },
			expectedOutput: false,
		},
		{
			name:           "MinLevel Warn, logging Warn",
			minLevel:       argus.Warn,
			logFunc:        func(logger *argus.Logger) *argus.Builder { return logger.Warn().Add("message", "logged") },
			expectedOutput: true,
			expectedLevel:  argus.Warn,
		},
		{
			name:           "MinLevel Error, logging Error",
			minLevel:       argus.Error,
			logFunc:        func(logger *argus.Logger) *argus.Builder { return logger.Error().Add("message", "logged") },
			expectedOutput: true,
			expectedLevel:  argus.Error,
		},
		{
			name:           "MinLevel Info, logging Debug",
			minLevel:       argus.Info,
			logFunc:        func(logger *argus.Logger) *argus.Builder { return logger.Debug().Add("message", "should be ignored") },
			expectedOutput: false,
		},
		{
			name:           "MinLevel Debug, logging Debug",
			minLevel:       argus.Debug,
			logFunc:        func(logger *argus.Logger) *argus.Builder { return logger.Debug().Add("message", "logged") },
			expectedOutput: true,
			expectedLevel:  argus.Debug,
		},
		{
			name:           "MinLevel Error, logging Warn",
			minLevel:       argus.Error,
			logFunc:        func(logger *argus.Logger) *argus.Builder { return logger.Warn().Add("message", "should be ignored") },
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			buf := new(bytes.Buffer)
			logger := argus.NewLogger(buf, argus.WithMinLevel(tc.minLevel))
			builder := tc.logFunc(logger)
			if err := builder.Print(); err != nil {
				t.Fatalf("Print failed: %v", err)
			}

			if tc.expectedOutput {
				if buf.Len() == 0 {
					t.Errorf("expected output, but buffer is empty")
				} else {
					var result map[string]any
					if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
						t.Fatalf("json.Unmarshal failed: %v", err)
					}
					levelVal, ok := result["level"].(string)
					if !ok {
						t.Errorf("expected level field to be string")
					}
					require.Equal(t, tc.expectedLevel, argus.Level(levelVal))
				}
			} else {
				if buf.Len() != 0 {
					t.Errorf("expected no output, but buffer has data: %s", buf.String())
				}
			}
		})
	}
}
