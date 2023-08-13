package simplestring

import (
	"fmt"
	"strings"
)

const FirstByte string = "+"

func Serialize(response string) string {
	// string that cannot contain a CR or LF character (no newlines are allowed)
	response = strings.ReplaceAll(response, "\r", "_")
	response = strings.ReplaceAll(response, "\n", "_")

	return fmt.Sprintf("%s%s\r\n", FirstByte, response)
}
