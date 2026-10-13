package validator

import (
	"strings"
	"testing"

	"go-hellogsm-score-calculator/internal/types"
)

func fives(n int) []int {
	result := make([]int, n)
	for i := range result {
		result[i] = 5
	}
	return result
}

func validRequest(graduationType types.GraduationType, liberalSystem string, artsLength int) types.MiddleSchoolAchievementReqDto {
	dto := types.MiddleSchoolAchievementReqDto{
		GraduationType:          string(graduationType),
		LiberalSystem:           liberalSystem,
		Achievement2_1:          fives(8),
		Achievement2_2:          fives(8),
		Achievement3_1:          fives(8),
		ArtsPhysicalAchievement: fives(artsLength),
		AbsentDays:              []int{0, 0, 0},
		AttendanceDays:          []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		VolunteerTime:           []int{0, 0, 0},
	}
	if liberalSystem == "자유학기제" {
		dto.FreeSemester = "1-1"
	}
	return dto
}

func TestValidateRequest_ArtsPhysicalAchievementLength(t *testing.T) {
	tests := []struct {
		name           string
		graduationType types.GraduationType
		liberalSystem  string
		artsLength     int
		wantErr        bool
	}{
		{"CANDIDATE 자유학년제 9개", types.CANDIDATE, "자유학년제", 9, false},
		{"CANDIDATE 자유학기제 15개", types.CANDIDATE, "자유학기제", 15, false},
		{"GRADUATE 자유학년제 12개", types.GRADUATE, "자유학년제", 12, false},
		{"GRADUATE 자유학기제 18개", types.GRADUATE, "자유학기제", 18, false},
		{"GRADUATE 자유학년제 11개 (부족)", types.GRADUATE, "자유학년제", 11, true},
		{"CANDIDATE 자유학년제 10개 (초과)", types.CANDIDATE, "자유학년제", 10, true},
		{"CANDIDATE 자유학기제 9개 (자유학년제 길이)", types.CANDIDATE, "자유학기제", 9, true},
		{"GRADUATE 자유학기제 12개 (자유학년제 길이)", types.GRADUATE, "자유학기제", 12, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateRequest(validRequest(tt.graduationType, tt.liberalSystem, tt.artsLength))
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				if !strings.Contains(err.Error(), "artsPhysicalAchievement must contain exactly") {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
		})
	}
}

func TestValidateRequest_GEDSkipsArtsPhysicalAchievementLength(t *testing.T) {
	score := 80.0
	dto := types.MiddleSchoolAchievementReqDto{GraduationType: string(types.GED), GedAvgScore: &score}

	if err := ValidateRequest(dto); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
