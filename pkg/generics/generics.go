package generics

import "fmt"

type comparator interface {
	compare(data *interface{}) int
}

type myData struct {
	h int
}

func (value *myData) compare(data *interface{}) int {
	datai := (*data).(myData)
	return value.h - datai.h
}
func IsCompare() {
	data := myData{
		h: 10,
	}
	var value comparator = &data
	var datai *interface{} = new(interface{})
	*datai = myData{h: 20}

	fmt.Println(value.compare(datai))
}
