package testmod

import "fmt"

// Hi returns an unfriendly greeting
func Hi(name string) string {
	return fmt.Sprintf("Please stop the hate, %s", name)
}
