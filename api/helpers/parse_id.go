package helpers

import (
	"strconv"
	"strings"
)

/*
parse id from a url path. p is the path.
pos is [1, ...). it stands for the position of the id
in the path. pass -1 for 'last'
*/
func ParseIDFromPath(p string, pos int) (int, error) {
	// if we split /api/v1/bla
	// we get a 4 elements long slice
	//
	// if the first character of p is not '/'
	// then decrement pos by 1 to make sure the index
	// is not out of bounds
	if string(p[0]) != "/" {
		pos = pos - 1
	}
	if pos == 0 {
		pos = 1
	}
	splitted := strings.Split(p, "/")
	if pos == -1 {
		pos = len(splitted) - 1
	}
	idStr := splitted[pos]
	idI64, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(idI64), nil
}
