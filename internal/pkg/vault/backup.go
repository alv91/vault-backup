package vault

import (
	"io"
)

func (v *Client) Backup(w io.Writer) error {

	err := v.vaultClient.Sys().RaftSnapshot(w)
	if err != nil {
		return err
	}

	return nil
}
