package xproperty

import (
	"github.com/Aoi-hosizora/ahlib/xcondition"
	"strings"
)

// Apply orderBy string through PropertyDict.
func (p PropertyDict) ApplyOrderBy(source string) string {
	result := make([]string, 0)
	if source == "" {
		return ""
	}

	sources := strings.Split(source, ",")
	for _, src := range sources {
		src = strings.TrimSpace(src)
		reverse := strings.HasSuffix(src, " desc")
		src = strings.Split(src, " ")[0]

		dest, ok := p[src]
		if !ok || dest == nil || len(dest.destProps) == 0 {
			continue
		}

		if dest.revert {
			reverse = !reverse
		}
		for _, prop := range dest.destProps {
			prop += xcondition.IfThenElse(reverse, " DESC", " ASC").(string) // XXX ASC
			result = append(result, prop)
		}
	}

	return strings.Join(result, ", ")
}

// Apply orderBy string through PropertyMapper.
func (p *PropertyMapper) ApplyOrderBy(source string) string {
	return p.dict.ApplyOrderBy(source)
}

// Apply orderBy (cypher version) string through PropertyDict.
func (p PropertyDict) ApplyCypherOrderBy(parent, source string) string {
	result := make([]string, 0)
	if source == "" {
		return ""
	}

	sources := strings.Split(source, ",")
	for _, src := range sources {
		src = strings.TrimSpace(src)
		reverse := strings.HasSuffix(src, " desc")
		src = strings.Split(src, " ")[0]

		dest, ok := p[src]
		if !ok || dest == nil || len(dest.destProps) == 0 {
			continue
		}

		if dest.revert {
			reverse = !reverse
		}
		for _, prop := range dest.destProps {
			prop = parent + "." + prop + xcondition.IfThenElse(reverse, " DESC", " ASC").(string) // P.XXX ASC
			result = append(result, prop)
		}
	}

	return strings.Join(result, ", ")
}

// Apply orderBy (cypher version) string through PropertyMapper.
func (p *PropertyMapper) ApplyCypherOrderBy(parent, source string) string {
	return p.dict.ApplyCypherOrderBy(parent, source)
}
