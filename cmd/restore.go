package cmd

import (
	"github.com/alv91/vault-backup/internal/app"
	"github.com/alv91/vault-backup/internal/pkg/s3"
	"github.com/alv91/vault-backup/internal/pkg/vault"
	"github.com/spf13/cobra"
)

var forceRestore bool

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a vault backup from raft snapshot",
	Run: func(cmd *cobra.Command, args []string) {

		vaultCfg := &vault.Config{
			Token:        vaultToken,
			Address:      vaultAddr,
			Namespace:    vaultNamespace,
			ForceRestore: forceRestore,
		}

		s3Cfg := &s3.Client{
			AccessKey:       s3AccessKey,
			SecretAccessKey: s3SecretKey,
			Region:          s3Region,
			Bucket:          s3Bucket,
			Endpoint:        s3Endpoint,
			FileName:        s3FileName,
		}

		app.Restore(vaultCfg, s3Cfg)
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)

	restoreCmd.Flags().BoolVarP(&forceRestore, "force", "f", false, "force restore")
}
