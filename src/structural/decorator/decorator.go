package decorator

type Provider interface{
	Send(string) string
}

func Do(provider Provider) string {

	return provider.Send("Hello, ")

}
