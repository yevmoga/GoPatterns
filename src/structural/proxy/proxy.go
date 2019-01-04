package proxy

import "fmt"

func Do() {
	proxy := CreateProxy(&Service{})
	fmt.Println(proxy.Get())
}
