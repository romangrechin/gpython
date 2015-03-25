// Bool objects

package py

type Bool bool

var (
	BoolType = NewType("bool", "bool(x) -> bool\n\nReturns True when the argument x is true, False otherwise.\nThe builtins True and False are the only two instances of the class bool.\nThe class bool is a subclass of the class int, and cannot be subclassed.")
	// Some well known bools
	False = Bool(false)
	True  = Bool(true)
)

// Type of this object
func (s Bool) Type() *Type {
	return BoolType
}

// Make a new bool - returns the canonical True and False values
func NewBool(t bool) Bool {
	if t {
		return True
	}
	return False
}

func (a Bool) M__bool__() Object {
	return a
}

func (a Bool) M__index__() Int {
	if a {
		return Int(1)
	}
	return Int(0)
}

func (a Bool) M__str__() Object {
	if a {
		return String("True")
	}
	return String("False")
}

// Check interface is satisfied
var _ I__bool__ = Bool(false)
var _ I__index__ = Bool(false)
var _ I__str__ = Bool(false)

// Convert an Object to an Bool
//
// Retrurns ok as to whether the conversion worked or not
func convertToBool(other Object) (Bool, bool) {
	switch b := other.(type) {
	case Bool:
		return b, true
	case Int:
		switch b {
		case 0:
			return False, true
		case 1:
			return True, true
		default:
			return False, false
		}
	case Float:
		switch b {
		case 0:
			return False, true
		case 1:
			return True, true
		default:
			return False, false
		}
	}
	return False, false
}

func (a Bool) M__eq__(other Object) Object {
	if b, ok := convertToBool(other); ok {
		return NewBool(a == b)
	}
	return False
}

func (a Bool) M__ne__(other Object) Object {
	if b, ok := convertToBool(other); ok {
		return NewBool(a != b)
	}
	return True
}
