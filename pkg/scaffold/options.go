package scaffold

type Options struct {
	Name    string
	Command map[string]bool
	RootDir string
}
type File struct {
	Name      string
	FilePaths map[string]string
	Content   string
	FileType  map[string]string
}
