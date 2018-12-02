package cmds

import (
	v "github.com/appscode/go/version"
	"github.com/appscode/kutil/tools/cli"
	"github.com/golang/glog"
	"github.com/kubevault/csi-driver/cmds/options"
	"github.com/kubevault/csi-driver/driver"
	"github.com/spf13/cobra"
)

func NewCmdInit() *cobra.Command {
	cfg := options.NewConfig()
	cmd := &cobra.Command{
		Use:               "init",
		Short:             "Initializes the driver.",
		DisableAutoGenTag: true,
		PreRun: func(c *cobra.Command, args []string) {
			cli.SendPeriodicAnalytics(c, v.Version.Version)
		},
		Run: func(cmd *cobra.Command, args []string) {
			drv, err := driver.NewDriver(cfg.Endpoint, cfg.NodeName)
			if err != nil {
				glog.Fatalln(err)
			}

			if err := drv.Run(); err != nil {
				glog.Fatalln(err)
			}
		},
	}
	cfg.AddFlags(cmd.Flags())
	return cmd
}
