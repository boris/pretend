package pretend

import (
	"fmt"
	"github.com/fatih/color"
)

type Script struct {
	name  string
	steps []Step
}

type Step struct {
	desc  string
	cmd   string
	group string // optional group name
}

func NewScript(name string) *Script {
	return &Script{name: name, steps: []Step{}}
}

func (s *Script) Step(desc, cmd string) {
	s.steps = append(s.steps, Step{desc: desc, cmd: cmd})
}

func (s *Script) Group(name string, fn func()) {
	s.steps = append(s.steps, Step{desc: name, cmd: ""})
	fn()
}

func (s *Script) Run() {
	cmdColor := color.New(color.FgGreen).SprintFunc()
	groupColor := color.New(color.FgBlue).SprintFunc()
	nameColor := color.New(color.FgCyan).SprintFunc()

	fmt.Printf("%s\n", nameColor(s.name))
	for _, step := range s.steps {
		if step.cmd == "" {
			fmt.Printf(groupColor(step.desc))
		} else {
			fmt.Printf("%s: %s", step.desc, cmdColor(step.cmd))
		}
		fmt.Scanln()
	}
}
