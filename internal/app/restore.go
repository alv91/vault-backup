package app

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/alv91/vault-backup/internal/pkg/s3"
	"github.com/alv91/vault-backup/internal/pkg/vault"
)

func Restore(vConfig *vault.Config, s3Config *s3.Client) {
	fmt.Println("Starting restore...")

	// create s3 client
	s3Client := s3.NewClient(s3Config.AccessKey, s3Config.SecretAccessKey, s3Config.Region, s3Config.Bucket, s3Config.Endpoint, s3Config.FileName)

	// get backup from s3Config
	body := s3Client.GetObject(s3Config.FileName).Body
	data, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err)

		return
	}

	// create new buffer writer
	buf := bytes.NewBuffer(data)
	r := bufio.NewReader(buf)

	// create vault client
	vaultClient, err := vault.NewClient(vConfig)
	if err != nil {
		fmt.Println(err)

		return
	}

	// restore vault backup
	err = vaultClient.Restore(r)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Printf("Restored backup with name '%s'.", s3Config.FileName)
}
