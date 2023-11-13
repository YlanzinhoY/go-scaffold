package internal

import (
	"fmt"
	"os"
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
