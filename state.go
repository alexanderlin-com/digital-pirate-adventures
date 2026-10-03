package main

import "math/rand"

type ShipKind int

const (
	Sloop ShipKind = iota
	Frigate
	Galleon
)

type Ship struct {
	Kind ShipKind

	Health       int
	MaxHealth    int
	Bandwidth    int
	MaxBandwidth int
	Speed        int
	Dodge        int
	Armor        float64
	Accuracy     int
	Damage       float64

	BroadsideType int
	SwivelType    int
	CrewType      int
	SailType      int
	RiggingType   int

	HasSwivel bool
	HasCrew   bool
}

// NewShip's Health/Bandwidth double as MaxHealth/MaxBandwidth at creation,
// used later to render accurate health/bandwidth bars rather than against
// some unrelated fixed scale.
func NewShip(kind ShipKind) Ship {
	switch kind {
	case Frigate:
		return Ship{Kind: Frigate, Health: 350, MaxHealth: 350, Bandwidth: 250, MaxBandwidth: 250, Speed: 7, Dodge: 3, Armor: 0.20, Accuracy: 3, Damage: 1, HasSwivel: true, HasCrew: false}
	case Galleon:
		return Ship{Kind: Galleon, Health: 500, MaxHealth: 500, Bandwidth: 200, MaxBandwidth: 200, Speed: 4, Dodge: 1, Armor: 0.30, Accuracy: 6, Damage: 1.25, HasSwivel: true, HasCrew: true}
	default:
		return Ship{Kind: Sloop, Health: 200, MaxHealth: 200, Bandwidth: 100, MaxBandwidth: 100, Speed: 10, Dodge: 7, Armor: 0.1, Accuracy: 0, Damage: 0.75, HasSwivel: false, HasCrew: false}
	}
}

type Enemy struct {
	Health    int
	MaxHealth int
	Speed     int
	Accuracy  int
	Dodge     int
}

func NewEnemy() Enemy {
	return Enemy{Health: 100, MaxHealth: 100, Speed: 6, Accuracy: 5, Dodge: 4}
}

type GameState struct {
	PlayerName string
	ShipName   string

	Player Ship
	Enemy  Enemy

	VisitedLibrary bool
	HasTextbook    bool
}

func NewGameState() *GameState {
	return &GameState{Enemy: NewEnemy()}
}

func d20() int {
	return rand.Intn(20) + 1
}

// resolveHitRoll is the to-hit formula from the Java version's Fight.resolveHit,
// with the die roll passed in rather than rolled internally, so it's deterministically
// testable instead of being entangled with System.out.println the way the Java version was.
func resolveHitRoll(roll, attackerAccuracy, weaponMod, defenderDodge int) bool {
	return roll+attackerAccuracy+weaponMod > 10+defenderDodge
}

func resolveHit(attackerAccuracy, weaponMod, defenderDodge int) bool {
	return resolveHitRoll(d20(), attackerAccuracy, weaponMod, defenderDodge)
}

// mitigate mirrors the Java version's Storm.lightning case-3 pattern: armor reduces
// incoming damage proportionally.
func mitigate(rawDamage int, defenderArmor float64) int {
	return int(float64(rawDamage) * (1 - defenderArmor))
}

// initiative mirrors Fight.java's initiative(): a speed + d20 contest, used both
// for turn order and for flee checks in the fight screen.
func initiative(playerSpeed, enemySpeed int) bool {
	return playerSpeed+d20() >= enemySpeed+d20()
}
