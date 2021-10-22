package user

func (t UserType) Equal(target UserType) bool {
	return t == target
}

func (t UserType) IsIn(targets ...UserType) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}