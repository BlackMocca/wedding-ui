package constants

type InputType string

const (
	INPUT_TYPE_TEXT     InputType = "text"
	INPUT_TYPE_PASSWORD InputType = "password"
)

type ButtonStyle int

const (
	// class primary button style
	BUTTON_STYLE_PRIMARY ButtonStyle = iota

	// class secondary button style
	BUTTON_STYLE_SECONDARY
)
