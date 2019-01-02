package bridge

import "fmt"

func Do()  {
	app := AppConstructor(&Xml{})
	fmt.Println(app.GetData())
}
