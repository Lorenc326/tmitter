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
	temp := map[int]bool{}
	for _, u := range users {
		temp[u.age] = true
	}
	for age, _ := range temp {
		if _, exist := temp[age*2]; exist {
			return true
		}
	}
	return false
}

func constrainedExactlyTwice(users []user) bool {
	if len(users) < 2 {
		return false
	}
	// as we have stale range 18-80 index is age-18
	ages := make([]bool, 63)
	for _, u := range users {
		ages[u.age-18] = true
	}
	// makes sense to iterate until maxAge/2
	// while counting - convert age back
	for age, _ := range ages[:23] {
		if ages[(age+18)*2] {
			return true
		}
	}
	return false
}
