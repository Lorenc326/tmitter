package user

func atLeastTwice(users []user) bool {
	if len(users) < 2 {
		return false
	}
	var min, max = users[0].age, users[0].age
	for i := 1; i < len(users); i++ {
		age := users[i].age
		if age > max {
			max = age
		} else if age < min {
			min = age
		} else {
			continue
		}
		if max >= min*2 {
			return true
		}
	}
	return false
}

func exactlyTwice(users []user) bool {
	if len(users) < 2 {
		return false
	}
	ages := map[int]bool{}
	for _, u := range users {
		ages[u.age] = true
		if ages[u.age*2] || (u.age%2 == 0 && ages[u.age/2]) {
			return true
		}
	}
	return false
}

func constrainedExactlyTwice(users []user) bool {
	if len(users) < 2 {
		return false
	}
	// as we have stale age range 18-80 index is age-18
	ages := make([]bool, 63)
	for _, u := range users {
		// if big sample of same age users is expected
		// consider skipping next ops for already checked
		//
		// if ages[u.age-18] {
		//	 continue
		// }
		ages[u.age-18] = true
		// we check if already exists minRangeAge * 2 or maxRangeAge / 2
		if (u.age <= 40 && ages[(u.age*2)-18]) || (u.age >= 36 && u.age%2 == 0 && ages[(u.age/2)-18]) {
			return true
		}
	}
	return false
}
