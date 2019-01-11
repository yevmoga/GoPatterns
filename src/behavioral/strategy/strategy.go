package strategy

import "fmt"

type Strategist interface {
	Execute(from, to string) string
}

type WalkStrategy struct{}

func (*WalkStrategy) Execute(from, to string) string {
	return fmt.Sprintf("please, go form %s to %s using feet", from, to)
}

type BicycleStrategy struct{}

func (*BicycleStrategy) Execute(from, to string) string {
	return fmt.Sprintf("please, go form %s to %s using bicycle", from, to)
}

type CarStrategy struct{}

func (*CarStrategy) Execute(from, to string) string {
	return fmt.Sprintf("please, go form %s to %s using car", from, to)
}
