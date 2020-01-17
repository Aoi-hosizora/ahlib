package xmapper

import (
	"reflect"
)

// public: MapFunc ExtraMapFunc EntityMapper DisposableMapOption

// Map Function from interface{} (is fromModel type) to interface{} (is toModel field type)
type MapFunc func(interface{}) interface{}

// Map Function, the last process of map, using fromObject and toObject
type ExtraMapFunc func(interface{}, interface{}) interface{}

// Save all mapper between entities
type EntityMapper struct {
	_entities []*entity
}

// Save all map rule between specific _fromType and _toType entity type
type entity struct {
	_mapper *EntityMapper

	_fromType reflect.Type
	_toType   reflect.Type

	// Save map rule between specific field
	// *_fieldDirectMapRule, *_fieldFromMapRule, _mapFunc
	_rules []_mapRule
}

// specific structField:
type _mapRule interface{}

// Save the direct map rule
type _fieldDirectMapRule struct {
	_toField reflect.StructField
	_mapFunc MapFunc
}

// Save the copy / nest map rule
type _fieldFromMapRule struct {
	_fromField reflect.StructField
	_toField   reflect.StructField
	_isNest    bool
}

type DisposableMapOption struct {
	_fromType reflect.Type
	_toType   reflect.Type
	_mapFunc  ExtraMapFunc
}

// Create a entity from entitiesMapper
func NewEntityMapper() *EntityMapper {
	return new(EntityMapper)
}
