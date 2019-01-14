package decorator

type Provider interface{
	Send(string) string
}
