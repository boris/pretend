package pretend

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"github.com/fatih/color"
)

type Script struct {
	Name  string `json:"name"`
	Steps []Step `json:"steps"`
}

type Step struct {
	Desc  string `json:"desc"`
	Cmd   string `json:"cmd,omitempty"`
	Steps []Step `json:"steps,omitempty"`
}

func (s *Script) Run() {
	nameColor := color.New(color.FgCyan).SprintFunc()
	fmt.Printf("%s\n", nameColor(s.Name))
	runSteps(s.Steps, 0)
}

func runSteps(steps []Step, indent int) {
	cmdColor := color.New(color.FgGreen).SprintFunc()
	groupColor := color.New(color.FgBlue).SprintFunc()
	indentStr := func(level int) string {
		return strings.Repeat("  ", level)
	}

	for _, step := range steps {
		prefix := indentStr(indent)
		if len(step.Steps) > 0 {
			fmt.Printf("%s%s\n", prefix, groupColor(step.Desc))
			runSteps(step.Steps, indent+1)
		} else {
			fmt.Printf("%s%s: %s (press Enter to continue)", prefix, step.Desc, cmdColor(step.Cmd))
			fmt.Scanln()
		}
	}
}

func ScriptFromFile(filePath string) (*Script, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open steps file: %w", err)
	}
	defer f.Close()

	var s Script
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&s); err != nil {
		return nil, fmt.Errorf("failed to decode steps file: %w", err)
	}

	return &s, nil
}
