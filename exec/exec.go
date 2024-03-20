package exec

const (
	Git = iota
	CommandLast
)

func GetCommandName(command int) string {
	switch command {
	case Git:
		return "git"
	}
	return ""
}
