package cli

type ArgSet struct {
	args map[int]*Arg
	keys []int
}

func (cmd *Command) Args() *ArgSet {
	if cmd.argSet == nil {
		cmd.argSet = &ArgSet{}
		cmd.argSet.args = make(map[int]*Arg)
		cmd.argSet.keys = make([]int, 0)
	}

	return cmd.argSet
}

type Arg struct {
	Name        string // name as it appears on command line
	value       string // value as set
	description string
}

func (set *ArgSet) Get(name string) string {
	value := ""

	for _, arg := range set.args {
		if arg.Name == name {
			value = arg.value
			break
		}
	}

	return value
}

// TODO: change from ... to a struct { name, description etc}.
func (set *ArgSet) Set(name string, params ...string) {
	key := len(set.args)
	set.keys = append(set.keys, key)

	description := ""
	if len(params) > 0 {
		description = params[0]
	}

	set.args[key] = &Arg{Name: name, description: description}
}

func (set *ArgSet) Length() int {
	return len(set.args)
}
