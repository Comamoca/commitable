package main

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

type PromptResult struct {
	Prefix               string
	Summary              string
	BreakingChangeMsg, h string
	RelatedIssue         string
}

func RunPrompt() PromptResult {
	// return prefix, commitMsg,

	prefixs := []string{
		"ð åãã¦ã®ã³ããã",
		"ð ãã¼ã¸ã§ã³ã¿ã°",
		"â¨ æ°æ©è½",
		"ð ãã°ä¿®æ­£",
		"â»ï¸  ãªãã¡ã¯ã¿ãªã³ã°",
		"ð ãã­ã¥ã¡ã³ã",
		"ð¨ ãã¶ã¤ã³UI/UX",
		"ð ããã©ã¼ãã³ã¹",
		"ð§ ãã¼ã«",
		"ð¨ ãã¹ã",
		"ð© éæ¨å¥¨è¿½å ",
		"ðï¸ åé¤",
		"ð§ WIP"}

	validate := func(input string) error {
		// if err != nil {
		// 	return errors.New("Invalid number")
		// }
		return nil
	}

	commitType := promptui.Select{
		// é¸æè¢ã®ã¿ã¤ãã«
		Label: "ã³ãããã¿ã¤ããå¤æ´",
		// é¸æè¢ã®éå
		Items: prefixs,
	}

	idx, _, err := commitType.Run()

	if err != nil {
		fmt.Println(err)
		// os.Exit(1)
	}

	summaryPrompt := promptui.Prompt{
		Label:    "å¤æ´ã®æ¦è¦ãè¿½å ",
		Validate: validate,
	}
	summary, err := summaryPrompt.Run()
	// fmt.Println(summary)

	commitPrompt := promptui.Prompt{
		Label:    "ã³ãããã¡ãã»ã¼ã¸ãè¿½å ",
		Validate: validate,
	}
	commitMsg, err := commitPrompt.Run()
	fmt.Println(commitMsg)

	breakingChangeConfirm := promptui.Prompt{
		Label:     "ç ´å£çå¤æ´ãããã¾ããï¼",
		IsConfirm: true,
	}
	var isBreakingChange string
	isBreakingChange, err = breakingChangeConfirm.Run()
	if err != nil {
		fmt.Println(err)
		// os.Exit(1)
	}
	// fmt.Println(isBreakingChange)

	var breakingChangeMsg string
	if isBreakingChange == "y" {
		breakingCangePrompt := promptui.Prompt{
			Label:    "ç ´å£çå¤æ´ã«å¯¾ããèª¬æãè¿½å ",
			Validate: validate,
		}
		breakingChangeMsg, err = breakingCangePrompt.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// fmt.Println(breakingChangeMsg)
	}

	relatedIssueConfirm := promptui.Prompt{
		Label:     "Issueã«é¢ä¿ããã³ãããã§ããï¼",
		IsConfirm: true,
	}
	var isRelatedIssue string
	isRelatedIssue, err = relatedIssueConfirm.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// fmt.Println(isBreakingChange)

	var relatedIssue string
	if isRelatedIssue == "y" {
		relatedIssuePrompt := promptui.Prompt{
			Label:    "é¢é£ããIssueçªå·ãè¿½å (ä¾: #1, #2...)",
			Validate: validate,
		}
		relatedIssue, err = relatedIssuePrompt.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// fmt.Println(relatedIssue)
	}

	result := PromptResult{
		Prefix:            prefixs[idx],
		Summary:           summary,
		BreakingChangeMsg: breakingChangeMsg,
		RelatedIssue:      relatedIssue,
	}

	return result
}
