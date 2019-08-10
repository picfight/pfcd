package pfcregtest

import (
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/fileops"
	"testing"
)

func TestFindDCR(t *testing.T) {
	path := fileops.Abs("../../picfight/pfcd")
	pin.D("path", path)
}
