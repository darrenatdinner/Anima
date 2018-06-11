package translate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// LanguageBank - Instance Of Translations In A Specific Language
type LanguageBank struct {
	LanguageCode    string
	TranslationBank map[string]string
}

// LoadBank - Load A Single LanguageBank From File
func (bank *LanguageBank) LoadBank(translationFile string) error {

	// Open The Given File
	jsonFile, errorFound := os.Open(translationFile)
	if errorFound != nil { // Clean And Error
		return errorFound
	}
	// Defer The Close Until After The Function
	defer jsonFile.Close()

	// Check Existing Data
	if bank.TranslationBank == nil {
		bank.TranslationBank = make(map[string]string)
	}
	// Read All Of The File Into Memory
	byteValue, errorFound := ioutil.ReadAll(jsonFile)

	if errorFound != nil { // Clean And Error
		return errorFound
	}
	// Grab The Data From The Loaded File
	var newTranslations = make(map[string]string)
	errorFound = json.Unmarshal(byteValue, &newTranslations)

	// Merge Into Existing Map
	for singleKey, singleValue := range newTranslations {
		bank.TranslationBank[singleKey] = singleValue
	}

	return errorFound
}

// LoadBankDirectory - Load All Of The Language Banks In The Given Directory
func (bank *LanguageBank) LoadBankDirectory(directoryPath string) error {

	dirFiles, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		return fmt.Errorf("error loading files from directory - %s", directoryPath)
	}
	// Swap For Threaded Loading````````````````````````````````````````````````````````
	for _, fileInfo := range dirFiles {
		// Skip Directories
		if fileInfo.IsDir() {
			continue
		}
		// Load File Into Memory
		err := bank.LoadBank(directoryPath + "/" + fileInfo.Name())
		if err != nil {
			return err
		}
	}
	return nil
}

// Translate - Retrieve A Translation In This Language
func (bank LanguageBank) Translate(translationKey string) (string, error) {
	if translation, found := bank.TranslationBank[translationKey]; found {
		return translation, nil
	}

	return "", nil
}
