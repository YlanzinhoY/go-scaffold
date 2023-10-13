package main

import (
	"fmt"
	"os"
)

func main() {
	projectDirs := []string{
		"cmd",
		"pkg",
		"internal",
		"api",
		"web",
		"config",
		"data",
		"scripts",
		"docs",
		"test",
	}

	files := []string{
		"dockerfile",
		"docker-compose.yml",
		"makefile",
	}

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório atual:", err)
		return
	}

	for _, dir := range projectDirs {
		dirPath := currentDir + "/" + dir
		err := os.Mkdir(dirPath, os.ModePerm)
		if err != nil {
			fmt.Printf("Erro ao criar o diretório %s: %v\n", dir, err)
		} else {
			fmt.Printf("Diretório %s criado com sucesso!\n", dir)
		}
	}

	for _, createFile := range files {
		_, err := os.Create(createFile)

		if err != nil {
			fmt.Printf("Erro ao criar o arquivo %s", createFile)
		} else {
			fmt.Printf("Arquivo %s criado com scuesso!\n", createFile)
		}

	}
}
