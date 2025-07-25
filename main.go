package main

import (
	"flag"
	"os"
	"github.com/boris/pretend/cmd/pretend"
)

func main(){
	scriptName := flag.String("script", "", "Name of the script to run")
	env := flag.String("env", "dev", "Environment name (e.g., dev, prod)")
	flag.Parse()

	switch *scriptName {
	case "deploy-web-app":
		deployWebApp(*env)
	default:
		flag.Usage()
		os.Exit(1)
	}
}

func deployWebApp(env string) {
	script := pretend.NewScript("Deploy Web App")

	script.Step("Check out the repository", "git clone https://example.com/repo.git")
	script.Step("Change directory to repo", "cd repo")

	script.Group("Install dependencies", func() {
	script.Step("Install Node.js dependencies", "npm install")
	script.Step("Install Python dependencies", "pip install -r requirements.txt")
	})

	if env == "prod" {
		script.Step("Keep an eye on the Mempool size", "Check the status in mempool.space")
	}

	script.Run()
}
