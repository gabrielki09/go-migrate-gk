package scaffold

type Options struct {
	Name    string
	Command map[string]bool
	RootDir string
}

type File struct {
	Name       string
	ModuleName string
	RootDir    string
	FilePaths  map[string]string
	FileType   map[string]string
}

type RepoPatternNames struct {
	NormalizedName string
	SnakeName      string
	PascalName     string

	RoutesPackage     string
	ControllerPackage string
	ServicePackage    string
	RepositoryPackage string
	RequestPackage    string
	ResponsePackage   string

	RoutesFuncName     string
	ControllerFuncName string
	ServiceFuncName    string
	RepositoryFuncName string
	RequestFuncName    string
	ResponseFuncName   string
}
