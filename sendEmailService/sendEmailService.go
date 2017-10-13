package sendEmailService

import (
	"os"
	"io/ioutil"
	"gomail-master"
	"crypto/tls"
)

func SendEmailService() (map[string]int)  {
	emailAndPath, err := ParseConfigFile()
	result := make(map[string]int)

	if err == nil {
		for email, path := range emailAndPath {
			files := make(map[string]int64)

			dirFiles, err := DirReader(path)
			if err != nil {
				result[email] = 2
			} else {
				for i := 0; i < len(dirFiles); i++ {
					unixTime := ModFileTime(path, dirFiles[i])
					files[dirFiles[i]] = unixTime
				}

				lastFile := FindLastFile(files)
				res := SendMail(email, path+lastFile)
				if res != nil {
					result[email] = 1
				} else {
					result[email] = 0
				}
			}
		}
	} else {
		// TODO: Config not found
	}
	return result
}

func SendMail(toMail string, pathFile string) (error) {
	m := gomail.NewMessage()
	m.SetHeader("From", "a.sveshnikova@topdelivery.ru")
	m.SetHeader("To", toMail)
	m.SetHeader("Subject", "Расчетные листки")
	m.SetBody("text/html", "Добрый день! Во вложении расчетный листок по заработной плате.")
	m.Attach(pathFile)

	d := gomail.NewDialer("smtp.yandex.ru", 25, "", "")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(m)
	if err != nil {
		return err
	} else {
		return err
	}
}

func ModFileTime(pathDir string, fileName string) (int64)  {
	info, err := os.Stat(pathDir + fileName)
	if err != nil {
		// TODO: handle errors (e.g. file not found)
	}

	unixTime := info.ModTime().Unix()
	return unixTime
}

func DirReader(pathDir string) ([]string, error) {
	var fileName = make([]string, 0)
	files, err := ioutil.ReadDir(pathDir)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if !f.IsDir() {
			fileName = append(fileName, f.Name())
		}
	}
	return fileName, nil
}

func FindLastFile(files map[string]int64) (string) {
	var lastFile string
	var unixTime int64 = 0

	for file, time := range files {
		if time > unixTime {
			unixTime = time
			lastFile = file
		}
	}
	return lastFile
}
