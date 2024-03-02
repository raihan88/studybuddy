/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/raihan88/studybuddy/cmd"
	"github.com/raihan88/studybuddy/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
