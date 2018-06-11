package translate_test

import (
	"os"
	"projectanima/AnimaTranslate/translate"
	"testing"
)

func checkForErrors(err error, t *testing.T) {
	if err == nil {
		return
	}
	t.Error(err.Error())
}

// Load Bank ----------
// Positive Tests
func TestLoadBank(t *testing.T) {
	englishBank := translate.LanguageBank{LanguageCode: "en"}

	errors := englishBank.LoadBank("./../translations/en/common.json")
	checkForErrors(errors, t)

	//t.Logf("%d %s", len(englishBank.TranslationBank), "Banks Loaded")
	if len(englishBank.TranslationBank) == 0 {
		t.Error("Bank Not Loaded")
	}
}

func TestLoadBank_Multiple(t *testing.T) {
	englishBank := translate.LanguageBank{LanguageCode: "en"}

	// Load First File
	errors := englishBank.LoadBank("./../translations/en/common.json")
	checkForErrors(errors, t)
	if englishBank.TranslationBank["common.yes"] == "" {
		t.Error("en.common Bank Not Loaded")
	}

	// Load Second File
	errors = englishBank.LoadBank("../translations/en/error.json")
	checkForErrors(errors, t)
	if englishBank.TranslationBank["error-Page.pagenotfound"] == "" {
		t.Error("en.error Bank Not Loaded")
	}
	if englishBank.TranslationBank["common.yes"] == "" {
		t.Error("en.common Bank Overwritten")
	}

}

// Negative Tests
func TestLoadBank_FileNotFound(t *testing.T) {
	englishBank := translate.LanguageBank{LanguageCode: "en"}

	errors := englishBank.LoadBank("./../translations/en/common2.json")

	_, validErrorType := errors.(*os.PathError)
	if validErrorType == false {
		t.Error("Incorrect File Error Not Logged Correctly")
	}

}

// LoadDirectory ----------
// Positive Tests
func TestLoadBankDirectory(t *testing.T) {
	englishBank := translate.LanguageBank{LanguageCode: "en"}

	errors := englishBank.LoadBankDirectory("./../translations/en")
	checkForErrors(errors, t)

	if englishBank.TranslationBank["common.yes"] == "" {
		t.Error("en.common Bank Not Loaded")
	}
	if englishBank.TranslationBank["error-Page.pagenotfound"] == "" {
		t.Error("en.error Bank Not Loaded")
	}
	if englishBank.TranslationBank["user-Name.firstname"] == "" {
		t.Error("en.user Bank Not Loaded")
	}
}

// Negative Tests
func TestLoadBankDirectory_DirectoryError(t *testing.T) {
	englishBank := translate.LanguageBank{LanguageCode: "en"}

	errors := englishBank.LoadBankDirectory("./../translations/en2")

	if errors == nil {
		t.Error("Directory Error... Error")
	}
}

// Translate -------------
// Positive Tests
func TestTranslate(t *testing.T) {
	englishBank := translate.LanguageBank{LanguageCode: "en"}

	errors := englishBank.LoadBankDirectory("./../translations/en")
	checkForErrors(errors, t)

	if text, _ := englishBank.Translate("common.yes"); text != "Yes" {
		t.Error("Incorrect Translation Found")
	}
}

// Negative Tests
func TestTranslate_MissingTranslation(t *testing.T) {
	englishBank := translate.LanguageBank{LanguageCode: "en"}

	errors := englishBank.LoadBankDirectory("./../translations/en")
	checkForErrors(errors, t)

	if text, _ := englishBank.Translate("common.yes33"); text != "" {
		t.Error("Incorrect Translation Found")
	}
}
