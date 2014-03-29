package rest

type Value map[string]interface{}

const (
    STATUS  = "status"
    MESSAGE = "message"
    OK      = "ok"
    FAIL    = "fail"
    ERROR   = "error"
)

/* --- value ----------------------------------------- */

func NewValue() Value {
    return make(Value)
}

func (this Value) Success() Value {
    this[STATUS] = OK
    return this
}

func (this Value) Failure(s string) Value {
    this[STATUS] = FAIL
    this[MESSAGE] = s
    return this
}

func (this Value) Error(s string) Value {
    this[STATUS] = ERROR
    this[MESSAGE] = s
    return this
}

func (this Value) Set(key string, value interface{}) Value {
    this[key] = value
    return this
}

func (this Value) String(key string) string {
    if v, ok := this[key]; ok {
        return v.(string)
    }

    return ""
}

func (this Value) Status() string {
    return this.String(STATUS)
}
