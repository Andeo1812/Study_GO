package special_print

import (
	"fmt"
	"time"
)

func Print() {
	today := time.Now()

	message := "Today is"

	output := fmt.Sprintf("Message - %s : %s\n", message, today)

	print(output)
}
