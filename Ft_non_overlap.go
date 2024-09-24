package poulpe

func Ft_non_overlap(intervals [][]int) int {
	Taille := len(intervals)
	Slice1 := 0
	AncienFin := intervals[0][1]
	for i := 1; i < Taille; i++ {
		Debut := intervals[i][0]
		fin := intervals[i][1]
		if Debut < fin {
			Slice1++
			if fin <= AncienFin {
				AncienFin = fin
			} else {
				AncienFin = fin
			}
		}
	}
	return Slice1
}
