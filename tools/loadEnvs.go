package tools

import (
	"os"
	"strings"
)

func ParseEnvFile(filename string) (envArray []string) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	textSplit := strings.Split(string(file), "\n")
	var countArray = int(0)
	for _, splittedText := range textSplit {
		// Try to check content
		if strings.Contains(splittedText, "=") && !strings.HasPrefix(splittedText, "=") {
			countArray++
			envArray = append(envArray, splittedText)
		} else {
			if countArray == 0 {
				ZapLogger("both").Info("Trying to add " + splittedText + " as an env but countArray is zero")
				continue
			} else if strings.HasPrefix(splittedText, "#") {
				ZapLogger("both").Info("Trying to add " + splittedText + " as an env but this env is commented")
				continue
			} else if splittedText == "" {
				ZapLogger("both").Info("Trying to add " + splittedText + " as an env but this line is empty")
				continue
			} else {
				envArray[countArray-1] = envArray[countArray-1] + splittedText
			}
		}
	}

	return envArray
}
