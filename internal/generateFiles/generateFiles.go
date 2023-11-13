package generatefiles

import (
	"fmt"
	"os"
)

func GenerateFiles(dirs, files []string) error {

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório atual:", err)
		return err
	}

	for _, dir := range dirs {
		dirPath := currentDir + "/" + dir
		err := os.Mkdir(dirPath, os.ModePerm)
		if err != nil {
			fmt.Printf("Erro ao criar o diretório %s: %v\n", dir, err)
			return err
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

	return nil

}
