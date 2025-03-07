package cli

import (
	"fmt"
	"github.com/faelmori/logz"
	. "github.com/faelmori/xtui/wrappers"
	"github.com/spf13/cobra"
	"strings"
)

func AppsCmdsList() []*cobra.Command {
	return []*cobra.Command{
		InstallApplicationsCommand(),
	}
}

func InstallApplicationsCommand() *cobra.Command {
	var depList []string
	var path string
	var yes, quiet bool

	cmd := &cobra.Command{
		Use:     "app-install",
		Aliases: []string{"install", "appInstall", "installApp", "aptInstall", "apt-install", "depInstall", "dep-install"},
		Short:   "Install applications and dependencies",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(depList) == 0 && len(args) == 0 {
				logz.Error("Empty applications list", map[string]interface{}{
					"context": "InstallApplicationsCommand",
					"error":   "no applications to install",
				})
				return fmt.Errorf("no applications to install")

			}
			newArgs := []string{strings.Join(depList, " "), path, fmt.Sprintf("%t", yes), fmt.Sprintf("%t", quiet)}
			args = append(args, newArgs...)

			availableProperties := getAvailableProperties()
			if len(availableProperties) > 0 {
				adaptedArgs := adaptArgsToProperties(args, availableProperties)
				return InstallDependenciesWithUI(adaptedArgs...)
			}

			return InstallDependenciesWithUI(args...)
		},
	}

	cmd.Flags().StringArrayVarP(&depList, "application", "a", []string{}, "Applications list to install")
	cmd.Flags().StringVarP(&path, "path", "p", "", "Apps installation path")
	cmd.Flags().BoolVarP(&yes, "yes", "y", false, "Automatic yes to prompts")
	cmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Quiet mode")

	return cmd
}

func getAvailableProperties() map[string]string {
	return map[string]string{
		"property1": "value1",
		"property2": "value2",
	}
}

func adaptArgsToProperties(args []string, properties map[string]string) []string {
	adaptedArgs := args
	for key, value := range properties {
		adaptedArgs = append(adaptedArgs, fmt.Sprintf("--%s=%s", key, value))
	}
	return adaptedArgs
}
