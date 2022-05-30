package algorithms

import (
	dse "dsa/general"
	"fmt"
)

func LinearSearch(list []int, ele int) (string, error) {

	stat := false
	msg := ""
	errmsg := fmt.Sprintf("%d does not found in given list", ele)

	for i, e := range list {
		if e == ele {
			stat = true
			msg = fmt.Sprintf("%d found at location %d", ele, i)
			break
		}
	}

	if !stat {
		return msg, &dse.DsaError{Err: errmsg}
	}

	return msg, nil
}
