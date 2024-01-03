package pages

import (
	"github.com/Blackmocca/wedding-ui/constants"
	"github.com/Blackmocca/wedding-ui/domain/components"
	"github.com/Blackmocca/wedding-ui/domain/core"
	"github.com/Blackmocca/wedding-ui/domain/core/api"
	"github.com/Blackmocca/wedding-ui/domain/core/validation"
	"github.com/Blackmocca/wedding-ui/domain/elements"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	tagCelebrateInput = "CelebrateText"
	tagCelebrateFrom  = "CelebrateFrom"
)

type Celebrate struct {
	app.Compo
	Reload             int
	celebrateText      *elements.InputTextArea
	celebrateFrom      *elements.InputText
	modal              *components.SuccessModal
	isModalSuccessShow bool
}

func (c *Celebrate) CelebrateText() *elements.InputTextArea {
	return c.celebrateText
}
func (c *Celebrate) CelebrateFrom() *elements.InputText {
	return c.celebrateFrom
}

func (c *Celebrate) OnInit() {
	c.celebrateText = elements.NewInputTextArea(c, tagCelebrateInput, &elements.InputTextAreaProp{
		BaseInput: elements.BaseInput{
			Id:           "celebrateText",
			Required:     true,
			PlaceHolder:  "เขียนคำอวยพรถึง บ่าวสาว ที่นี้",
			ValidateFunc: []validation.ValidateRule{validation.Required},
		},
		Row: 3,
	})
	c.celebrateFrom = elements.NewInputText(c, tagCelebrateFrom, &elements.InputTextProp{
		BaseInput: elements.BaseInput{
			Id:           "celebrateFrom",
			Required:     true,
			PlaceHolder:  "ลงชื่อผู้อวยพร",
			ValidateFunc: []validation.ValidateRule{validation.Required},
		},
	})
	c.modal = components.NewSuccessModal("ขอบคุณสำหรับการอวยพรพวกเราครับ")
	c.isModalSuccessShow = true
}

func (c *Celebrate) Event(ctx app.Context, event constants.Event, data interface{}) {
	switch event {
	case constants.EVENT_ON_VALIDATE_INPUT:
		if v, ok := data.(*elements.InputTextArea); ok {
			elem := core.CallMethod(c, v.Tag).(*elements.InputTextArea)
			elem.Value = elem.GetValue()
			elem.ValidateError = validation.Validate(elem.GetValue(), elem.ValidateFunc...)
		}
		if v, ok := data.(*elements.InputText); ok {
			elem := core.CallMethod(c, v.Tag).(*elements.InputText)
			elem.Value = elem.GetValue()
			elem.ValidateError = validation.Validate(elem.GetValue(), elem.ValidateFunc...)
		}
	}
}

func (c *Celebrate) isValidatePass() bool {
	c.Event(nil, constants.EVENT_ON_VALIDATE_INPUT, c.celebrateText)
	c.Event(nil, constants.EVENT_ON_VALIDATE_INPUT, c.celebrateFrom)

	var allValidates = []error{
		c.celebrateText.ValidateError,
		c.celebrateFrom.ValidateError,
	}
	var isError error
	for _, err := range allValidates {
		if err != nil {
			isError = err
		}
	}

	return isError == nil
}

func (c *Celebrate) save(ctx app.Context, e app.Event) {
	if !c.isValidatePass() {
		app.Log("validate fail")
		c.Update()
		return
	}

	var celebrateText = c.celebrateText.GetValue()
	var celebrateFrom = c.celebrateFrom.GetValue()

	if err := api.WeddingAPI.Create(ctx, celebrateText, celebrateFrom); err != nil {
		panic(err)
	}

	c.isModalSuccessShow = true
	c.Update()
}

func (c *Celebrate) Render() app.UI {
	return app.Div().Class("w-screen h-dvh overflow-hidden bg-secondary-base").Body(
		app.If(c.isModalSuccessShow, c.modal),
		c.celebrateText,
		c.celebrateFrom,
		/* button */
		app.Div().Class("flex flex-row items-center justify-end").Body(
			elements.NewButton(constants.BUTTON_STYLE_SECONDARY).
				Text("Submit").
				OnClick(c.save),
		),
	)
}
