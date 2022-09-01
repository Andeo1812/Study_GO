package race

import (
	"fmt"
)

func PrintName(name string) {
	message := "My name is"

	output := fmt.Sprintf("%s - %s\n", message, name)

	print(output)
}
