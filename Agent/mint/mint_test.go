package mint

import "testing"

func TestMint(t *testing.T) {
	BatchNewColdBootClient("c_conf", 0, 999999999999999)
}

func TestWritefile(t *testing.T) {
	writeFile("sss", "./svg/s.svg")
}
