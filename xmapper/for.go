package xmapper

// Add direct field mapper rule, ignore if toField is not exist
func (e *entity) ForMember(toFieldString string, mapFunc MapFunc) *entity {
	toField, ok := e._toType.FieldByName(toFieldString)
	if !ok {
		return e
	}
	rule := &_fieldDirectMapRule{
		_toField: toField,
		_mapFunc: mapFunc,
	}
	e._rules = append(e._rules, rule)
	return e
}

// Add nest field mapper rule, ignore if fromField or toField is not exist
func (e *entity) ForNest(fromFieldString string, toFieldString string) *entity {
	fromField, ok := e._fromType.FieldByName(fromFieldString)
	if !ok {
		return e
	}
	toField, ok := e._toType.FieldByName(toFieldString)
	if !ok {
		return e
	}
	rule := &_fieldFromMapRule{
		_fromField: fromField,
		_toField:   toField,
		_isNest:    true,
	}
	e._rules = append(e._rules, rule)
	return e
}

// Add copy field mapper rule, ignore if (fromField or toField is not exist) or (field type if different)
func (e *entity) ForCopy(fromFieldString string, toFieldString string) *entity {
	fromField, ok := e._fromType.FieldByName(fromFieldString)
	if !ok {
		return e
	}
	toField, ok := e._toType.FieldByName(toFieldString)
	if !ok {
		return e
	}
	if fromField.Type != toField.Type {
		return e
	}
	rule := &_fieldFromMapRule{
		_fromField: fromField,
		_toField:   toField,
		_isNest:    false,
	}
	e._rules = append(e._rules, rule)
	return e
}

// Add extra function mapper rule
func (e *entity) ForExtra(extraMapFunc ExtraMapFunc) *entity {
	e._rules = append(e._rules, extraMapFunc)
	return e
}
