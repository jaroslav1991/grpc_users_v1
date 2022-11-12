package validators

func ValidatePassword(password string) bool {
	if len(password) < 5 {
		return false
	}
	return true
}
