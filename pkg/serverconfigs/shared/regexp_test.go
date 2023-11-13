package shared_test

import (
	"testing"

	"github.com/iwind/TeaGo/assert"
	"github.com/oy1978/EdgeCommon/pkg/serverconfigs/shared"
)

func TestRegexp(t *testing.T) {
	var a = assert.NewAssertion(t)

	a.IsTrue(shared.RegexpFloatNumber.MatchString("123"))
	a.IsTrue(shared.RegexpFloatNumber.MatchString("123.456"))
	a.IsFalse(shared.RegexpFloatNumber.MatchString(".456"))
	a.IsFalse(shared.RegexpFloatNumber.MatchString("abc"))
	a.IsFalse(shared.RegexpFloatNumber.MatchString("123."))
	a.IsFalse(shared.RegexpFloatNumber.MatchString("123.456e7"))
	a.IsTrue(shared.RegexpNamedVariable.MatchString("${abc.efg}"))
	a.IsTrue(shared.RegexpNamedVariable.MatchString("${abc}"))
	a.IsFalse(shared.RegexpNamedVariable.MatchString("{abc.efg}"))
}
