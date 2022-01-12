package leapyear

func checkLeap(year int) string {
	if year%400 == 0 {
		return "Leap year"
	} else if year%100 == 0 {
		return "Not a leap year"
	} else if year%4 == 0 {
		return "Leap year"
	}
	return "Not a leap year"

}
