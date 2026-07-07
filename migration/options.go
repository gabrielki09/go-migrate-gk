package migration

type Command string

const (
	CommandUp     Command = "up"
	CommandDown   Command = "down"
	CommandFresh  Command = "fresh"
	CommandStatus Command = "status"
)

type Options struct {
	Dir     string
	Command Command
}

type Migration struct {
	Version  string
	Name     string
	UpFile   string
	DownFile string
}
