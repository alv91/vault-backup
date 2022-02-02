package vault

import (
	"io"
)

func (v *Client) Restore(f io.Reader) error {

	err := v.vaultClient.Sys().RaftSnapshotRestore(f, v.forceRestore)
	if err != nil {
		return err
	}

	return nil
}
