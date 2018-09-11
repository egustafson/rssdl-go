package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "rssdl",
	Short: "RSS-dl is an RSS (Podcast) downloader",
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "conf", "", "config file (default is ./rssdl.yaml)")

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(dlCmd)
}

func initConfig() {
	if cfgFile != "" {
		// if set through the rootCmd.PersistentFlags
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("rssdl")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config file: ", err)
		os.Exit(1)
	}
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize (or reset) the database.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init: " + cfgFile)
	},
}

var listCmd = &cobra.Command{
	Use:   "list [feed-id]",
	Short: "list the feeds or list articles in a feed.",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list: " + cfgFile)
	},
}

var addCmd = &cobra.Command{
	Use:   "add [RSS-URL]",
	Short: "Add an RSS feed (URL) for tracking.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add: " + cfgFile)
	},
}

var updateCmd = &cobra.Command{
	Use:   "update [feed-id]",
	Short: "Update all, or a single feed by ingesting the RSS.",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update: " + cfgFile)
	},
}

var dlCmd = &cobra.Command{
	Use:   "dl [article-id]",
	Short: "Download the content from an article.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dl: " + cfgFile)
	},
}
