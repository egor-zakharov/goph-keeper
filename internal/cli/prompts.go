package cli

import "github.com/c-bata/go-prompt"

func (e *Executor) ShowPrompts(d prompt.Document) []prompt.Suggest {
	var suggests []prompt.Suggest

	if d.FindEndOfCurrentWord() == 0 {
		suggests = []prompt.Suggest{
			{Text: "signUp", Description: "Example: signUp;login;password"},
			{Text: "signIn", Description: "Example: signIn;login;password"},

			{Text: "create-card", Description: "Example: create-card;card_number;card_expire;card_holder;card_cvv"},
			{Text: "get-cards", Description: "Example: get-cards"},
			{Text: "update-card", Description: "Example: update-card;card_id;card_number;card_expire;card_holder;card_cvv"},
			{Text: "delete-card", Description: "Example: delete-card;card_id"},

			{Text: "upload-file", Description: "Example: upload-file;file_path;meta"},
			{Text: "get-files", Description: "Example:get-files"},
			{Text: "download-file", Description: "Example:download-file;file_id;file_name"},
			{Text: "delete-file", Description: "Example:delete-file;file_id"},

			{Text: "create-text", Description: "Example: create-text;meta;text"},
			{Text: "get-text", Description: "Example: get-text"},
			{Text: "update-text", Description: "Example: update-text;text_id;meta;text"},
			{Text: "delete-text", Description: "Example: delete-text;text_id"},

			{Text: "create-auth", Description: "Example: create-auth;meta;text"},
			{Text: "get-auth", Description: "Example: get-auth"},
			{Text: "update-auth", Description: "Example: update-auth;text_id;meta;text"},
			{Text: "delete-auth", Description: "Example: delete-text;text_id"},

			{Text: "exit", Description: "Example: exit"},
		}
	}

	return prompt.FilterHasPrefix(suggests, d.GetWordBeforeCursor(), false)
}
