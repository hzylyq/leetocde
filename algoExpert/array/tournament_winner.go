package array

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	var currentBestTeam string
	points := make(map[string]int, 0)
	for idx, competition := range competitions {
		result := results[idx]
		homeTeam, awayTeam := competition[0], competition[1]
		winner := awayTeam
		if result == HOME_TEAM_WON {
			winner = homeTeam
		}
		points[winner] += 3
		if points[winner] > points[currentBestTeam] {
			currentBestTeam = winner
		}

	}

	return currentBestTeam
}
