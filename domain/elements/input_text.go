package elements

import (
	"github.com/Blackmocca/wedding-ui/constants"
	"github.com/Blackmocca/wedding-ui/domain/core"
	"github.com/Blackmocca/wedding-ui/domain/core/validation"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type InputTextProp struct {
	BaseInput
	inputType constants.InputType
}

type inputState struct {
	value         string
	isValidateErr bool
}

type InputText struct {
	app.Compo
	Parent core.ParentNotify
	Tag    string
	InputTextProp

	state inputState
}

func NewInputText(parent core.ParentNotify, tag string, prop *InputTextProp) *InputText {
	return &InputText{
		Parent: parent,
		Tag:    tag,
		InputTextProp: InputTextProp{
			BaseInput: prop.BaseInput,
			inputType: constants.INPUT_TYPE_TEXT,
		},
		state: inputState{
			value:         prop.Value,
			isValidateErr: prop.ValidateError != nil,
		},
	}
}

func NewInputPassword(parent core.ParentNotify, tag string, prop *InputTextProp) *InputText {
	return &InputText{
		Parent: parent,
		Tag:    tag,
		InputTextProp: InputTextProp{
			BaseInput: prop.BaseInput,
			inputType: constants.INPUT_TYPE_TEXT,
		},
		state: inputState{
			value:         prop.Value,
			isValidateErr: prop.ValidateError != nil,
		},
	}
}

func (i *InputText) GetValue() string {
	return i.state.value
}

func (i *InputText) onChangeInput(ctx app.Context, e app.Event) {
	value := ctx.JSSrc().Get("value")
	validateErr := validation.Validate(value.String(), i.ValidateFunc...)
	i.state.value = value.String()
	i.state.isValidateErr = (validateErr != nil)

	i.Parent.Event(nil, constants.EVENT_ON_VALIDATE_INPUT, i)

	e.PreventDefault()
}

func (i *InputText) Render() app.UI {
	class := "w-full leading-6 border-2 border-primary-base px-2 py-1 h-11 focus:border-green-600 focus:outline-none"
	if i.state.isValidateErr || i.BaseInput.ValidateError != nil || i.ValidateError != nil {
		class += " border-red-500 "
	}
	return app.Input().
		ID(i.Id).
		Class(class).
		Disabled(i.Disabled).
		Type(string(i.inputType)).
		Value(i.state.value).
		Placeholder(i.PlaceHolder).
		Required(i.Required).
		OnChange(i.onChangeInput)
}
