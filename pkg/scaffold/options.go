package scaffold

type Options struct {
	Name             string
	SeparateByFolder bool
	Command          map[string]bool
}
type File struct {
	Name             string
	FilePaths        map[string]string
	Content          string
	SeparateByFolder bool
	FileType         map[string]string
}
