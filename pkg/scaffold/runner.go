package scaffold

func Run(option Options) error {
	if err := option.Validate(); err != nil {
		return err
	}

	dirs, err := resolveFileDir(option.Command, option.RootDir)
	if err != nil {
		return err
	}

	goMod, err := getModuleName()
	if err != nil {
		return err
	}

	fileConfig := File{
		Name:       option.Name,
		FilePaths:  dirs,
		ModuleName: goMod,
		RootDir:    option.RootDir,
	}

	return createFiles(fileConfig, option)

}
