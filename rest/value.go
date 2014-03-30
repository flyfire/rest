package rest

import "strconv"

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
        if s, ok := v.(string); ok {
            return s
        }
    }

    return ""
}

func (this Value) Int(key string) int {
    if v, ok := this[key]; ok {
        switch v.(type) {
        case string:
            if n, err := strconv.Atoi(v.(string)); err == nil {
                return n
            }
        case int:
            return v.(int)
        }
    }

    return 0
}

func (this Value) Status() string {
    return this.String(STATUS)
}
