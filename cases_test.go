package main

type CalcMaxCombTestCase struct {
	description string
	input       string
	expected    int
}

var testCases = []CalcMaxCombTestCase{
	{
		description: "正の数3つ",
		input:       "10 20 30",
		expected:    6000,
	},
	{
		description: "最大値の積の組み合わせに負の数が含まれるパターン",
		input:       "10 20 -30 -40",
		expected:    24000,
	},
	{
		description: "最大値の積の組み合わせに重複した正の数が含まれるパターン",
		input:       "10 20 30 30",
		expected:    18000,
	},
	{
		description: "2番目、3番目に大きい数の積より1番目、2番目に小さい数の積の方が大きいパターン",
		input:       "10 20 30 -40 -50 -60",
		expected:    90000,
	},
	{
		description: "最大値の積の組み合わせに重複した負の数が含まれるパターン",
		input:       "10 20 30 -40 -40",
		expected:    48000,
	},
	{
		description: "全て負の数のパターン",
		input:       "-10 -20 -30 -40 -50 -60",
		expected:    -6000,
	},
	{
		description: "全て負の数のパターン",
		input:       "-10 -20 -30 -40 -50 -60",
		expected:    -6000,
	},
}
