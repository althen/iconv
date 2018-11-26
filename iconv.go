package iconv

import "io"

func convName(name string) string {
	if name == "utf-8" {
		return UTF8
	}
	if name == "gb18030" {
		return GB18030
	}
	return name
}

// ConvertString ...
func ConvertString(input string, fromEncoding string, toEncoding string) (string, error) {
	converter, err := NewConverter(convName(fromEncoding), convName(toEncoding))
	if err != nil {
		return "", err
	}
	return converter.ConvertString(input)
}

// ConvertBytes ...
func ConvertBytes(input []byte, fromEncoding string, toEncoding string) ([]byte, error) {
	converter, err := NewConverter(fromEncoding, toEncoding)
	if err != nil {
		return nil, err
	}
	return converter.ConvertBytes(input)
}

// Convert ...
func Convert(reader io.Reader, fromEncoding string, toEncoding string) (io.Reader, error) {
	converter, err := NewConverter(fromEncoding, toEncoding)
	if err != nil {
		return nil, err
	}
	return converter.Convert(reader), nil
}
