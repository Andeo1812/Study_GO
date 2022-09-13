package options

// FlagsData
const (
	showCountStrUsage = "Подсчитать количество повторов строки во входных данных. Вывести это число перед строкой отделив пробелом."
	showCountStrFlag  = "c"

	showUnUniqStrUsage = "Вывести только те строки, которые повторились во входных данных."
	showUnUniqStrFlag  = "d"

	showUniqStrUsage = "Вывести только те строки, которые не повторились во входных данных."
	showUniqStrFlag  = "u"

	registerNotImportantUsage = "Не учитывать регистр букв."
	registerNotImportantFlag  = "i"

	skipCountWordsUsage = "Не учитывать первые num_fields полей в строке."
	//  Полем в строке является непустой набор символов отделённый пробелом.
	skipCountWordsFlag = "f"

	skipCountCharsUsage = "Не учитывать первые num_chars символов в строке."
	//  При использовании вместе с параметром -f учитываются первые символы после num_fields
	//  полей (не учитывая пробел-разделитель после последнего поля).
	skipCountCharsFlag = "s"
)
