// +build !dumpkeys

package dumpkeys

import (
	ci "github.com/libp2p/go-libp2p-core/crypto"
)

func DumpKey(
		localAddr string,
		remoteAddr string,
		local ci.StretchedKeys,
		localCipherType string,
		localHMACType string,
		remote ci.StretchedKeys,
		remoteCipherType string,
		remoteHMACType string,
	) error {
	// simply does nothing - with dumpkeys.go it emulates C macro behavior
	return nil;
}
