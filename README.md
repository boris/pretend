# pretend

**pretend** is a Go library for documenting operational workflows as interactive, human-in-the-loop scripts.  
It lets you describe your manual procedures as step-by-step “pretend” scripts, pausing at each step for confirmation, without making any real changes to your systems.

_"Good documentation is almost runnable"_

## Features

- Write interactive “do-nothing” scripts in Go
- Clearly separate groups/phases of steps
- Colorful terminal output for readability
- Pause for operator confirmation after each step
- Perfect for onboarding, runbooks, or documenting processes before automation

## Installation

```bash
go get github.com/boris/pretend
```

Also requires `fatih/color` (automatically installed by go get).

## Usage
```go
package main

import "github.com/yourusername/pretend"

func main() {
    script := pretend.NewScript("Deploy Web App")

    script.Step("Clone repository", "git clone https://example.com/repo.git")
    script.Step("Change directory", "cd repo")

    script.Group("Install dependencies", func() {
        script.Step("Install Node.js dependencies", "npm install")
        script.Step("Install Python dependencies", "pip install -r requirements.txt")
    })

    script.Run()
}
```

See [main.go](./main.go) for a complete example.

## Why?

- Share operational runbooks in a way that’s unambiguous and interactive
- Document “how-to”s for processes you might automate later
- Help new teammates safely learn infrastructure workflows
- Provide a bridge between human and automated operations
