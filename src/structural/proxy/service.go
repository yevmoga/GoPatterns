package proxy

type Servicer interface {
	Get() string
}

type Service struct {}

func (s *Service) Get() string {
	return "you call service"
}

type Proxy struct {
	Service Servicer
}

func (proxy *Proxy) Get() string {
	return proxy.Service.Get() + " with some proxy"
}

func CreateProxy(service Servicer) Servicer {
	return &Proxy{
		Service: service,
	}
}






