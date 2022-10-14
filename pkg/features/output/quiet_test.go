package output

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/output"
	"testing"
)

func TestQuietIsEnabled(t *testing.T) {
	conf := cli.Configuration{Quiet: true}
	quiet := Quiet{}
	if !quiet.IsEnabled(&conf) {
		t.Fatal("Execution for quiet flag not working")
	}
}

func TestQuietExecute(t *testing.T) {
	_ = Quiet{}.Execute(nil)

	if output.Level() != output.LevelSilence {
		t.Fatalf("Expected log level to be LevelSilence but got %d", output.Level())
	}
}
