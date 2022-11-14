package utils

import (
	"reflect"

	"github.com/pkg/errors"

	"github.com/parinyapt/Rail_Project_Backend/models"
)

func GetStructTag(config models.ParameterGetStructTag) (string, error) {
	field, ok := reflect.TypeOf(config.Structx).FieldByName(config.FieldName)
	if !ok {
		return "", errors.Wrap(errors.New("[Custom]->Error=False"), "[Error]->Reflect Field Error")
	}

	return string(field.Tag.Get(config.TagName)), nil
}
