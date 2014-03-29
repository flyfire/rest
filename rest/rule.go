package rest

type (
    Rule interface {
        Check(Value) bool
    }

    RangeRule struct {
        Name string
        Min  int
        Max  int
    }
)

/* --- rule --------------------------------------- */

func (this *RangeRule) Check(value Value) bool {
    v, ok := value[this.Name]
    if !ok {
        return false
    }

    switch o := v.(type) {
    case string:
        if n := len(o); n < this.Min || n > this.Max {
            return false
        }
    case int:
        if o < this.Min || o > this.Max {
            return false
        }
    }

    return true
}

func RuleCheck(value Value, rules ...Rule) bool {
    for _, r := range rules {
        if !r.Check(value) {
            return false
        }
    }

    return true
}
