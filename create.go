package bagit

import (
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// Create creates a new bagit archive
func (b *Bagit) Create(srcDir string, outDir string, hashalg string, addHeader string) error {

	var err error

	var mapheader map[string]interface{}
	if len(addHeader) != 0 {
		mapheader = getaddHeader(addHeader)
	}

	if hashalg == "md5" {
		log.Println("WARNING: md5 has known collisions. You should not use md5.")
		log.Println("WARNING: Press Ctrl + C to cancel or wait 5 seconds to continue...")
		time.Sleep(5 * time.Second)
	}

	// create bagit directory
	err = os.Mkdir(outDir, 0700)
	e(err)

	// create payload dir
	err = os.Mkdir(outDir+"/data", 0700)
	e(err)

	// create bagit.txt tag file
	fd, err := os.Create(outDir + "/bagit.txt")
	e(err)
	defer fd.Close()

	_, err = fd.WriteString("BagIt-Version: " + BagitVer + "\n")
	e(err)
	_, err = fd.WriteString("Tag-File-Character-Encoding: " + TagFileCharEnc)
	e(err)

	// create manifest-ALG.txt file
	fm, err := os.Create(outDir + "/manifest-" + hashalg + ".txt")
	e(err)
	defer fm.Close()

	// create bag-info.txt file
	fi, err := os.Create(outDir + "/bag-info.txt")
	e(err)
	defer fi.Close()

	// copy source to data dir in new bag, calculate oxum and count bytes of payload
	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			b.Oxum.Filecount++
			fsize, err := os.Stat(path)
			e(err)
			b.Oxum.Bytes += fsize.Size()
			_, err = fm.WriteString(hex.EncodeToString(hashit(path, hashalg)) + " " + path + "\n")
			copy(path, outDir+"/data/"+path)

		} else {
			os.MkdirAll(outDir+"/data/"+path, 0700)
		}
		return nil
	})
	e(err)

	// write bag-info.txt
	oxumbytes := int(b.Oxum.Bytes)
	_, err = fi.WriteString("Bag-Software-Agent: bagit <https://github.com/steffenfritz/bagit>\n")
	_, err = fi.WriteString("Bagging-Date: " + b.Timestamp + "\n")
	_, err = fi.WriteString("Payload-Oxum: " + strconv.Itoa(oxumbytes) + "." + strconv.Itoa(b.Oxum.Filecount) + "\n")

	// add additional headers to bag-info.txt
	if len(mapheader) != 0 {
		for k, v := range mapheader {
			_, err = fi.WriteString(k + ": " + v.(string) + "\n")
		}
	}

	// ToDo: Write optional fetch.txt file

	// ToDo: write optional tagmanifest files

	return err
}

// getaddHeader gets additional headers from a json file
func getaddHeader(addHeader string) map[string]interface{} {
	jsonFile, err := os.Open(addHeader)
	e(err)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return result
}
