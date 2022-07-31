package main

import "fmt"
import "os"

func main() {
	result := RunPrompt()

	msg, err := genCommitMsg(result)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(msg)

	err = Commit(msg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
