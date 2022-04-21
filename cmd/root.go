package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfgFile        string
	vaultAddr      string
	vaultToken     string
	vaultNamespace string
	vaultTimeout   time.Duration
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
	Short: "Tool for backing vault using snapshots",
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
	rootCmd.PersistentFlags().StringVarP(&vaultAddr, "vault-address", "a", "https://127.0.0.1:8200", "vault address")
	rootCmd.PersistentFlags().StringVarP(&vaultNamespace, "vault-namespace", "n", "admin", "vault namespace")
	rootCmd.PersistentFlags().StringVarP(&vaultToken, "vault-token", "t", "", "vault token")
	rootCmd.PersistentFlags().DurationVar(&vaultTimeout, "vault-timeout", 60*time.Second, "vault client timeout")
	rootCmd.PersistentFlags().StringVar(&s3AccessKey, "s3-access-key", "", "s3 access key")
	rootCmd.PersistentFlags().StringVar(&s3SecretKey, "s3-secret-key", "", "s3 secret key")
	rootCmd.PersistentFlags().StringVar(&s3Bucket, "s3-bucket", "", "s3 bucket")
	rootCmd.PersistentFlags().StringVar(&s3Region, "s3-region", "eu-central-1", "s3 region")
	rootCmd.PersistentFlags().StringVar(&s3Endpoint, "s3-endpoint", "", "s3 endpoint")
	rootCmd.PersistentFlags().StringVar(&s3FileName, "s3-filename", "backup-latest.snap", "s3 filename to restore")

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

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	_ = viper.BindPFlags(rootCmd.PersistentFlags())
	bindFlags(rootCmd)

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// Bind each cobra flag to its associated viper configuration (config file and environment variable)
// This is required because viper doesn't work with cobra flags that are also bound to a variable
// (e.g. using StringVar to bind a flag to a string variable). See https://github.com/spf13/viper/issues/671.
func bindFlags(cmd *cobra.Command) {

	cmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && viper.IsSet(f.Name) {
			val := viper.Get(f.Name)
			cmd.PersistentFlags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})

}
