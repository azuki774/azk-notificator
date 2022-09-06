package main

import (
	"azk-notificator/internal/factory"
	"fmt"

	"github.com/spf13/cobra"
)

var runOption factory.ServerRunOption

// startCmd represents the start command

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		fmt.Println("start called")
		srv, err := factory.NewServer(&runOption)
		if err != nil {
			return err
		}

		return srv.Start()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	startCmd.Flags().StringVar(&runOption.Host, "host", "", "listen host")
	startCmd.Flags().StringVar(&runOption.Port, "port", "80", "listen port")
	startCmd.Flags().StringVar(&runOption.QueueHost, "queue-host", "localhost", "queue-DB host")
	startCmd.Flags().StringVar(&runOption.QueuePort, "queue-port", "6379", "queue-DB port")
	startCmd.Flags().StringVar(&runOption.QueuePort, "queue-pass", "6379", "queue-DB password")
}
