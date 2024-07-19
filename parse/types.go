package parse

type App struct {
	Code        string
	Description string
	Commands    []Command
}

type Command struct {
	Code        string
	Description string
	aliases     []string
	Children    []ChildCommand
}

type ChildCommand struct {
	Code        string
	Description string
	aliases     []string
	Parameters  []Parameter
}

type Parameter struct {
	Code        string
	Description string
	Optional    bool
	Flag        bool
}

type helpInfo struct {
	code         string
	description  string
	usageSuffix  string
	childrenName string
	children     []helpInfo
}

type parseOutput struct {
	helpMessage string
}
