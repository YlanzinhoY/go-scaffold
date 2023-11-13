package internal

import (
	generatefiles "github.com/ylanzinhoy/go-scaffold/internal/generateFiles"
)

func MaximumFiles() {
	projectDirs := []string{
		"cmd",
		"pkg",
		"internal",
		"api",
		"build",
		"deployments",
		"examples",
		"web",
		"assets",
		"githooks",
		"third_party",
		"vendor",
		"website",
		"tools",
		"configs",
		"data",
		"scripts",
		"docs",
		"test",
	}

	files := []string{
		"dockerfile",
		"docker-compose.yml",
		"Makefile",
	}

	err := generatefiles.GenerateFiles(projectDirs, files)

	if err != nil {
		panic(err)
	}
	
}
