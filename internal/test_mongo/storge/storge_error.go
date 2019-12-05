package storge

import "strings"

//@author xiaolan
//@lastUpdate 2019-10-08
//@param _err mongodb的错误类型
func IsDuplicate(_err error) bool {

	if strings.HasPrefix(_err.Error(), "E11000") {
		return true
	}
	return false
}


//@author xiaolan
//@lastUpdate 2019-10-09
//@param _err mongodb的错误类型
func IsNotFound(_err error) bool {

	if _err.Error() == "not found" {
		return true
	}
	return false
}
