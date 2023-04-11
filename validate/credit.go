package validate

func validateCredit(credit, stake uint) bool {
	if credit == 0 {
		return false
	}
	if stake == 0 {
		return false
	}
	if credit < stake {
		return false
	}
	return true
}
