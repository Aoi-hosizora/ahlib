package xproperty

import (
	"errors"
	"reflect"
)

// PropertyMappers represents a property mappers container, contains a list of PropertyMapper.
type PropertyMappers struct {
	mappers []*PropertyMapper
}

// PropertyMapper represents a property mapper, contains the type information and property values.
type PropertyMapper struct {
	// source represents mapping `source`.
	source interface{}

	// destination represents mapping `destination`.
	destination interface{}

	// srcType represents source type.
	srcType reflect.Type

	// destType represents destination type.
	destType reflect.Type

	// dict represents the mapping properties.
	dict PropertyDict
}

// PropertyDict represents a dictionary of PropertyMapperValue.
type PropertyDict map[string]*PropertyMapperValue

// PropertyMapperValue represents a value stored in PropertyMapper, used to represents some property settings.
type PropertyMapperValue struct {
	// id represents the property order `id`.
	id int

	// revert represents if need `revert` in sort for destination.
	revert bool

	// arg represents some other arguments for property.
	arg interface{}

	// destinations represents the destination properties.
	destinations []string
}

// GetId returns the id in PropertyMapperValue.
func (p *PropertyMapperValue) GetId() int {
	return p.id
}

// GetRevert returns the revert in PropertyMapperValue.
func (p *PropertyMapperValue) GetRevert() bool {
	return p.revert
}

// GetArg returns the arg in PropertyMapperValue.
func (p *PropertyMapperValue) GetArg() interface{} {
	return p.arg
}

// GetDestinations returns the destinations in PropertyMapperValue.
func (p *PropertyMapperValue) GetDestinations() []string {
	return p.destinations
}

// New creates an empty PropertyMappers.
func New() *PropertyMappers {
	return &PropertyMappers{
		mappers: make([]*PropertyMapper, 0),
	}
}

const (
	nilModelPanic  = "xproperty: nil model"
	nilMapperPanic = "xproperty: nil mapper"
)

var (
	mapperNotFoundErr = errors.New("xproperty: mapper not found")
)

// NewMapper creates a PropertyMapper, panics when using nil model.
func NewMapper(src interface{}, dest interface{}, dict PropertyDict) *PropertyMapper {
	if src == nil || dest == nil {
		panic(nilModelPanic)
	}
	if dict == nil {
		dict = PropertyDict{}
	}
	srcTyp := reflect.TypeOf(src)
	destTyp := reflect.TypeOf(dest)

	return &PropertyMapper{source: src, destination: dest, srcType: srcTyp, destType: destTyp, dict: dict}
}

// NewValue creates a PropertyMapperValue with revert and destinations fields.
func NewValue(revert bool, destinations ...string) *PropertyMapperValue {
	return &PropertyMapperValue{
		revert:       revert,
		destinations: destinations,
	}
}

// NewValueCompletely creates a PropertyMapperValue with all fields.
func NewValueCompletely(id int, revert bool, arg interface{}, destinations ...string) *PropertyMapperValue {
	return &PropertyMapperValue{
		id:           id,
		revert:       revert,
		arg:          arg,
		destinations: destinations,
	}
}

// WithId creates a new PropertyMapperValue with id.
func (p *PropertyMapperValue) WithId(id int) *PropertyMapperValue {
	p.id = id
	return p
}

// WithRevert creates a new PropertyMapperValue with revert.
func (p *PropertyMapperValue) WithRevert(revert bool) *PropertyMapperValue {
	p.revert = revert
	return p
}

// WithArg creates a new PropertyMapperValue with arg.
func (p *PropertyMapperValue) WithArg(arg interface{}) *PropertyMapperValue {
	p.arg = arg
	return p
}

// WithDestinations creates a new PropertyMapperValue with destinations.
func (p *PropertyMapperValue) WithDestinations(destinations []string) *PropertyMapperValue {
	p.destinations = destinations
	return p
}

// GetDict returns the PropertyDict from PropertyMapper.
func (p *PropertyMapper) GetDict() PropertyDict {
	return p.dict
}

// AddMapper adds a PropertyMapper to PropertyMappers, panics when use nil PropertyMapper.
func (p *PropertyMappers) AddMapper(mapper *PropertyMapper) {
	if mapper == nil {
		panic(nilMapperPanic)
	}

	for _, m := range p.mappers {
		if m.srcType == mapper.srcType || m.destType == mapper.destType {
			m.dict = mapper.dict
			return
		}
	}
	p.mappers = append(p.mappers, &PropertyMapper{ // deep copy
		source:      mapper.source,
		destination: mapper.destination,
		srcType:     mapper.srcType,
		destType:    mapper.destType,
		dict:        mapper.dict,
	})
}

// AddMappers adds some PropertyMapper-s to PropertyMappers, panics when use nil PropertyMapper.
func (p *PropertyMappers) AddMappers(mappers ...*PropertyMapper) {
	for _, m := range mappers {
		p.AddMapper(m)
	}
}

// ====
// core
// ====

// GetMapper returns the PropertyMapper from PropertyMappers, panics when using nil model, returns error when mapper not found.
func (p *PropertyMappers) GetMapper(src interface{}, dest interface{}) (*PropertyMapper, error) {
	if src == nil || dest == nil {
		panic(nilModelPanic)
	}

	drcTyp := reflect.TypeOf(src)
	destTyp := reflect.TypeOf(dest)
	for _, mapper := range p.mappers {
		if mapper.srcType == drcTyp && mapper.destType == destTyp {
			return mapper, nil
		}
	}
	return nil, mapperNotFoundErr
}

// GetDefaultMapper returns the PropertyMapper from PropertyMappers, returns an empty PropertyDict if not found, panics when using nil model.
func (p *PropertyMappers) GetDefaultMapper(src interface{}, dest interface{}) *PropertyMapper {
	mapper, err := p.GetMapper(src, dest)
	if err != nil {
		return NewMapper(src, dest, PropertyDict{})
	}
	return mapper
}

// _mappers represents a global PropertyMappers.
var _mappers = New()

// AddMapper adds a PropertyMapper to PropertyMappers.
func AddMapper(mapper *PropertyMapper) {
	_mappers.AddMapper(mapper)
}

// AddMappers adds some PropertyMapper-s to PropertyMappers.
func AddMappers(mappers ...*PropertyMapper) {
	_mappers.AddMappers(mappers...)
}

// GetMapper returns the PropertyMapper from PropertyMappers, panics when using nil model, returns error when mapper not found.
func GetMapper(src interface{}, dest interface{}) (*PropertyMapper, error) {
	return _mappers.GetMapper(src, dest)
}

// GetDefaultMapper returns the PropertyMapper from PropertyMappers, returns an empty PropertyDict if not found, panics when using nil model.
func GetDefaultMapper(src interface{}, dest interface{}) *PropertyMapper {
	return _mappers.GetDefaultMapper(src, dest)
}
