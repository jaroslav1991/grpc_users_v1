package posts

func EmptyTitle(title string) bool {
	if len(title) > 0 {
		return true
	}
	return false
}

func EmptyMessage(message string) bool {
	if len(message) > 0 {
		return true
	}
	return false
}
