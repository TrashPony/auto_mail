package sendEmailService

import "os"

func Delete(email string) (bool) {
	emailAndPath, err := ParseConfigFile()
	if err != nil{
		panic(err)
	}

	var writer bool = false
	for existingEmail,_:= range emailAndPath {
		if existingEmail == email {
			delete(emailAndPath, existingEmail)
			writer = true
		}
	}

	if writer {
		f, err := os.Create("./config")
		for existingEmail, existingPath := range emailAndPath {
			if _, err = f.WriteString("\n" + existingEmail + ";" + existingPath); err != nil {
				panic(err)
				return false
			}
		}
		return writer
	}
	return writer
}
