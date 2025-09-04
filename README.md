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

Or, you can get the binary directly from the [releases page](https://github.com/boris/pretend/releases).

## Usage
Describe your steps in JSON, then run `pretend -help` for options.
```json
{
  "name": "Deploy Web App",
  "steps": [
    { "desc": "Clone repository", "cmd": "git clone https://example.com/repo.git" },
    {
      "desc": "Install dependencies",
      "steps": [
        { "desc": "Python", "cmd": "pip install -r requirements.txt" }
      ]
    }
  ]
}

```

See [examples](./examples) directory for a complete example.

## Why?

- Share operational runbooks in a way that’s unambiguous and interactive
- Document “how-to”s for processes you might automate later
- Help new teammates safely learn infrastructure workflows
- Provide a bridge between human and automated operations
