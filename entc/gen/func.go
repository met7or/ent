package gen

import (
	"fmt"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"unicode"

	"fbc/ent/field"

	"github.com/go-openapi/inflect"
)

var (
	rules   = ruleset()
	acronym = make(map[string]bool)
	funcs   = template.FuncMap{
		"ops":        ops,
		"add":        add,
		"order":      order,
		"snake":      snake,
		"pascal":     pascal,
		"extend":     extend,
		"xrange":     xrange,
		"receiver":   receiver,
		"plural":     plural,
		"aggregate":  aggregate,
		"primitives": primitives,
		"singular":   rules.Singularize,
		"quote":      strconv.Quote,
		"base":       filepath.Base,
		"keys":       keys,
		"join":       join,
		"lower":      strings.ToLower,
		"upper":      strings.ToUpper,
		"hasSuffix":  strings.HasSuffix,
	}
)

// ops returns all operations for given type.
func ops(f *Field) []Op {
	switch t := f.Type; {
	case t == field.TypeBool:
		return boolOps
	case t == field.TypeString && strings.ToLower(f.Name) != "id":
		return stringOps
	default:
		return numericOps
	}
}

// xrange generates a slice of len n.
func xrange(n int) (a []int) {
	for i := 0; i < n; i++ {
		a = append(a, i)
	}
	return
}

// plural a name.
func plural(name string) string {
	p := rules.Pluralize(name)
	if p == name {
		p += "Slice"
	}
	return p
}

// pascal converts the given column name into a PascalCase.
//
//	user_info => UserInfo
//	full_name => FullName
//  user_id   => UserID
//
func pascal(s string) string {
	words := strings.Split(s, "_")
	for i, w := range words {
		upper := strings.ToUpper(w)
		if acronym[upper] {
			words[i] = upper
		} else {
			words[i] = rules.Capitalize(w)
		}
	}
	return strings.Join(words, "")
}

// snake converts the given struct or field name into a snake_case.
//
//	Username => username
//	FullName => full_name
//	HTTPCode => http_code
//
func snake(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		// put '_' if it is not a start or end of a word, current letter is an uppercase letter,
		// and previous letter is a lowercase letter (cases like: "UserInfo"), or next letter is
		// also a lowercase letter and previous letter is not "_".
		if i > 0 && i < len(s)-1 && unicode.IsUpper(r) &&
			(unicode.IsLower(rune(s[i-1])) ||
				unicode.IsLower(rune(s[i+1])) && unicode.IsLetter(rune(s[i-1]))) {
			b.WriteString("_")
		}
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}

// receiver returns the receiver name of the given type.
//
//	User      => u
//	UserQuery => uq
//
func receiver(s string) (r string) {
	words := strings.Split(snake(s), "_")
	for _, w := range words {
		r += w[:1]
	}
	return
}

// scope wraps the Type object with extended context.
type scope struct {
	*Type
	Scope map[interface{}]interface{}
}

// extend extends the parent block with a KV pairs.
//
//	{{ with $scope := extend $ "key" "value" }}
//		{{ template "setters" $scope }}
//	{{ end}}
//
func extend(t *Type, kv ...interface{}) *scope {
	s := &scope{Type: t, Scope: make(map[interface{}]interface{})}
	if len(kv)%2 != 0 {
		panic("invalid number of parameters")
	}
	for i := 0; i < len(kv); i += 2 {
		s.Scope[kv[i]] = kv[i+1]
	}
	return s
}

// add calculates summarize list of variables.
func add(xs ...int) (n int) {
	for _, x := range xs {
		n += x
	}
	return
}

func ruleset() *inflect.Ruleset {
	rules := inflect.NewDefaultRuleset()
	// add common initialisms. copied from golint.
	for _, w := range []string{
		"API", "ASCII", "CPU", "CSS", "DNS", "GUID", "UID", "UI",
		"RHS", "RPC", "SLA", "SMTP", "SSH", "TLS", "TTL", "HTML",
		"HTTP", "HTTPS", "ID", "IP", "JSON", "LHS", "QPS", "RAM",
		"UUID", "URI", "URL", "UTF8", "VM", "XML", "XSRF", "XSS",
	} {
		acronym[w] = true
		rules.AddAcronym(w)
	}
	return rules
}

// order returns a map of sort orders.
// The key is the function name, and the value its database keyword.
func order() map[string]string {
	return map[string]string{
		"asc":  "incr",
		"desc": "decr",
	}
}

// aggregate returns a map between all agg-functions and if they accept a field name as a parameter or not.
func aggregate() map[string]bool {
	return map[string]bool{
		"min":   true,
		"max":   true,
		"sum":   true,
		"mean":  true,
		"count": false,
	}
}

// keys returns the given map keys.
func keys(v reflect.Value) ([]string, error) {
	if k := v.Type().Kind(); k != reflect.Map {
		return nil, fmt.Errorf("expect map for keys, got: %v", k)
	}
	keys := make([]string, v.Len())
	for i, v := range v.MapKeys() {
		keys[i] = v.String()
	}
	return keys, nil
}

// primitives returns all primitives types.
func primitives() []string {
	return []string{field.TypeString.String(), field.TypeInt.String(), field.TypeFloat64.String(), field.TypeBool.String()}
}

// join is a wrapper around strings.Join to provide consistent output.
func join(a []string, sep string) string {
	sort.Strings(a)
	return strings.Join(a, sep)
}