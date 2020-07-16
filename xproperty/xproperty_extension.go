package xproperty

import (
	"github.com/Aoi-hosizora/ahlib/xcondition"
	"strings"
)

func (p *PropertyMapper) ApplyOrderBy(source string) string {
	result := make([]string, 0)
	if source == "" {
		return ""
	}

	sources := strings.Split(source, ",")
	for _, src := range sources {
		src = strings.TrimSpace(src)
		reverse := strings.HasSuffix(src, " desc")
		src = strings.Split(src, " ")[0]

		dest, ok := p.dict[src]
		if !ok || dest == nil || len(dest.destProps) == 0 {
			continue
		}

		if dest.revert {
			reverse = !reverse
		}
		props := dest.destProps

		for _, prop := range props {
			prop += xcondition.IfThenElse(reverse, " DESC", " ASC").(string)
			result = append(result, prop)
		}
	}

	return strings.Join(result, ", ")
}
