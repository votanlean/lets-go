package main
type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	ans int
}

func (s Simple) Get() int {
	return s.ans
}

func (p *Simple) Set(u int) {
	p.ans = u
}

func fI(it Simpler) int {
	println("init simple %v", it.Get())
	it.Set(10)
	println("after set simple %v", it.Get())
	return it.Get()
}

func main() {
	simple := Simple{
		ans: 1,
	}
	var si Simpler
	si = &simple
	fI(si)
}