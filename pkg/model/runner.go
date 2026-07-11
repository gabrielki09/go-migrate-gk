package model

import (
	"fmt"
	"log"
)

func Run(option Options) error {
	dirs, err := resolveFileDir(option.Command)
	if err != nil {
		log.Println("erro aqui 1")
		return err
	}

	if option.Command["uuid_use"] && option.Command["id_use"] {
		return fmt.Errorf("somente um tipo de ID pode ser utilizado.")
	}

	fileConfig := File{
		Name:             option.Name,
		FilePaths:        dirs,
		SeparateByFolder: option.SeparateByFolder,
	}

	return createFiles(fileConfig, option)

}
