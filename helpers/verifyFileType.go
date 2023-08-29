package helpers

func VerfiyFileType(fileContentType string) (isValid bool) {
	contentTypeAccept := []string{"image/png", "image/jpg", "image/jpeg"}

	isValid = false
	for _, cType := range contentTypeAccept {
		if cType == fileContentType {
			isValid = true
			break
		}
	}

	return isValid
}
