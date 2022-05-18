package fasthttp

import (
	"strings"
	"unsafe"

	"github.com/clubpay/ronykit"
	"github.com/clubpay/ronykit/utils"
	"github.com/goccy/go-reflect"
)

// Param is a single URL parameter, consisting of a key and a value.
type Param struct {
	Key   string
	Value string
}

// Params is a Param-slice, as returned by the httpMux.
// The slice is ordered, the first URL parameter is also the first slice value.
// It is therefore safe to read values by the index.
type Params []Param

// ByName returns the value of the first Param which key matches the given name.
// If no matching Param is found, an empty string is returned.
func (ps Params) ByName(name string) string {
	for _, p := range ps {
		if p.Key == name {
			return p.Value
		}
	}

	return ""
}

type (
	DecoderFunc func(bag Params, data []byte) ronykit.Message
)

var (
	_ ronykit.RouteSelector     = Selector{}
	_ ronykit.RESTRouteSelector = Selector{}
	_ ronykit.RPCRouteSelector  = Selector{}
)

// Selector implements ronykit.RouteSelector and
// also ronykit.RPCRouteSelector and ronykit.RESTRouteSelector
type Selector struct {
	Method    string
	Path      string
	Predicate string
	Decoder   DecoderFunc
}

func (r Selector) GetMethod() string {
	return r.Method
}

func (r Selector) GetPath() string {
	return r.Path
}

func (r Selector) GetPredicate() string {
	return r.Predicate
}

func (r Selector) Query(q string) interface{} {
	switch q {
	case queryDecoder:
		return r.Decoder
	case queryMethod:
		return r.Method
	case queryPath:
		return r.Path
	case queryPredicate:
		return r.Predicate
	}

	return nil
}

var tagKey = "paramName"

// SetTag set the tag name which ReflectDecoder looks to extract parameters from Path and Query params.
// Default value: paramName
func SetTag(tag string) {
	tagKey = tag
}

// emptyInterface is the header for an interface{} value.
type emptyInterface struct {
	typ  uint64
	word unsafe.Pointer
}

type paramCaster struct {
	offset uintptr
	name   string
	opt    string
	typ    reflect.Type
}

func reflectDecoder(enc ronykit.Encoding, factory ronykit.MessageFactoryFunc) DecoderFunc {
	rVal := reflect.ValueOf(factory())
	rType := rVal.Type()
	if rType.Kind() != reflect.Ptr {
		panic("x must be a pointer to struct")
	}
	if rVal.Elem().Kind() != reflect.Struct {
		panic("x must be a pointer to struct")
	}

	var pcs []paramCaster

	for i := 0; i < reflect.Indirect(rVal).NumField(); i++ {
		f := reflect.Indirect(rVal).Type().Field(i)
		if tagValue := f.Tag.Get(tagKey); tagValue != "" {
			valueParts := strings.Split(tagValue, ",")
			if len(valueParts) == 1 {
				valueParts = append(valueParts, "")
			}

			pcs = append(
				pcs,
				paramCaster{
					offset: f.Offset,
					name:   valueParts[0],
					opt:    valueParts[1],
					typ:    f.Type,
				},
			)
		}
	}

	return func(bag Params, data []byte) ronykit.Message {
		v := factory()
		var err error
		if len(data) > 0 {
			err = ronykit.UnmarshalMessage(data, v, enc)
			if err != nil {
				return err
			}
		}

		for idx := range pcs {
			x := bag.ByName(pcs[idx].name)
			if x == "" {
				continue
			}

			ptr := unsafe.Add((*emptyInterface)(unsafe.Pointer(&v)).word, pcs[idx].offset)

			switch pcs[idx].typ.Kind() {
			case reflect.Int64:
				*(*int64)(ptr) = utils.StrToInt64(x)
			case reflect.Int32:
				*(*int32)(ptr) = utils.StrToInt32(x)
			case reflect.Uint64:
				*(*uint64)(ptr) = utils.StrToUInt64(x)
			case reflect.Uint32:
				*(*uint32)(ptr) = utils.StrToUInt32(x)
			case reflect.Int:
				*(*int)(ptr) = utils.StrToInt(x)
			case reflect.Uint:
				*(*uint)(ptr) = utils.StrToUInt(x)
			case reflect.String:
				*(*string)(ptr) = string(utils.S2B(x))
			case reflect.Bool:
				if strings.ToLower(x) == "true" {
					*(*bool)(ptr) = true
				}
			}
		}

		return v.(ronykit.Message) //nolint:forcetypeassert
	}
}
