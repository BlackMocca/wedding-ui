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

type Celebrate struct {
	app.Compo

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
	for _, err := range allValidates {
		if err != nil {
			return false
		}
	}

	return true
}

func (c Celebrate) save(ctx app.Context, e app.Event) {
	if !c.isValidatePass() {
		app.Log("validate fail")
		return
	}

	var celebrateText = c.celebrateText.GetValue()
	var celebrateFrom = c.celebrateFrom.GetValue()

	// celebrate := models.NewCelebrate(celebrateText, celebrateFrom)
	// if err := shared.GetPsqlClient().Create(context.Background(), celebrate); err != nil {
	// 	app.Log("fail to save celebrate text:", err.Error())
	// }

	if err := api.WeddingAPI.Create(ctx, celebrateText, celebrateFrom); err != nil {
		panic(err)
	}

	app.Log(celebrateText, celebrateFrom)
}

func (c Celebrate) Render() app.UI {
	return app.Div().Class("w-screen h-dvh overflow-hidden bg-secondary-base").Body(
		app.A().Href("/").Text("this is "),
		app.P().Text("asdsadsa"),
		c.celebrateText,
		c.celebrateFrom,
		/* button */
		app.Div().Class("col-span-2 flex flex-row items-center justify-end gap-4").Body(
			elements.NewButton(constants.BUTTON_STYLE_SECONDARY).
				Text("Submit").
				OnClick(c.save),
		),
	)
}
