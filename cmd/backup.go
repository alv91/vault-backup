package cmd

import (
	"github.com/alv91/vault-backup/internal/app"
	"github.com/alv91/vault-backup/internal/pkg/s3"
	"github.com/alv91/vault-backup/internal/pkg/vault"
	"github.com/spf13/cobra"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup vault secrets using raft snapshot",
	Run: func(cmd *cobra.Command, args []string) {

		vaultCfg := &vault.Config{
			Token:     vaultToken,
			Address:   vaultAddr,
			Namespace: vaultNamespace,
			Timeout:   vaultTimeout,
		}

		s3Cfg := &s3.Client{
			AccessKey:       s3AccessKey,
			SecretAccessKey: s3SecretKey,
			Region:          s3Region,
			Bucket:          s3Bucket,
			Endpoint:        s3Endpoint,
			FileName:        s3FileName,
		}

		err := app.Backup(vaultCfg, s3Cfg)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
}
