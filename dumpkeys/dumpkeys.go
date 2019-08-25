// +build dumpkeys

package dumpkeys

import (
	"fmt"
	"os"

	ci "github.com/libp2p/go-libp2p-core/crypto"
	b64 "encoding/base64"
)

func DumpKey(
		localAddr string,
		remoteAddr string,
		local ci.StretchedKeys,
		localCT string,
		remote ci.StretchedKeys,
		remoteCT string,
	) error {

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

	localKey := b64.StdEncoding.EncodeToString(local.CipherKey)
	localIV := b64.StdEncoding.EncodeToString(local.IV)
	localMac := b64.StdEncoding.EncodeToString(local.MacKey)

	remoteKey := b64.StdEncoding.EncodeToString(remote.CipherKey)
	remoteIV := b64.StdEncoding.EncodeToString(remote.IV)
	remoteMac := b64.StdEncoding.EncodeToString(remote.MacKey)

	if _, err = f.WriteString(fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s\n",
		localAddr,
		remoteAddr,
		localKey,
		localIV,
		localMac,
		localCT,
		remoteKey,
		remoteIV,
		remoteMac,
		remoteCT,
	)); err != nil {
		return err;
	}

	return nil;
}
