//go:build !windows

package input_autocomplete

// We make the assumption that on non-Windows OSes support ANSI escape
// sequences by default.

func EnableVirtualTerminalWindows() error {
	return nil
}
