package gpu

import (
	"log"

	"github.com/beclab/Olares/cli/pkg/pipelines"
	"github.com/spf13/cobra"
)

func NewCmdEnableGpu() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enable",
		Short: "Enable GPU drivers for Olares node",
		Run: func(cmd *cobra.Command, args []string) {
			if err := pipelines.EnableGpuNode(); err != nil {
				log.Fatalf("error: %v", err)
			}
		},
	}
	return cmd
}
