package app

import (
	"context"
	"time"

	"github.com/spf13/cobra"

	"github.com/flyer103/demo-operator/pkg/operator"
)

var (
	kubeconfig     string
	watchNamespace string
	resyncSeconds  uint32
)

var serverCmd = &cobra.Command{
	Use:           "server",
	Short:         "Launch server",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		config := &operator.OperatorConfig{
			KubeConfigPath: kubeconfig,
			WatchNamespace: watchNamespace,
			ResyncPeriod:   time.Duration(resyncSeconds) * time.Second,
		}
		operator, err := operator.NewOperator(config)
		if err != nil {
			return err
		}

		ctx := context.TODO()
		stopc := make(chan struct{})

		return operator.Run(ctx, stopc)
	},
}

func init() {
	serverCmd.Flags().StringVarP(&kubeconfig, "kubeconfig", "c", "", "path to kube config")
	serverCmd.Flags().StringVar(&watchNamespace, "watchNamespace", "",
		"the namespace which Operator watches")
	serverCmd.Flags().Uint32Var(&resyncSeconds, "resyncSeconds", 30,
		"resync seconds")

	rootCmd.AddCommand(serverCmd)
}
