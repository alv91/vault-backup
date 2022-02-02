package vault

import (
	vault "github.com/hashicorp/vault/api"
)

const DEFAULT_VAULT_NAMESPACE = "admin"

type Client struct {
	vaultClient  *vault.Client
	forceRestore bool
}

type Config struct {
	Address      string
	Token        string
	Namespace    string
	ForceRestore bool
	TmpPath      string
	FileName     string
}

func NewClient(config *Config) (*Client, error) {
	vaultConfig := vault.DefaultConfig()

	vaultConfig.Address = config.Address

	client, err := vault.NewClient(vaultConfig)
	if err != nil {
		return nil, err
	}

	if config.Token != "" {
		client.SetToken(config.Token)
	}

	if config.Namespace == "" {
		config.Namespace = DEFAULT_VAULT_NAMESPACE
	}

	client.SetNamespace(config.Namespace)

	return &Client{
		vaultClient:  client,
		forceRestore: config.ForceRestore,
	}, nil
}
