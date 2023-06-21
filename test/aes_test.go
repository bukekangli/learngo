package test

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
	"wps.cn/lib/go/crypto"
)

var EncryptKey = []byte("ypkljluxulvkvhfyksnpswgqnsnjaqdj")

func TestEncryptFilename(t *testing.T) {
	type storage struct {
		InputFilePath  string
		OutputFilePath string
		Filename       string
	}
	dir := "/Users/shanglikang/Downloads/file_id_220228_230227"
	targetDir := "./file_id_220228_230227_with_name"
	if _, err := os.Stat(targetDir); err == nil {
		os.RemoveAll(targetDir)
	}
	os.Mkdir(targetDir, 0777)
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	storageList := make([]*storage, 0)
	for _, file := range files {
		storageList = append(storageList, &storage{
			InputFilePath:  strings.Join([]string{dir, file.Name()}, "/"),
			OutputFilePath: strings.Join([]string{targetDir, file.Name()}, "/"),
			Filename:       file.Name(),
		})
	}
	fileIdMap := make(map[int64]string)
	existCount := 0
	sumCount := 0
	key := []byte("ypkljluxulvkvhfyksnpswgqnsnjaqdj")
	for _, test := range storageList {
		f, err := os.Open(test.InputFilePath)
		if err != nil {
			panic(err.Error())
		}
		defer f.Close()
		outputFile, err := os.Create(test.OutputFilePath)
		if err != nil {
			panic(err)
		}
		defer outputFile.Close()
		scanner := bufio.NewScanner(f)
		fileIdList := make([]int64, 0)
		count := 0
		for scanner.Scan() {
			if count > 1000 {
				break
			}
			count++
			fileId, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				t.Errorf("parse err: %s", err.Error())
				continue
			}
			fileIdList = append(fileIdList, fileId)
			if _, exist := fileIdMap[fileId]; exist {
				existCount++
			} else {
				fileIdMap[fileId] = GetFileById(fileId)
			}
		}
		sumCount += len(fileIdList)
		for _, fileId := range fileIdList {
			var encryptFilename []byte
			if val, exist := fileIdMap[fileId]; exist {
				encryptVal, err := crypto.AesEncrypt(key, []byte(val), crypto.Pkcs7Padding)
				if err != nil {
					t.Errorf("AesEncrypt error: %s", err.Error())
				} else {
					base64Val := base64.StdEncoding.EncodeToString(encryptVal)
					encryptFilename = []byte(base64Val)
				}
			}
			//outputFile.WriteString(fmt.Sprintf("%d\t%s\n", fileId, encryptFilename))
			b := []byte(strconv.Itoa(int(fileId)))
			b = append(b, '\t')
			b = append(b, encryptFilename...)
			b = append(b, '\n')
			outputFile.Write(b)
		}
		t.Logf("filename: %s len(fileIdList)=%d ", test.InputFilePath, len(fileIdList))
	}

	f, err := os.Create("./file_id_220228_230227_with_name.zip")
	if err != nil {
		panic(err)
	}
	w := zip.NewWriter(f)
	for _, storage := range storageList {
		data, err := ioutil.ReadFile(storage.OutputFilePath)
		if err != nil {
			panic(err)
		}
		f, err := w.Create(storage.OutputFilePath)
		if err != nil {
			panic(err)
		}
		_, err = f.Write(data)
		if err != nil {
			panic(err)
		}
	}
	err = w.Close()
	if err != nil {
		panic(err)
	}

	t.Logf("existCount: %d sumCount: %d sumUniqueCount: %d", existCount, sumCount, len(fileIdMap))
}

func GetFileById(fileId int64) string {
	return fmt.Sprintf("filename_%d", fileId)
}

func TestEncrypt(t *testing.T) {
	data := []byte("asdfasdkjfajsljdflajslkjdf;lajsdjkfl;akjsd;ljkflkajsdl;jkf;lajksd;ljfad")
	data = []byte("filename_100000273922")
	encryptData, err := crypto.AesEncrypt(EncryptKey, data, crypto.Pkcs7Padding)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	//encryptData = bytes.ReplaceAll(encryptData, []byte("\n"), []byte(""))
	encryptData = []byte(strings.ReplaceAll(string(encryptData), "\r", ""))
	f.Write(encryptData)
	dataFromFile, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	originData, err := crypto.AesDecrypt(EncryptKey, dataFromFile, crypto.Pkcs7Unpadding)
	if err != nil {
		panic(err)
	}
	t.Logf("'%s' %t", string(originData), string(data) == string(originData))
}

func TestDecrypt(t *testing.T) {
	data := []byte("Ÿ\u0017íð$îò$ÎEÄÃÿ6\n|¾¢íÒ•P\f•²”+YH=Rî")
	data = []byte("g~û\"@\vëj¡\u0000š›Z¿3¼tjê•~ø\u001Aþ\u001F\"ø—¹\nþÔ")
	data = []byte("‘]|\u009Dàj«k×Y\u0017{|*\u008DÙñÄ‚=Ö²\u0080%\nÚ\u008D6Cx\u001F)")
	data = []byte("ÃKÎ¿ª6¢îëc|¬\u0000ð{Z„ LAyÏ·-HÝZž\v·›Þ")
	data = []byte("‘]|\u009Dàj«k×Y\u0017{|*\u008DÙñÄ‚=Ö²\u0080%\nÚ\u008D6Cx\u001F)")
	decryptData, err := crypto.AesDecrypt(EncryptKey, data, crypto.Pkcs7Unpadding)
	if err != nil {
		panic(err)
	}
	t.Logf(string(decryptData))
}

func TestDecryptFile(t *testing.T) {
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

func TestDecryptFileV2(t *testing.T) {
	filename := "/Users/shanglikang/go/src/github.com/bukekangli/learngo/test/open_android_220228.txt"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	dataStr := string(data)
	dataList := strings.Split(dataStr, "\n")
	for _, line := range dataList {
		items := strings.SplitAfterN(line, "\t", 2)
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

func TestDecryptFileWithSelfSplit(t *testing.T) {
	filename := "/Users/shanglikang/go/src/github.com/bukekangli/learngo/test/open_android_220228.txt"
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(ScanLines)
	for scanner.Scan() {
		items := strings.SplitAfterN(scanner.Text(), "\t", 2)
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

func TestDecryptFileWithBase64(t *testing.T) {
	filename := "/Users/shanglikang/go/src/github.com/bukekangli/learngo/test/file_id_220228_230227_with_name/open_android_220228.txt"
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(ScanLines)
	for scanner.Scan() {
		items := strings.SplitN(scanner.Text(), "\t", 2)
		//t.Logf("%#v", items)
		if len(items) != 2 {
			t.Errorf("len(items)=%d", len(items))
			continue
		}
		val, err := base64.StdEncoding.DecodeString(items[1])
		if err != nil {
			panic(err)
		}
		data, err := crypto.AesDecrypt(EncryptKey, val, crypto.Pkcs7Unpadding)
		if err != nil {
			t.Errorf(err.Error())
			continue
		}
		t.Logf("file_id: %s filename: %s", items[0], string(data))
	}
}
func runesToUTF8Manual2(rs []rune) []byte {
	size := 0
	for _, r := range rs {
		size += utf8.RuneLen(r)
	}

	bs := make([]byte, size)

	count := 0
	for _, r := range rs {
		count += utf8.EncodeRune(bs[count:], r)
	}

	return bs
}

func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}
