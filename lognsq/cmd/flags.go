package cmd

import "github.com/spf13/cobra"

func topicFlag(cmd *cobra.Command, cfgd *cfgData) {
	s, _ := cmd.Flags().GetString("topic")
	if s != "" {
		cfgd.Topic = s
	}
}

func addressFlag(cmd *cobra.Command, cfgd *cfgData) {
	s, _ := cmd.Flags().GetString("nsqd-tcp-address")
	if s != "" {
		cfgd.Address = s
	}
}

func regexFlag(cmd *cobra.Command, cfgd *cfgData) {
	s, _ := cmd.Flags().GetString("regex-filter")
	if s != "" {
		cfgd.Regex = s
	}
}

func printFlag(cmd *cobra.Command, cfgd *cfgData) {
	b, _ := cmd.Flags().GetBool("print-log")
	if b {
		cfgd.StderrLogs = true
	}
}
