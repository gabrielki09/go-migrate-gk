package model

import (
	"log"
)

func Run(option Options) error {
	dirs, err := resolveFileDir(option.Command)
	if err != nil {
		log.Println("erro aqui 1")
		return err
	}

	fileConfig := File{
		Name:             option.Name,
		FilePaths:        dirs,
		SeparateByFolder: option.SeparateByFolder,
	}

	return createFiles(fileConfig, option)

}
