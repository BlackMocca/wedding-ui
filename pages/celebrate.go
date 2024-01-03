package pages

import (
	"github.com/Blackmocca/wedding-ui/constants"
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

var (
	svgRing = constants.SVG_RING_WEDDING_STRING
)

type Celebrate struct {
	app.Compo
	Reload        int
	celebrateText *elements.InputTextArea
	celebrateFrom *elements.InputText
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
			PlaceHolder:  "",
			ValidateFunc: []validation.ValidateRule{validation.Required},
		},
		Row: 3,
	})
	c.celebrateFrom = elements.NewInputText(c, tagCelebrateFrom, &elements.InputTextProp{
		BaseInput: elements.BaseInput{
			Id:           "celebrateFrom",
			Required:     true,
			PlaceHolder:  "",
			ValidateFunc: []validation.ValidateRule{validation.Required},
		},
	})
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
	case constants.EVENT_CLOSE_MODAL:
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

	c.Update()
}

func (c *Celebrate) Render() app.UI {
	return app.Div().Class("w-screen h-dvh bg-secondary-base").Body(
		app.Div().Class("flex flex-col w-full h-dvh justify-center items-center pt-4 gap-6").Body(
			app.Div().Class("flex flex-col w-full items-center").Body(
				app.Raw(svgRing),
				app.P().Class("text-xl text-primary-base font-medium pt-4").Text("มาร่วมอวยพร บ่าว-สาว"),
			),
			app.Div().Class("flex flex-col w-10/12").Body(
				// <label for="price" class="block text-sm font-medium leading-6 text-gray-900">Price</label>
				app.Label().Class("text-base text-primary-base font-regular pb-2").For(c.celebrateFrom.Id).Text("จาก"),
				c.celebrateFrom,
			),
			app.Div().Class("flex flex-col w-10/12").Body(
				app.Label().Class("text-base text-primary-base font-regular pb-2").For(c.celebrateText.Id).Text("เขียนคำอวยพร"),
				c.celebrateText,
			),
			/* button */
			app.Div().Class("flex flex-col w-10/12 pt-5").Body(
				elements.NewButton(constants.BUTTON_STYLE_PRIMARY).
					Text("ส่งคำอวยพร").
					OnClick(c.save),
			),
		),

		app.P().Class("absolute w-full text-sm text-primary-base font-medium text-center bottom-0 pb-4").Text("© 2024 NengHuag Wedding. All Rights Reserved"),
	)
}
