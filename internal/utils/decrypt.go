package utils

func Decrypt(text *[]byte) {
	for ind, elem := range *text {
		(*text)[ind] = decrypt_func(elem)
	}
}

func decrypt_func(elem byte) byte {
	return elem - 5
}
