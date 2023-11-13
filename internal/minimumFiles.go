package internal

import generatefiles "github.com/ylanzinhoy/go-scaffold/internal/generateFiles"

func MinimumFiles() {
	projectDirs := []string{
		"cmd",
		"pkg",
		"internal",
		"api",
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
