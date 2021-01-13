package branch

import (
	"fmt"

	"github.com/planetscale/cli/config"
	"github.com/spf13/cobra"
)

// StatusCmd gets the status of a database branch using the PlanetScale API.
func StatusCmd(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status <db_name> <branch_name>",
		Short: "Check the status of a branch of a database",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return cmd.Usage()
			}

			_, err := cfg.NewClientFromConfig()
			if err != nil {
				return err
			}

			// TODO: Call PlanetScale API here to get database branch status.

			fmt.Println("database branch status has not been implemented yet")
			return nil
		},
	}

	return cmd
}