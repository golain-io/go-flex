/**
* 	part.go has been copied and modified from github.com/Masterminds/squirrel
 */

package goflex

import (
	"fmt"
	"io"
)

type part struct {
	pred interface{}
	args []interface{}
}

func newPart(pred interface{}, args ...interface{}) Fluxer {
	return &part{pred, args}
}

func (p part) ToFlux() (flux string, args []interface{}, err error) {
	switch pred := p.pred.(type) {
	case nil:
		// no-op
	case Fluxer:
		flux, args, err = nestedToFlux(pred)
	case string:
		flux = pred
		args = p.args
	default:
		err = fmt.Errorf("expected string or Fluxer, not %T", pred)
	}
	return
}

func nestedToFlux(f Fluxer) (string, []interface{}, error) {
	// @TODO: Do we need rawFluxer?
	// if raw, ok := f.(rawSqlizer); ok {
	// 	return raw.toSqlRaw()
	// } else {
	return f.ToFlux()
}

func appendToFlux(parts []Fluxer, w io.Writer, sep string, args []interface{}) ([]interface{}, error) {
	for i, p := range parts {
		partSql, partArgs, err := nestedToFlux(p)
		if err != nil {
			return nil, err
		} else if len(partSql) == 0 {
			continue
		}

		if i > 0 {
			_, err := io.WriteString(w, sep)
			if err != nil {
				return nil, err
			}
		}

		_, err = io.WriteString(w, partSql)
		if err != nil {
			return nil, err
		}
		args = append(args, partArgs...)
	}
	return args, nil
}
