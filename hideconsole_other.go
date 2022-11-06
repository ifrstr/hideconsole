//go:build !windows
//+build !windows

package hideconsole

func HideConsole() error {
	return nil
}
