package main

func splitSuitsAndNumbers(chars string) ([]string, []string) {
	s := []string{"s", "h", "d", "c"}
	suits := []string{}
	nums := []string{}
	isSuit := false
	for _, char := range chars {
		c := string(char)
		if containsString(s, c) {
			suits = append(suits, c)
			isSuit = true
		} else {
			if isSuit {
				nums = append(nums, c)
			} else {
				nums[len(nums)-1] = nums[len(nums)-1] + c
			}
			isSuit = false
		}
	}
	return suits, nums
}

func containsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
