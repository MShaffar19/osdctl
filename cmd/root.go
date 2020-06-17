package cmd

import (
	"flag"
	"github.com/openshift/osd-utils-cli/cmd/list"
	"os"

	awsv1alpha1 "github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1"
	"github.com/spf13/cobra"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/kubectl/pkg/util/templates"
)

func init() {
	awsv1alpha1.AddToScheme(scheme.Scheme)

	NewCmdRoot(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
}

// rootCmd represents the base command when called without any subcommands
func NewCmdRoot(streams genericclioptions.IOStreams) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "osd-utils-cli",
		Short: "OSD CLI",
		Long:  `CLI tool to provide OSD related utilities`,
		Run:   help,
	}

	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)

	// Reuse kubectl global flags to provide namespace, context and credential options
	kubeFlags := genericclioptions.NewConfigFlags(false)
	kubeFlags.AddFlags(rootCmd.PersistentFlags())

	// add sub commands
	rootCmd.AddCommand(newCmdReset(streams, kubeFlags))
	rootCmd.AddCommand(newCmdSet(streams, kubeFlags))
	rootCmd.AddCommand(list.NewCmdList(streams, kubeFlags))

	// add options command to list global flags
	templates.ActsAsRootCommand(rootCmd, []string{"options"})
	rootCmd.AddCommand(newCmdOptions(streams))

	return rootCmd
}

func help(cmd *cobra.Command, _ []string) {
	cmd.Help()
}