// +build dumpkeys

package dumpkeys

import (
	"fmt"
	"os"

	ci "github.com/libp2p/go-libp2p-core/crypto"
	b64 "encoding/base64"
)

func DumpKey(localAddr string, remoteAddr string, privKey ci.PrivKey) error {
	fmt.Println("dumping keys...")
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

	raw_key, err := privKey.Raw();
	if err != nil {
		return err;
	}

	key := b64.StdEncoding.EncodeToString([]byte(raw_key))

	if _, err = f.WriteString(localAddr + "," + remoteAddr + "," + key + "\n"); err != nil {
		return err;
	}

	return nil;
}
