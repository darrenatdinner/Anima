package Procedures

// TranslationUpsert - Procedure To Update A Translation In The Database If It Exists
//	Add It If It Doesn't
var TranslationUpsert = `{
    find_translation( func: eq(key, %s AND eq(language, %s)) {
		uid
    }
}`

// LoadLanguage - Load the Complete Translation List From A Specific Language
var LoadLanguage = `
	get_translations( func: eq(language, %s)){
		key
		text
		language
	}`
