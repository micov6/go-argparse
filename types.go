package parse

type App struct {
	Code        string
	Description string
	Commands    []*Command
}

type Command struct {
	Code        string
	Description string
	aliases     []string
	Children    []*ChildCommand
}

type ChildCommand struct {
	Code           string
	Description    string
	CommandHandler func(parameterValues map[string]ParameterValue) error
	Parameters     []*Parameter
	aliases        []string
}

type Parameter struct {
	Code        string
	Description string
	IsOptional  bool
	IsNumber    bool
	IsBoolean   bool
	aliases     []string // NOTE: do not support for arg parsing, just for help
}

type ParameterValue struct {
	StringValue  string
	NumberValue  int
	BooleanValue bool
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
