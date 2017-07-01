package config

import "github.com/spf13/cobra"

func NewCmdConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "config",
		Short:   "OSM configuration",
		Example: "osm config view",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(newCmdSet())
	cmd.AddCommand(newCmdUse())
	cmd.AddCommand(newCmdView())
	cmd.AddCommand(newCmdGet())
	cmd.AddCommand(newCmdCurrent())
	return cmd
}