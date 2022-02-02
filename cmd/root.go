package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile        string
	vaultAddr      string
	vaultToken     string
	vaultNamespace string
	s3AccessKey    string
	s3SecretKey    string
	s3Bucket       string
	s3Region       string
	s3Endpoint     string
	s3FileName     string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vault-backup",
	Short: "Tool for backuping vault using snapshots",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.vault-backup.yaml)")
	rootCmd.PersistentFlags().StringVarP(&vaultAddr, "vault-address", "a", "https://127.0.0.1:8200", "Vault address")
	rootCmd.PersistentFlags().StringVarP(&vaultNamespace, "vault-namespace", "n", "admin", "Vault namespace")
	rootCmd.PersistentFlags().StringVarP(&vaultToken, "vault-token", "t", "", "Token for Vault API")
	rootCmd.PersistentFlags().StringVar(&s3AccessKey, "s3-access-key", "", "s3 access key")
	rootCmd.PersistentFlags().StringVar(&s3SecretKey, "s3-secret-key", "", "s3 secret key")
	rootCmd.PersistentFlags().StringVar(&s3Bucket, "s3-bucket", "", "s3 bucket")
	rootCmd.PersistentFlags().StringVar(&s3Region, "s3-region", "eu-central-1", "s3 region")
	rootCmd.PersistentFlags().StringVar(&s3Endpoint, "s3-endpoint", "", "s3 endpoint")
	rootCmd.PersistentFlags().StringVar(&s3FileName, "s3-filename", "", "s3 filename to restore (default: latest)")

	_ = rootCmd.MarkPersistentFlagRequired("vault-address")
	_ = rootCmd.MarkPersistentFlagRequired("vault-token")
	_ = rootCmd.MarkPersistentFlagRequired("s3-access-key")
	_ = rootCmd.MarkPersistentFlagRequired("s3-secret-key")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".vault-backup" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".vault-backup")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
