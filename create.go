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
func (b *Bagit) Create(verbose bool) error {

	var err error

	var mapheader map[string]interface{}
	if len(*b.AddHeader) != 0 {
		mapheader = getaddHeader(*b.AddHeader)
	}

	if *b.HashAlg == "md5" {
		log.Println("WARNING: md5 has known collisions. You should not use md5.")
		log.Println("WARNING: Press Ctrl + C to cancel or wait 5 seconds to continue...")
		time.Sleep(5 * time.Second)
	}

	// create bagit directory
	err = os.Mkdir(*b.OutDir, 0700)
	e(err)

	// create payload dir
	err = os.Mkdir(*b.OutDir+"/data", 0700)
	e(err)

	// create bagit.txt tag file
	fd, err := os.Create(*b.OutDir + "/bagit.txt")
	e(err)
	defer fd.Close()

	_, err = fd.WriteString("BagIt-Version: " + BagitVer + "\n")
	e(err)
	_, err = fd.WriteString("Tag-File-Character-Encoding: " + TagFileCharEnc)
	e(err)

	// create manifest-ALG.txt file
	fm, err := os.Create(*b.OutDir + "/manifest-" + *b.HashAlg + ".txt")
	e(err)
	defer fm.Close()

	// create bag-info.txt file
	fi, err := os.Create(*b.OutDir + "/bag-info.txt")
	e(err)
	defer fi.Close()

	// Add provided fetch manifest for remote resources to manifest file
	if len(*b.FetchFile) != 0 {
	}

	// copy source to data dir in new bag, calculate oxum and count bytes of payload
	err = filepath.Walk(*b.SrcDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			b.Oxum.Filecount++
			fsize, err := os.Stat(path)
			e(err)
			b.Oxum.Bytes += fsize.Size()
			_, err = fm.WriteString(hex.EncodeToString(hashit(path, *b.HashAlg)) + " " + path + "\n")
			copy(path, *b.OutDir+"/data/"+path)

		} else {
			os.MkdirAll(*b.OutDir+"/data/"+path, 0700)
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
