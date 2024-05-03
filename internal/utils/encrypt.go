package utils

func Encrypt(text *[]byte) {
	for ind, elem := range *text {
		(*text)[ind] = encrypt_func(elem)
	}
}

func encrypt_func(elem byte) byte {
	return elem - 5
}
