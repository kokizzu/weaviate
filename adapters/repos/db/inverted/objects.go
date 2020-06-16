//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

package inverted

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema"
)

func (a *Analyzer) Object(input map[string]interface{}, props []*models.Property) ([]Property, error) {
	propsMap := map[string]*models.Property{}
	for _, prop := range props {
		propsMap[prop.Name] = prop
	}

	out := make([]Property, len(input))
	i := 0
	for key, value := range input {
		prop, ok := propsMap[key]
		if !ok {
			// not a registered property, skip
			continue
		}

		if len(prop.DataType) != 1 {
			// must be a ref prop or something else is wrong, skip
			continue
		}
		var hasFrequency bool

		var items []Countable
		switch schema.DataType(prop.DataType[0]) {
		case schema.DataTypeText:
			hasFrequency = true
			asString, ok := value.(string)
			if !ok {
				return nil, fmt.Errorf("expected property %s to be of type string, but got %T", key, value)
			}
			items = a.Text(asString)
		case schema.DataTypeString:
			hasFrequency = true
			asString, ok := value.(string)
			if !ok {
				return nil, fmt.Errorf("expected property %s to be of type string, but got %T", key, value)
			}
			items = a.String(asString)
		case schema.DataTypeInt:
			hasFrequency = false
			asInt, ok := value.(int64)
			if !ok {
				return nil, fmt.Errorf("expected property %s to be of type int64, but got %T", key, value)
			}

			var err error
			items, err = a.Int(int(asInt))
			if err != nil {
				return nil, errors.Wrapf(err, "analyze property %s", key)
			}
		case schema.DataTypeNumber:
			hasFrequency = false
			asFloat, ok := value.(float64)
			if !ok {
				return nil, fmt.Errorf("expected property %s to be of type float64, but got %T", key, value)
			}

			var err error
			items, err = a.Float(asFloat) // convert to int before analyzing
			if err != nil {
				return nil, errors.Wrapf(err, "analyze property %s", key)
			}
		case schema.DataTypeBoolean:
			hasFrequency = false
			asBool, ok := value.(bool)
			if !ok {
				return nil, fmt.Errorf("expected property %s to be of type bool, but got %T", key, value)
			}

			var err error
			items, err = a.Bool(asBool) // convert to int before analyzing
			if err != nil {
				return nil, errors.Wrapf(err, "analyze property %s", key)
			}

		default:
			// ignore unsupported prop type
			continue
		}

		out[i] = Property{
			Name:         key,
			Items:        items,
			HasFrequency: hasFrequency,
		}
		i++
	}

	return out[:i], nil
}