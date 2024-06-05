package util

import "github.com/manifoldco/promptui"

// Create promptui Input Prompt
func PromptInput(label string) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
	}
	return prompt.Run()
}

// Create promptui Input Prompt with Default Value
func PromptInputWithDefault(label, defaultValue string) (string, error) {
	prompt := promptui.Prompt{
		Label:   label,
		Default: defaultValue,
	}
	return prompt.Run()
}

// Create promptui Select Items
func PromptSelect(label string, items []string) (string, error) {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}
	_, result, err := prompt.Run()
	return result, err
}

// Create promptui Select Items with Default Value
func PromptSelectWithDefault(label string, items []string, defaultValue string) (string, error) {
	prompt := promptui.Select{
		Label:     label,
		Items:     items,
		CursorPos: indexOf(items, defaultValue),
	}
	_, result, err := prompt.Run()
	return result, err
}

// Find the index of the item from a list of items
func indexOf(items []string, item string) int {
	for i, v := range items {
		if v == item {
			return i
		}
	}
	return 0
}
