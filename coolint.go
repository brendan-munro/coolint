package coolint

type CoolInt int

func (_ CoolInt) String() string {
	return "It's cool."
}
