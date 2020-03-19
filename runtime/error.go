package runtime

/*
* @
* @Author:
* @Date: 2020/3/19 16:15
 */

// An errorString represents a runtime error described by a single string.
type errorString string

func (e errorString) RuntimeError() {}

func (e errorString) Error() string {
	return "runtime error: " + string(e)
}