package filter

import (
	"path/filepath"

	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/features/conditions"
	"github.com/timo-reymann/deterministic-zip/pkg/output"
)

// Exclude given patterns like a block list
type Exclude struct {
}

// DebugName returns the debuggable name
func (e Exclude) DebugName() string {
	return "Exclude"
}

// IsEnabled checks if exclude flags are present
func (e Exclude) IsEnabled(c *cli.Configuration) bool {
	return conditions.HasElements(&c.Exclude)
}

// Execute exclude against source files and mutate back the cleaned ones
func (e Exclude) Execute(c *cli.Configuration) error {
	var fileExcluded bool
	files := make([]string, 0)
	excludes := c.Exclude

	for _, f := range c.SourceFiles {
		fileExcluded = false
		for _, pattern := range excludes {
			isMatch, _ := filepath.Match(pattern, f)
			println(pattern, f, isMatch)

			if isMatch {
				fileExcluded = true
				break
			}
		}

		if !fileExcluded {
			files = append(files, f)
		} else {
			output.Debugf("%s doesnt match exclude patterns, skipping", f)
		}
	}

	c.SourceFiles = files
	return nil
}
