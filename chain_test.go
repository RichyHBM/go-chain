package main

import (
	"errors"
	"testing"
)

func exampleThenPass() (bool, error) {
	return true, nil
}

func exampleThenPass1(i int) (bool, error) {
	return i == 1, nil
}

func exampleThenPass2(i int, j int) (bool, error) {
	return i == j, nil
}

func exampleThenPassB(b bool) (int, error) {
	if b {
		return 1, nil
	} else {
		return 0, nil
	}
}

func exampleThenPassC(i int) (bool, error) {
	if i == 1 {
		return true, nil
	} else if i == 0 {
		return false, nil
	} else {
		return false, errors.New("Bad data")
	}
}

func TestPassChained(t *testing.T) {
	Value, Error := Result(Then(Then(Run(
		exampleThenPass),
		exampleThenPassB),
		exampleThenPassC))

	if Error != nil {
		t.Fatalf(`TestPass returned error: %v`, Error)
	} else if Value != true {
		t.Fatalf(`TestPass returned false`)
	}
}

func TestPass(t *testing.T) {
	Chain := Run(exampleThenPass)
	Chain2 := Then(Chain, exampleThenPassB)
	Chain3 := Then(Chain2, exampleThenPassC)
	Value, Error := Result(Chain3)

	if Error != nil {
		t.Fatalf(`TestPass returned error: %v`, Error)
	} else if Value != true {
		t.Fatalf(`TestPass returned false`)
	}
}

func TestPass1(t *testing.T) {
	Chain := Run1(exampleThenPass1, 1)
	Chain2 := Then(Chain, exampleThenPassB)
	Chain3 := Then(Chain2, exampleThenPassC)
	Value, Error := Result(Chain3)

	if Error != nil {
		t.Fatalf(`TestPass returned error: %v`, Error)
	} else if Value != true {
		t.Fatalf(`TestPass returned false`)
	}
}

func TestPass2(t *testing.T) {
	Chain := Run2(exampleThenPass2, 1, 1)
	Chain2 := Then(Chain, exampleThenPassB)
	Chain3 := Then(Chain2, exampleThenPassC)
	Value, Error := Result(Chain3)

	if Error != nil {
		t.Fatalf(`TestPass returned error: %v`, Error)
	} else if Value != true {
		t.Fatalf(`TestPass returned false`)
	}
}

func TestFail(t *testing.T) {
	Chain := Run1(exampleThenPass1, 1)
	Chain2 := Then(Chain, exampleThenPassB)
	Chain2.Value = -1
	Chain3 := Then(Chain2, exampleThenPassC)
	_, Error := Result(Chain3)

	if Error == nil {
		t.Fatalf(`TestPass returned error: %v`, Error)
	}
}

func TestManual(t *testing.T) {
	Value, Error := exampleThenPass1(1)
	Chain := &Chain[bool] {
		Value: Value,
		Error: Error,
	}

	Chain2 := Then(Chain, exampleThenPassB)
	Chain3 := Then(Chain2, exampleThenPassC)
	Result, Error := Result(Chain3)
	
	if Error != nil {
		t.Fatalf(`TestPass returned error: %v`, Error)
	} else if Result != true {
		t.Fatalf(`TestPass returned false`)
	}
}
