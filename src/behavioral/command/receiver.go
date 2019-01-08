package command

type Chef struct {}

func (c *Chef) Execute(d string) string {
	return "executed dish " + d
}
