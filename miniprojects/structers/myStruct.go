package main

type myStruct struct {
	On    bool
	Ammo  int
	Power int
}

func (val *myStruct) Shoot() bool {

	if !val.On {
		return false
	}

	if val.Ammo > 0 {
		val.Ammo--
		return true
	} else {
		return false
	}
}

func (val *myStruct) RideBike() bool {
	if !val.On {
		return false
	}

	if val.Power > 0 {
		val.Power--
		return true
	} else {
		return false
	}
}