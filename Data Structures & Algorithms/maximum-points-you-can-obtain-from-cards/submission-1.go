/*
start: 7:01
idea: 2pointer
t: s:
end:
*/
func maxScore(cardPoints []int, k int) int {
	if len(cardPoints) == 0 || len(cardPoints) < k {
		return 0
	}

	totalScore := 0
	for i:=0; i<len(cardPoints); i++ {
		totalScore += cardPoints[i]
	}

	if len(cardPoints) == k {
		return totalScore
	}

	maxScore := 0
	score := 0
	start := 0
	for end := 0; end<len(cardPoints); end++ {
		score += cardPoints[end]

		if end-start+1 == len(cardPoints)-k {
			maxScore = max(maxScore, totalScore-score)
			score -= cardPoints[start]
			start += 1
		}
	}

	return maxScore
}
