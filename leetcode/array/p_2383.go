package array

func MinNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	return minNumberOfHours(initialEnergy, initialExperience, energy, experience)
}

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	costEnergy := 1
	costExperience := 0
	for i := len(energy) - 1; i >= 0; i-- {
		costEnergy += energy[i]
		costExperience = max(experience[i]+1, costExperience-experience[i])
	}
	return max(costExperience-initialExperience, 0) + max(costEnergy-initialEnergy, 0)
}
