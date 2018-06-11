package translate_test

import (
	"projectanima/AnimaDB/domain"
	"projectanima/AnimaTranslate/translate"
	"testing"
)

// CreateBank

// Positive Tests
func TestCreateBank(t *testing.T) {

	englishBank, errorFound :=
		translate.CreateBank(translate.LanguageBank{}, "en", "../translations")

	if errorFound != nil {
		t.Error("Bank Failed To Be Created")
	}
	if tValue, _ := englishBank.Translate("common-Pack.identifier"); tValue != "en" {
		t.Error("Translation Not Loaded")
	}
}

func TestCreateBank_RegionLoad(t *testing.T) {

	englishBank, errorFound :=
		translate.CreateBank(translate.LanguageBank{}, "en-us", "../translations")

	if errorFound != nil {
		t.Error("Bank Failed To Be Created")
	}
	if tValue, _ := englishBank.Translate("common-Pack.identifier"); tValue != "en-us" {
		t.Error("Incorrect Translation Loaded")
	}
}

/*func TestCreateBank_LanguageSingleton(t *testing.T) {

	englishBank, errorFound :=
		translate.CreateBank(translate.LanguageBank{}, "en", "../translations")

	if errorFound != nil {
		t.Error("Bank Failed To Be Created")
	}

	// Recalling The Function Should Return The Same Instance
	englishRecall, _ :=
		translate.CreateBank(translate.LanguageBank{}, "en", "../translations")

	if &englishBank != &englishRecall {
		t.Error("New Instance Of Bank Created")
	}
}*/

// Negative Tests
func TestCreateBank_(t *testing.T) {

	englishBank, errorFound :=
		translate.CreateBank(translate.LanguageBank{}, "en", "../translations")

	if errorFound != nil {
		t.Error("Bank Failed To Be Created")
	}
	if tValue, _ := englishBank.Translate("common-Pack.identifier"); tValue != "en" {
		t.Error("Translation Not Loaded")
	}
}
func TestCreateBank_NoLanguageError(t *testing.T) {

	_, errorFound :=
		translate.CreateBank(translate.LanguageBank{}, "", "../translations")

	if errorFound == nil {
		t.Error("Language Not Given Error Not Thrown")
	}
}
func TestCreateBank_NullDirectoryError(t *testing.T) {
	translate.ToggleLanguageCaching()
	_, errorFound :=
		translate.CreateBank(translate.LanguageBank{}, "en", "")

	if errorFound == nil {
		t.Error("Directory Error Not Thrown")
	}
	translate.ToggleLanguageCaching()
}
func TestCreateBank_FileGivenError(t *testing.T) {
	translate.ToggleLanguageCaching()
	_, errorFound :=
		translate.CreateBank(translate.LanguageBank{}, "en", "../translations.json")

	if errorFound == nil {
		t.Error("File Error Not Thrown")
	}
	translate.ToggleLanguageCaching()
}
func TestCreateBank_LanguageDirectoryError(t *testing.T) {
	translate.ToggleLanguageCaching()
	_, errorFound :=
		translate.CreateBank(translate.LanguageBank{}, "gg", "../translations")

	if errorFound == nil {
		t.Error("Language Error Not Thrown")
	}
	translate.ToggleLanguageCaching()
}
func TestCreateBank_RegionalDirectoryError(t *testing.T) {
	translate.ToggleLanguageCaching()
	_, errorFound :=
		translate.CreateBank(translate.LanguageBank{}, "en-pt", "../translations")

	if errorFound == nil {
		t.Error("Regional Language Error Not Thrown")
	}
	translate.ToggleLanguageCaching()
}
func TestCreateBank_LanguageBankTypeError(t *testing.T) {
	translate.ToggleLanguageCaching()
	_, errorFound :=
		translate.CreateBank(domain.Account{}, "en", "../translations")

	if errorFound == nil {
		t.Error("Language Bank Type Error Not Thrown")
	}
	translate.ToggleLanguageCaching()
}
