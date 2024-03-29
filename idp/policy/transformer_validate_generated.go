// NOTE: automatically generated file -- DO NOT EDIT

package policy

import (
	"userclouds.com/infra/ucerr"
)

// Validate implements Validateable
func (o Transformer) Validate() error {
	fieldLen := len(o.Name)
	if fieldLen < 1 || fieldLen > 128 {
		return ucerr.Friendlyf(nil, "Transformer.Name length has to be between 1 and 128 (length: %v)", fieldLen)
	}
	if err := o.InputType.Validate(); err != nil {
		return ucerr.Wrap(err)
	}
	if err := o.InputConstraints.Validate(); err != nil {
		return ucerr.Wrap(err)
	}
	if err := o.OutputConstraints.Validate(); err != nil {
		return ucerr.Wrap(err)
	}
	if err := o.TransformType.Validate(); err != nil {
		return ucerr.Wrap(err)
	}
	// .extraValidate() lets you do any validation you can't express in codegen tags yet
	if err := o.extraValidate(); err != nil {
		return ucerr.Wrap(err)
	}
	return nil
}
