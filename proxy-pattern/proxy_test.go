package proxy_pattern

import "testing"

func TestProxy(t *testing.T) {
	str1 := "Get this text from real object"
	str2 := "Get this text from proxy"

	p := new(Proxy)
	ro := new(RealObject)
	res1 := ro.DoSomething(str1)
	res2 := p.DoSomething(str2)

	if str1 != res1 {
		t.Errorf("Expect result is %s, but got %s", str1, res1)
	}

	if str2 != res2 {
		t.Errorf("Expect result is %s, but got %s", str2, res2)
	}
}