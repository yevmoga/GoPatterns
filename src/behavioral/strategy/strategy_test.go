package strategy

import "testing"

func TestStrategy(t *testing.T) {

	testCases := []struct {
		name           string
		strategy       Strategist
		from           string
		to             string
		expectedResult string
	}{
		{"first way", &CarStrategy{}, "Railway Station", "Central Street", "please, go form Railway Station to Central Street using car"},
		{"second way", &WalkStrategy{}, "Central Street", "Restaurant", "please, go form Central Street to Restaurant using feet"},
	}

	ctx := Context{}

	for _, testCase := range testCases {
		ctx.SetStrategy(testCase.strategy)
		way := ctx.Do(testCase.from, testCase.to)
		if way != testCase.expectedResult {
			t.Errorf("Expect result for testcase \"%s\" to equal \"%s\", but \"%s\".\n", testCase.name, way, testCase.expectedResult)
		}
	}
}
