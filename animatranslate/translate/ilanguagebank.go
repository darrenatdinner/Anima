package translate

import (
	"fmt"
	"strings"
)

// ILanguageBank - Specifies Data And Methods To Load And Use a Language Bank
type ILanguageBank interface {
	LoadBank(string) error
	LoadBankDirectory(string) error
	Translate(string) (string, error)
}

// Languages Are Held In Memory Statically
var loadedBanks = make(map[string]ILanguageBank)

// Use The Language Cache
var languageCacheEnabled = true

// ToggleLanguageCaching - Toggle On Or Off The Caching Of A Language *To Refresh A Language
//	Toggle, Call CreateBank Again, Then Toggle Cacheing Back On If Desired
func ToggleLanguageCaching() {
	languageCacheEnabled = !languageCacheEnabled
}

// CreateBank - Creation Function For Creating A Language Bank
func CreateBank(languageInterface interface{}, language string, directory string) (ILanguageBank, error) {
	// Validate Basic Information Given
	var languageDirectory string
	var regionDirectory string

	// If The Bank Already Exists Give The Pre-Existing
	if foundBank, ok := loadedBanks[language]; ok && languageCacheEnabled {
		return foundBank, nil
	}
	// Language Missing
	if language == "" {
		return nil, fmt.Errorf("no language given")
	}
	// Else, Create A New One To Load
	if directory == "" {
		return nil, fmt.Errorf("no directory structure found")
	}
	languageStructure := strings.Split(directory, "/")
	if strings.Contains(languageStructure[len(languageStructure)-1], ".") {
		return nil, fmt.Errorf("file name found, expecting directory")
	}
	languageChoice := strings.Split(language, "-")
	languageDirectory = directory + "/" + languageChoice[0]
	// Support Regional Directories
	if len(languageChoice) > 1 {
		regionDirectory = directory + "/" + strings.Join(languageChoice, "/")
	}

	switch languageInterface.(type) {
	case LanguageBank:
		{
			// Handle Locale Information Driven By Directory input
			langBank := new(LanguageBank)
			langBank.LanguageCode = language
			if loadError := langBank.LoadBankDirectory(languageDirectory); loadError != nil {
				return nil, loadError
			}
			if regionDirectory != "" {
				if loadError := langBank.LoadBankDirectory(regionDirectory); loadError != nil {
					loadedBanks[language] = langBank
					return langBank, fmt.Errorf("regional directory failed to load")
				}
			}
			loadedBanks[language] = langBank
			return langBank, nil
		}
	}
	return nil, fmt.Errorf("unable to create a concrete ILanguageBank")
}
