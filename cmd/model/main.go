package main

import (
	"flag"
	"log"
	"strings"

	"github.com/gabrielki09/go-migrate-gk/pkg/model"
)

func main() {
	var (
		modelFlag        = flag.String("model", "", "Comando para criação de arquivo padrão da model")
		uuidUse          = flag.Bool("uuid", false, "Comando para criação do model com uuid")
		idUse            = flag.Bool("id", false, "Comando para criação do model com id (int)")
		separateByFolder = flag.Bool("S", false, "Comando para separação de pastas por model")
		requests         = flag.Bool("R", false, "Comando para criação da request")
		resource         = flag.Bool("r", false, "Comando para criação de resources")
		seed             = flag.Bool("s", false, "Comando para criação do seeder")
		migration        = flag.Bool("m", false, "Comando para criação da migration")
		controller       = flag.Bool("c", false, "Comando para criação do controller")
		all              = flag.Bool("a", false, "Comando para separação de pastas por model")
	)

	flag.Parse()

	flags := make(map[string]bool)

	if strings.TrimSpace(*modelFlag) == "" {
		log.Fatal("flag model é obrigatória.")
	}

	flags["model"] = true
	flags["uuid_use"] = *uuidUse
	flags["id_use"] = *idUse
	flags["separate_by_folder"] = *separateByFolder
	flags["requests"] = *requests
	flags["resource"] = *resource
	flags["seed"] = *seed
	flags["migration"] = *migration
	flags["controller"] = *controller

	if *all {
		for key := range flags {
			if key == "uuid_use" || key == "id_use" {
				continue
			}

			flags[key] = true
		}
	}

	options := model.Options{
		Name:             *modelFlag,
		SeparateByFolder: flags["separate_by_folder"],
		Command:          flags,
	}

	if err := model.Run(options); err != nil {
		log.Fatal(err)
	}
}
