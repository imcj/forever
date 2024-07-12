package main

import (
	"forever/task"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	rootCmd := &cobra.Command{
		Use: "forever",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	runCmd := &cobra.Command{
		Use: "run",
		Run: func(cmd *cobra.Command, args []string) {
			var config string
			cmd.PersistentFlags().StringVarP(&config, "config", "c", "", "config file (default is ./forever.yml)")

			if config == "" {
				config = "./forever.yml"
			}

			cfg, err := task.LoadConfig(config)
			if err != nil {
				logrus.Errorf("load config err: %v", err)
				return
			}

			runner, err := task.NewRunner(cfg)
			if err != nil {
				logrus.Errorf("new runner err: %v", err)
				return
			}

			runner.Run()
		},
	}

	rootCmd.AddCommand(runCmd)
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
