package strategy

type Context struct {
	strategy Strategist
}

func (ctx *Context) Do(from, to string) string {
	return ctx.strategy.Execute(from, to)
}

func (ctx *Context) SetStrategy(strategy Strategist) {
	ctx.strategy = strategy
}




