package main

type Button interface {
	OnClick() string
	OnDoubleClick() string
}
