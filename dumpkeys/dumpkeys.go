// +build dumpkeys

package dumpkeys

import (
	"os"

	ci "github.com/libp2p/go-libp2p-core/crypto"
)

func dumpKey(localAddr string, key ci.PrivKey) {
	keylogPath := os.Getenv("LIBP2P_SECIO_KEYLOG")
	if env == "" {
		// according to the rules both the build tag and the env variable should be provided
		return;
	}

	f, err := os.OpenFile(keylogPath, os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(localAddr + ","); err != nil {
		panic(err)
	}
	if _, err = f.WriteString(s.localKey); err != nil {
		panic(err)
	}
}