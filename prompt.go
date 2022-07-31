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
		"🎉 初めてのコミット",
		"🔖 バージョンタグ",
		"✨ 新機能",
		"🐛 バグ修正",
		"♻️  リファクタリング",
		"📚 ドキュメント",
		"🎨 デザインUI/UX",
		"🐎 パフォーマンス",
		"🔧 ツール",
		"🚨 テスト",
		"💩 非推奨追加",
		"🗑️ 削除",
		"🚧 WIP"}

	validate := func(input string) error {
		// if err != nil {
		// 	return errors.New("Invalid number")
		// }
		return nil
	}

	commitType := promptui.Select{
		// 選択肢のタイトル
		Label: "コミットタイプを変更",
		// 選択肢の配列
		Items: prefixs,
	}

	idx, _, err := commitType.Run()

	if err != nil {
		fmt.Println(err)
		// os.Exit(1)
	}

	summaryPrompt := promptui.Prompt{
		Label:    "変更の概要を追加",
		Validate: validate,
	}
	summary, err := summaryPrompt.Run()
	// fmt.Println(summary)

	commitPrompt := promptui.Prompt{
		Label:    "コミットメッセージを追加",
		Validate: validate,
	}
	commitMsg, err := commitPrompt.Run()
	fmt.Println(commitMsg)

	breakingChangeConfirm := promptui.Prompt{
		Label:     "破壊的変更がありますか？",
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
			Label:    "破壊的変更に対する説明を追加",
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
		Label:     "Issueに関係するコミットですか？",
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
			Label:    "関連するIssue番号を追加(例: #1, #2...)",
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
