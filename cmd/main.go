package main

import (
	"flag"
	"log"

	scaffold "github.com/gabrielki09/go-scaffold-gk/pkg/scaffold"
)

func main() {
	var (
		path           = flag.String("path", "", "Comando para informar o caminho principal das operações")
		modelFlag      = flag.String("m", "", "Comando para criação de arquivo padrão da model")
		uuidUse        = flag.Bool("uuid", false, "Comando para criação do model com uuid")
		idUse          = flag.Bool("id", false, "Comando para criação do model com id (int)")
		requests       = flag.Bool("R", false, "Comando para criação da request")
		resource       = flag.Bool("r", false, "Comando para criação de resources")
		seed           = flag.Bool("s", false, "Comando para criação do seeder")
		migration      = flag.Bool("M", false, "Comando para criação da migration")
		controller     = flag.Bool("c", false, "Comando para criação do controller")
		all            = flag.Bool("a", false, "Comando -a habilitara os demais comandos, menos -repo, -uuid e -id")
		repo           = flag.Bool("repo", false, "Comando para criação de repository pattern")
		createRepoPath = flag.Bool("create-path", false, "Comando para que caso o caminho do repository pattern não exista, seja criado")
	)

	flag.Parse()

	flags := make(map[string]bool)

	flags["m"] = true
	flags["uuid_use"] = *uuidUse
	flags["id_use"] = *idUse
	flags["requests"] = *requests
	flags["resource"] = *resource
	flags["seed"] = *seed
	flags["M"] = *migration
	flags["controller"] = *controller
	flags["repo"] = *repo
	flags["create_repo_path"] = *createRepoPath

	if *all {
		for key := range flags {
			if key == "uuid_use" || key == "id_use" {
				continue
			}

			flags[key] = true
		}
	}

	if *repo {
		flags["routes"] = true
		flags["controller"] = true
		flags["service"] = true
	}

	options := scaffold.Options{
		Name:    *modelFlag,
		Command: flags,
		RootDir: *path,
	}

	if err := scaffold.Run(options); err != nil {
		log.Fatal(err)
	}

	log.Println("PASSOU SEM ERROS")
}
