package proxy

type Servicer interface {
	Get(string) string
}

type Service struct {}

func (*Service) Get(s string) string {
	return s
}

type Proxy struct {
	Service Servicer
}

func (proxy *Proxy) Get(s string) string {
	return proxy.Service.Get(s) + " with some proxy"
}

func CreateProxy(service Servicer) Servicer {
	return &Proxy{
		Service: service,
	}
}






