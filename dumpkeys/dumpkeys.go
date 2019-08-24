// +build dumpkeys

package dumpkeys

import (
	"os"

	ci "github.com/libp2p/go-libp2p-core/crypto"
)

func dumpKey(localAddr string, key ci.PrivKey) error {
	keylogPath := os.Getenv("LIBP2P_SECIO_KEYLOG")
	if keylogPath == "" {
		// according to the rules both the build tag and the env variable should be provided
		return nil;
	}

	f, err := os.OpenFile(keylogPath, os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0600)
	if err != nil {
		return err;
	}

	defer f.Close()

	if _, err = f.WriteString(localAddr + ","); err != nil {
		return err;
	}
	if _, err = f.WriteString(key); err != nil {
		return err;
	}

	return nil;
}
