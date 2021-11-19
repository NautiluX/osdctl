package account

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"

	"github.com/openshift/osdctl/cmd/account/get"
	"github.com/openshift/osdctl/cmd/account/list"
	"github.com/openshift/osdctl/cmd/account/mgmt"
	"github.com/openshift/osdctl/cmd/account/servicequotas"
	"github.com/openshift/osdctl/pkg/k8s"
)

// NewCmdAccount implements the base account command
func NewCmdAccount(streams genericclioptions.IOStreams, flags *genericclioptions.ConfigFlags) *cobra.Command {
	accountCmd := &cobra.Command{
		Use:               "account",
		Short:             "AWS Account related utilities",
		Args:              cobra.NoArgs,
		DisableAutoGenTag: true,
		Run:               help,
	}

	client, err := k8s.NewClient(flags)
	if err != nil {
		panic(err)
	}
	accountCmd.AddCommand(get.NewCmdGet(streams, flags, client))
	accountCmd.AddCommand(list.NewCmdList(streams, flags, client))
	accountCmd.AddCommand(servicequotas.NewCmdServiceQuotas(streams, flags))
	accountCmd.AddCommand(mgmt.NewCmdMgmt(streams, flags))
	accountCmd.AddCommand(newCmdReset(streams, flags, client))
	accountCmd.AddCommand(newCmdSet(streams, flags, client))
	accountCmd.AddCommand(newCmdConsole(streams, flags))
	accountCmd.AddCommand(newCmdCli(streams, flags))
	accountCmd.AddCommand(newCmdCleanVeleroSnapshots(streams))
	accountCmd.AddCommand(newCmdVerifySecrets(streams, flags, client))
	accountCmd.AddCommand(newCmdRotateSecret(streams, flags))
	accountCmd.AddCommand(newCmdGenerateSecret(streams, flags))

	return accountCmd
}

func help(cmd *cobra.Command, _ []string) {
	cmd.Help()
}
