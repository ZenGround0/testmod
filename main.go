package testmod

import "fmt"

// Hi returns an unfriendly greeting
func Hi(name string, caps bool) string {
	if caps {
		return fmt.Sprintf("PLEASE STOP THE HATE, %s", name)
	}
	return fmt.Sprintf("Please stop the hate, %s", name)
}
