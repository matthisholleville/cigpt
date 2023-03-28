package main

import "gitlab.com/cigpt-ai/cigpt/cmd"

var version = "dev"

func main() {
	cmd.Execute(version)
}
