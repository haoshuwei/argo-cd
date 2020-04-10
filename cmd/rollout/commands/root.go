package commands

import (
	"github.com/argoproj/argo-cd/rollout"
	"github.com/spf13/cobra"
)

func NewRolloutCommand() *cobra.Command {

	var (
		clientOptions rollout.ClientOptions
	)

	var command = &cobra.Command{
		Use:   "rollout",
		Short: "appcenter rollout app",
		Run: func(c *cobra.Command, args []string) {
			c.HelpFunc()(c, args)
		},
	}

	command.AddCommand(NewCanaryCommand(&clientOptions))


	command.PersistentFlags().StringVar(&clientOptions.Kubeconfig, "kubeconfig", "~/.kube/config", "Kubeconfig file location")
	command.PersistentFlags().StringVar(&clientOptions.Namespace, "namespace", "default", "Canary namespace")
	return command
}
