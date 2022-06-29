package algorithms

import "reflect"

type BSortList interface {
	Bsort() interface{}
}

type Ibslist struct {
	List interface{}
}

func (ibsl Ibslist) Bsort() interface{} {
	values := reflect.ValueOf(ibsl.List)
	for i := 0; i < values.Len(); i++ {
		swapped := false
		for j := 0; j < values.Len()-1; j++ {
			if values.Index(j).Int() > values.Index(j+1).Int() {
				x, y := values.Index(j), values.Index(j+1)
				values.Index(j).Set(y)
				values.Index(j + 1).Set(x)
				swapped = true
			}
		}

		if swapped {
			break
		}
	}

	return values.Interface()
}

func BubbleSort(bs BSortList) interface{} {
	return bs.Bsort()
}
