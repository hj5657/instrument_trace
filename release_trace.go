//go:build !debug
// +build !debug

package trace

func Trace() func() {
	return func() {

	}
}
