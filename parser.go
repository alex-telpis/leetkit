package leetkit

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var errCycle = errors.New("cycle detected")

// Parse parses Leetcode's string representation of a data structure or primitive type.
// For simplicity, Parse panics instead of returning an error.
// For parsing complex types, consider using dedicated functions like ParseTreeNode, ParseIntMatrix, etc.
func Parse[T any](s string) T { return panicOnParseErr(s, tryParse[T]) }

func tryParse[T any](s string) (T, error) {
	var res T

	s = strings.TrimSpace(s)
	t := reflect.TypeOf(res)
	v, err := parseReflect(s, t)
	if err != nil {
		return res, err
	}

	rv := reflect.ValueOf(&res)
	rv.Elem().Set(v)

	return res, nil
}

// Sprint returns a string representation of the input.
// The name is referring to the fmt.Sprint function.
// For usage simplicity, Sprint panics instead of returning an error.
func Sprint(v any) string {
	res, err := trySprint(v)
	if err != nil {
		panic("failed to serialize: " + err.Error())
	}

	return res
}

func trySprint(v any) (string, error) {
	if str, ok := v.(fmt.Stringer); ok {
		return str.String(), nil
	}

	s, err := sprintReflect(reflect.ValueOf(v))
	if err != nil {
		return "", err
	}

	return s, nil
}

func parseReflect(raw string, refType reflect.Type) (reflect.Value, error) {
	refVal := reflect.Value{}

	switch refType.Kind() {
	case reflect.Bool:
		if raw != "true" && raw != "false" {
			return refVal, fmt.Errorf("invalid bool input: %s", raw)
		}
		b := raw == "true"
		return reflect.ValueOf(b), nil

	case reflect.Uint8: // byte
		if len(raw) != 3 || raw[0] != '"' && raw[0] != '\'' || raw[2] != raw[0] {
			return refVal, fmt.Errorf("invalid byte input: %s", raw)
		}
		return reflect.ValueOf(raw[1]), nil

	case reflect.String:
		s, err := strconv.Unquote(raw)
		if err != nil {
			return refVal, fmt.Errorf("invalid string input: %s", raw)
		}
		return reflect.ValueOf(s), nil

	case reflect.Int, reflect.Int32:
		i, err := strconv.Atoi(raw)
		if err != nil {
			return refVal, fmt.Errorf("invalid int input: %s", raw)
		}
		return reflect.ValueOf(i), nil

	case reflect.Int64:
		i, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			return refVal, fmt.Errorf("invalid int64 input: %s", raw)
		}
		return reflect.ValueOf(i), nil

	case reflect.Uint, reflect.Uint32:
		i, err := strconv.ParseUint(raw, 10, 32)
		if err != nil {
			return refVal, fmt.Errorf("invalid uint input: %s", raw)
		}
		return reflect.ValueOf(uint(i)), nil

	case reflect.Uint64:
		i, err := strconv.ParseUint(raw, 10, 64)
		if err != nil {
			return refVal, fmt.Errorf("invalid uint64 input: %s", raw)
		}
		return reflect.ValueOf(i), nil

	case reflect.Float64:
		f, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			return refVal, fmt.Errorf("invalid float64 input: %s", raw)
		}
		return reflect.ValueOf(f), nil

	case reflect.Slice:
		splits, err := split(raw)
		if err != nil {
			return refVal, fmt.Errorf("invalid slice input: %s", raw)
		}

		sl := reflect.MakeSlice(refType, 0, len(splits))
		for _, s := range splits {
			e, err := parseReflect(s, refType.Elem())
			if err != nil {
				return refVal, err
			}
			sl = reflect.Append(sl, e)
		}
		return sl, nil

	case reflect.Ptr:
		elemType := refType.Elem()
		switch elemType {
		case reflect.TypeOf((*TreeNode)(nil)).Elem():
			root, err := tryParseTreeNode(raw)
			if err != nil {
				return refVal, err
			}
			return reflect.ValueOf(root), nil

		case reflect.TypeOf((*ListNode)(nil)).Elem():
			head, err := tryParseListNode(raw)
			if err != nil {
				return refVal, err
			}
			return reflect.ValueOf(head), nil

		case reflect.TypeOf((*NaryTreeNode)(nil)).Elem():
			head, err := tryParseNaryTreeNode(raw)
			if err != nil {
				return refVal, err
			}
			return reflect.ValueOf(head), nil

		}
	}

	return refVal, fmt.Errorf("unexpected type %s", refType.String())
}

func sprintReflect(v reflect.Value) (string, error) {
	switch v.Kind() {
	case reflect.Slice:
		sb := &strings.Builder{}
		sb.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				sb.WriteByte(',')
			}
			_s, err := sprintReflect(v.Index(i))
			if err != nil {
				return "", err
			}
			sb.WriteString(_s)
		}
		sb.WriteByte(']')
		return sb.String(), nil

	case reflect.Ptr:
		switch tt := v.Elem().Interface().(type) {
		case TreeNode:
			return tt.String(), nil

		case NaryTreeNode:
			return tt.String(), nil

		case ListNode:
			return tt.String(), nil

		default:
			return "", fmt.Errorf("unexpected type %T", v.Elem().Interface())
		}

	case reflect.String:
		return fmt.Sprintf(`"%s"`, v.Interface()), nil

	case reflect.Uint8:
		return fmt.Sprintf(`"%c"`, v.Interface()), nil

	case reflect.Float64:
		return fmt.Sprintf(`%.5f`, v.Interface()), nil

	default:
		return fmt.Sprintf(`%v`, v.Interface()), nil
	}
}

func split(s string) ([]string, error) {
	s = strings.TrimSpace(s)
	invalidErr := fmt.Errorf("invalid slice input: %s", s)

	if len(s) <= 1 || s[0] != '[' || s[len(s)-1] != ']' {
		return nil, invalidErr
	}

	var splits []json.RawMessage
	if err := json.Unmarshal([]byte(s), &splits); err != nil {
		return nil, invalidErr
	}
	res := make([]string, len(splits))
	for i, v := range splits {
		res[i] = string(v)
	}
	return res, nil
}

// panicOnParseErr is a wrapper around a parser function that will panic instead of forwarding the error.
func panicOnParseErr[T any](input string, f func(string) (T, error)) T {
	res, err := f(input)
	if err != nil {
		panic("failed to parse: " + err.Error())
	}

	return res
}
