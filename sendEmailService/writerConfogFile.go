package sendEmailService

import "os"

func Write(email string, path string) (bool) {
	emailAndPath, err := ParseConfigFile()
	if err != nil{
		panic(err)
	}
	f, err := os.OpenFile("./config", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var writer bool = true

	for existingEmail, existingPath := range emailAndPath {
		if existingEmail == email {
			if existingPath == path {
				writer = false
				break
			} else {
				emailAndPath[existingEmail] = path
				writer = true
				break
			}
		}
	}

	if writer {
		f, err := os.Create("./config")
		emailAndPath[email] = path
		for existingEmail, existingPath := range emailAndPath {
			a := existingPath[len(existingPath)-1]
			if a == '\\' { // проверка на / в конце пути дериктории, если нет досталвляем
				if _, err = f.WriteString("\n" + existingEmail + ";" + existingPath); err != nil {
					panic(err)
					return false
				}
			} else {
				if _, err = f.WriteString("\n" + existingEmail + ";" + existingPath + "\\"); err != nil {
					panic(err)
					return false
				}
			}
		}
		return true
	}
	return false
}
