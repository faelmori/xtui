package cli

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/faelmori/xtui/components"
	"github.com/faelmori/xtui/types"
	"github.com/faelmori/xtui/wrappers"
	"github.com/spf13/cobra"
	"testing"
	"time"
)

func FormsCmdsList() []*cobra.Command {
	inputCmd := InputFormCommand()
	loaderCmd := LoaderFormCommand()

	return []*cobra.Command{
		inputCmd,
		loaderCmd,
	}
}

func InputFormCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "input-form",
		Aliases: []string{"input", "formInput", "inputForm", "formInput", "form-input"},
		Short:   "Form inputs for any command",
		Long:    "Form inputs screen, interactive mode, for any command with flags",
		RunE: func(cmd *cobra.Command, args []string) error {
			config := types.Config{
				Title: "Dynamic Form",
				Fields: types.FormFields{
					Title: "Login",
					Fields: []types.FormField{
						&types.InputField{
							Ph:  "Username",
							Tp:  "text",
							Val: "",
							Req: true,
							Min: 3,
							Max: 20,
							Err: "Username is required and must be between 3 and 20 characters.",
							Vld: func(value string) error {
								if len(value) < 3 || len(value) > 20 {
									return fmt.Errorf("username must be between 3 and 20 characters")
								}
								return nil
							},
						},
						&types.InputField{
							Ph:  "Password",
							Tp:  "password",
							Val: "",
							Req: true,
							Min: 6,
							Max: 20,
							Err: "Password is required and must be between 6 and 20 characters.",
							Vld: func(value string) error {
								if len(value) < 6 || len(value) > 20 {
									return fmt.Errorf("password must be between 6 and 20 characters")
								}
								return nil
							},
						},
					},
				},
			}
			_, err := components.ShowForm(config)
			return err
		},
	}

	return cmd
}

func LoaderFormCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "loader-form",
		Aliases: []string{"loader", "formLoader", "loaderForm", "formLoader", "form-loader"},
		Short:   "Form loader for any command",
		Long:    "Form loader screen, interactive mode, for any command with flags",
		RunE: func(cmd *cobra.Command, args []string) error {
			messages := make(chan tea.Msg)
			go func() {
				messages <- wrappers.LoaderMsg{Message: "Loading dynamic properties..."}
				time.Sleep(2 * time.Second)
				messages <- wrappers.LoaderMsg{Message: "Dynamic properties loaded successfully."}
				time.Sleep(1 * time.Second)
				messages <- wrappers.LoaderCloseMsg{}
			}()
			return wrappers.StartLoader(messages)
		},
	}

	return cmd
}

// Unit tests for InputFormCommand
func TestInputFormCommand(t *testing.T) {
	cmd := InputFormCommand()
	if cmd.Use != "input-form" {
		t.Errorf("expected 'input-form', got '%s'", cmd.Use)
	}
	if cmd.Short != "Form inputs for any command" {
		t.Errorf("expected 'Form inputs for any command', got '%s'", cmd.Short)
	}
	if cmd.Long != "Form inputs screen, interactive mode, for any command with flags" {
		t.Errorf("expected 'Form inputs screen, interactive mode, for any command with flags', got '%s'", cmd.Long)
	}
}

// Unit tests for LoaderFormCommand
func TestLoaderFormCommand(t *testing.T) {
	cmd := LoaderFormCommand()
	if cmd.Use != "loader-form" {
		t.Errorf("expected 'loader-form', got '%s'", cmd.Use)
	}
	if cmd.Short != "Form loader for any command" {
		t.Errorf("expected 'Form loader for any command', got '%s'", cmd.Short)
	}
	if cmd.Long != "Form loader screen, interactive mode, for any command with flags" {
		t.Errorf("expected 'Form loader screen, interactive mode, for any command with flags', got '%s'", cmd.Long)
	}
}
