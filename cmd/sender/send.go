package main

import (
	"azk-notificator/internal/factory"
	"fmt"

	"github.com/spf13/cobra"
)

var sendOption factory.SenderRunOption

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		fmt.Println("send called")
		sender := factory.NewSender(&sendOption)
		defer sender.Logger.Sync()
		return sender.Run()
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	sendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	sendCmd.Flags().StringVar(&sendOption.QueueHost, "queue-host", "localhost", "queue-DB host")
	sendCmd.Flags().StringVar(&sendOption.QueuePort, "queue-port", "6379", "queue-DB port")
	sendCmd.Flags().StringVar(&sendOption.QueuePass, "queue-pass", "", "queue-DB password")
}
