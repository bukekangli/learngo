package test

import (
	"bufio"
	"os"
	"strings"
	"testing"
	"wps.cn/lib/go/crypto"
)

func TestDecryptFromFile(t *testing.T) {
	filename := "/Users/shanglikang/go/src/github.com/bukekangli/learngo/test/file_id_220228_230227_with_name/open_android_220228.txt"
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), "\t")
		//t.Logf("%#v", items)
		if len(items) != 2 {
			t.Errorf("len(items)=%d", len(items))
			continue
		}
		data, err := crypto.AesDecrypt(EncryptKey, []byte(items[1]), crypto.Pkcs7Unpadding)
		if err != nil {
			t.Errorf(err.Error())
			continue
		}
		t.Logf("file_id: %s filename: %s", items[0], string(data))
	}
}
