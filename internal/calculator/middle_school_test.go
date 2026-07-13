/package calculator

import (
	"go-hellogsm-score-calculator/internal/types"
	"testing"
)

func fullMarks(n int) []int {
	achievements := make([]int, n)
	for i := range achievements {
		achievements[i] = 5
	}
	return achievements
}

func TestCalcGeneralSubjectsSemesterScore_Candidate(t *testing.T) {
	dto := types.MiddleSchoolAchievementCalcDto{
		Achievement1_2: fullMarks(5),
		Achievement2_1: fullMarks(5),
		Achievement2_2: fullMarks(5),
		Achievement3_1: fullMarks(5),
	}

	result := CalcGeneralSubjectsSemesterScore(dto, types.CANDIDATE)

	cases := map[string]struct {
		got  float64
		want float64
	}{
		"score1_2": {RatToFloat64(result.Score1_2), 18.000},
		"score2_1": {RatToFloat64(result.Score2_1), 45.000},
		"score2_2": {RatToFloat64(result.Score2_2), 45.000},
		"score3_1": {RatToFloat64(result.Score3_1), 72.000},
		"score3_2": {RatToFloat64(result.Score3_2), 0.000},
	}

	for name, c := range cases {
		if c.got != c.want {
			t.Errorf("%s = %v, want %v", name, c.got, c.want)
		}
	}

	total := RatToFloat64(CalcGeneralSubjectsTotalScore(result))
	if total != 180.000 {
		t.Errorf("total = %v, want 180.000", total)
	}
}

func TestCalcGeneralSubjectsSemesterScore_Graduate(t *testing.T) {
	dto := types.MiddleSchoolAchievementCalcDto{
		Achievement1_2: fullMarks(5),
		Achievement2_1: fullMarks(5),
		Achievement2_2: fullMarks(5),
		Achievement3_1: fullMarks(5),
		Achievement3_2: fullMarks(5),
	}

	result := CalcGeneralSubjectsSemesterScore(dto, types.GRADUATE)

	cases := map[string]struct {
		got  float64
		want float64
	}{
		"score1_2": {RatToFloat64(result.Score1_2), 18.000},
		"score2_1": {RatToFloat64(result.Score2_1), 36.000},
		"score2_2": {RatToFloat64(result.Score2_2), 36.000},
		"score3_1": {RatToFloat64(result.Score3_1), 45.000},
		"score3_2": {RatToFloat64(result.Score3_2), 45.000},
	}

	for name, c := range cases {
		if c.got != c.want {
			t.Errorf("%s = %v, want %v", name, c.got, c.want)
		}
	}

	total := RatToFloat64(CalcGeneralSubjectsTotalScore(result))
	if total != 180.000 {
		t.Errorf("total = %v, want 180.000", total)
	}
}

func TestCalcGeneralSubjectsSemesterScore_GraduateMissingFirstGrade(t *testing.T) {
	dto := types.MiddleSchoolAchievementCalcDto{
		Achievement1_2: nil,
		Achievement2_1: fullMarks(5),
		Achievement2_2: fullMarks(5),
		Achievement3_1: fullMarks(5),
		Achievement3_2: fullMarks(5),
	}

	result := CalcGeneralSubjectsSemesterScore(dto, types.GRADUATE)

	if got := RatToFloat64(result.Score1_2); got != 0.000 {
		t.Errorf("score1_2 = %v, want 0.000 when achievement1_2 is nil", got)
	}
}