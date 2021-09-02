package helpers

import "strconv"

func StrToInt(str string) (int, error) {
	int64V, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	intV := int(int64V)
	return intV, nil
}
