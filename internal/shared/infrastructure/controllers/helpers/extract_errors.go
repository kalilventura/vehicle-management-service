package helpers

func ExtractValidationErrors(err error) map[string]string {
	return map[string]string{
		"errors": err.Error(),
	}
}
