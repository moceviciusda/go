package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	if areEqual(l1, l2) {
		return RelationEqual
	}

	if contains(l1, l2) {
		return RelationSuperlist
	}

	if contains(l2, l1) {
		return RelationSublist
	}

	return RelationUnequal
}

func areEqual(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i, num := range l1 {
		if num != l2[i] {
			return false
		}
	}

	return true
}

func contains(l1, l2 []int) bool {
	if len(l1) < len(l2) {
		return false
	}

	// if len(l2) == 0 {
	// 	return true
	// }

	for i := range l1 {
		if len(l1[i:]) < len(l2) {
			return false
		}

		if areEqual(l1[i:i+len(l2)], l2) {
			return true
		}
		// if num1 == l2[0] {
		// 	match := true
		// 	for j, num2 := range l2 {
		// 		if num2 != l1[i+j] {
		// 			match = false
		// 			break
		// 		}
		// 	}

		// 	if match {
		// 		return true
		// 	}
		// }
	}

	// return true
	return false
}
