package main

import "testing"

func calculateBMI(weight float64, height float64) float64 {
	return weight / (height * height)
}

func TestCalculateBMI(t *testing.T) {
	// 正常なケース
	bmi := calculateBMI(70, 1.75)
	expected := 22.86
	if bmi < expected-0.01 || bmi > expected+0.01 { // 浮動小数点の比較は誤差を考慮
		t.Errorf("BMI計算が期待通りではありません。期待値: %f, 実際: %f", expected, bmi)
	}

	// エッジケースのテスト
	bmi = calculateBMI(0, 1.75)
	if bmi != 0 {
		t.Errorf("体重が0の場合、BMIは0であるべきです。")
	}
}
