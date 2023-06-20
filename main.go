package main

import "fmt"

type SaftyPlacer interface {
	placeSafeties()
}

// ICE
// sandy rocks
// concrete

type RockClimber struct {
	rocksClimbed int
	kind         int
	sp           SaftyPlacer
}

func newRockClimber(sp SaftyPlacer) *RockClimber {
	return &RockClimber{
		sp: sp,
	}
}

type IceSafetyPlacer struct {
	// db
	// data
	// api
}

func (sp IceSafetyPlacer) placeSafeties() {
	// switch sp.kind {
	// case 1:
	// 	// ICE
	// case 2:
	// 	// sand
	// case 3:
	// 	// concrete
	// }
	fmt.Println("placing my ICE safeties...")
}

type NOPSafetyPlacer struct{}

func (sp NOPSafetyPlacer) placeSafeties() {
	fmt.Println("placing NO safeties...")
}

func (rc *RockClimber) climbRock() {
	rc.rocksClimbed++
	if rc.rocksClimbed == 10 {
		rc.sp.placeSafeties()
	}
}

func main() {
	// rc := &RockClimber{}
	// rc := newRockClimber(IceSafetyPlacer{})
	rc := newRockClimber(NOPSafetyPlacer{})
	for i := 0; i < 11; i++ {
		rc.climbRock()
	}
}
