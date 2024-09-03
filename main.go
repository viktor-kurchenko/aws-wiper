package main

import (
	"time"

	"viktor-kurchenko/aws-wiper/internal"

	"github.com/spf13/cobra"
)

var (
	clusters   *[]string
	tries      *int
	retryPause *time.Duration
	cmd        = &cobra.Command{
		Use:   "aws-wiper",
		Short: "AWS wiper tool",
		Run: func(cmd *cobra.Command, args []string) {
			if err := internal.Wipe(*clusters, *tries, *retryPause); err != nil {
				panic(err)
			}
		},
	}
)

func main() {
	_ = cmd.Execute()
}

func init() {
	clusters = cmd.Flags().StringSliceP("clusters", "c", []string{}, "Comma-separated list (e.g.: us-west-1=eks1,us-west-2=eks2)")
	_ = cmd.MarkFlagRequired("clusters")
	tries = cmd.Flags().IntP("tries", "t", 5, "Cleanup tries count")
	retryPause = cmd.Flags().DurationP("pause", "p", time.Minute, "Cleanup retry pause")
}
