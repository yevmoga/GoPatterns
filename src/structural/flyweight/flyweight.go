package flyweight

type Flyweighter interface {
	GetName() string
	SetName(string)
}

type FlyweightFactory struct {
	pool map[int]Flyweighter
}

func (f *FlyweightFactory) GetFlyweight(state int) Flyweighter {
	if f.pool == nil {
		f.pool = make(map[int]Flyweighter)
	}

	if _, ok := f.pool[state]; !ok {
		f.pool[state] = &Flyweight{state:state}
	}

	return f.pool[state]
}

type Flyweight struct {
	name  string
	state int
}

func (f *Flyweight) GetName() string {
	return f.name
}

func (f *Flyweight) SetName(name string) {
	f.name = name
}
