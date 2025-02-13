package secio

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"testing"
)

type keyInfo struct {
	cipherKey []byte
	iv        []byte
	macKey    []byte
}

func getTestKeyInfo() *keyInfo {
	return &keyInfo{
		cipherKey: []byte("this is a test keyaaaaaaaaaaaaaa"),
		iv:        make([]byte, 16),
		macKey:    []byte("key for the mac"),
	}
}

func getTestingWriter(w io.Writer, ki *keyInfo) (*etmWriter, error) {
	c, err := aes.NewCipher(ki.cipherKey)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(c, ki.iv)

	mac, err := newMac("SHA256", ki.macKey)
	if err != nil {
		return nil, err
	}

	return NewETMWriter(w, stream, mac).(*etmWriter), nil
}

func getTestingReader(r io.Reader, ki *keyInfo) (*etmReader, error) {
	c, err := aes.NewCipher(ki.cipherKey)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewCFBDecrypter(c, ki.iv)

	mac, err := newMac("SHA256", ki.macKey)
	if err != nil {
		return nil, err
	}

	return NewETMReader(r, stream, mac).(*etmReader), nil
}

func TestBasicETMStream(t *testing.T) {
	buf := new(bytes.Buffer)

	ki := getTestKeyInfo()
	w, err := getTestingWriter(buf, ki)
	if err != nil {
		t.Fatal(err)
	}

	before := []byte("hello world")
	err = w.WriteMsg(before)
	if err != nil {
		t.Fatal(err)
	}

	r, err := getTestingReader(buf, ki)
	if err != nil {
		t.Fatal(err)
	}

	msg, err := r.ReadMsg()
	if err != nil {
		t.Fatal(err)
	}

	if string(before) != string(msg) {
		t.Fatal("got wrong message")
	}

	w.Close()
	r.Close()
}

func TestBasicETMStreamRaw(t *testing.T) {
	fmt.Println("asdasd");
	buf := new(bytes.Buffer)

	ki := getTestKeyInfo()
	w, err := getTestingWriter(buf, ki)
	if err != nil {
		t.Fatal(err)
	}

	before := []byte("hello world")
	wlen, err := w.Write(before)
	if err != nil {
		t.Fatal(err)
	}

	if wlen != len(before) {
		t.Errorf("Write length mismatch. %d != %d", wlen, len(before))
	}

	r, err := getTestingReader(buf, ki)
	if err != nil {
		t.Fatal(err)
	}

	after := make([]byte, wlen)
	rlen, err := r.Read(after)
	if err != nil {
		t.Fatal(err)
	}

	if rlen != wlen {
		t.Errorf("Read length mismatch. %d != %d", rlen, wlen)
	}

	if string(before) != string(after) {
		t.Fatal("got wrong message")
	}

	w.Close()
	r.Close()
}
