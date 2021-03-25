package parsers

type ParsingError string

func (pe ParsingError) Error() string {
	return string(pe)
}
