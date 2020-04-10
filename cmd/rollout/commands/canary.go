package commands

import (
	"github.com/argoproj/argo-cd/rollout"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
)

// NewCommand returns a new instance of an appcenter command
func NewCanaryCommand(config *rollout.ClientOptions) *cobra.Command {

	var command = &cobra.Command{
		Use:   "canary",
		Short: "appcenter rollout canary",
		Run: func(c *cobra.Command, args []string) {
			c.HelpFunc()(c, args)
		},
	}

	command.AddCommand(NewCanarySetCommand(config))
	command.AddCommand(NewCanaryListCommand(config))
	return command
}

func NewCanarySetCommand(config *rollout.ClientOptions) *cobra.Command {

	var command = &cobra.Command{
		Use:   "setWeight",
		Short: "appcenter rollout canary set",
		Example: "appcenter rollout canary setWeight <canaryName> 10",
		Run: func(c *cobra.Command, args []string) {
			if len(args) < 2 {
				log.Fatal("Missing canary name or weight")
			}
			controller := rollout.NewCanaryController(config.Kubeconfig)
			weightInt, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatalf("Convert %s to int error %v", args[1], err)
			}
			controller.SetWeight(weightInt, args[0], config.Namespace)
		},
	}

	return command
}

func NewCanaryListCommand(config *rollout.ClientOptions) *cobra.Command {
	var command = &cobra.Command{
		Use:   "list",
		Short: "appcenter list canary",
		Example: "appcenter list canary",
		Run: func(c *cobra.Command, args []string) {
			controller := rollout.NewCanaryController(config.Kubeconfig)
			controller.ListCanary(config.Namespace)
		},
	}

	return command
}
