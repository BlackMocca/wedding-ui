package constants

type LocalStorageKey string

const (
	STORAGE_CONNECTION_LIST LocalStorageKey = "connection-list"
)

type AssetPath string

const (
	LOGO_NO_BACKGROUND AssetPath = "/web/resources/assets/logo/logo-no-background.svg"
	ICON_FAVOURITE     AssetPath = "/web/resources/assets/icon/favourite.png"
	ICON_LOADING       AssetPath = "/web/resources/assets/icon/loading.png"
	ICON_SUCCESS       AssetPath = "/web/resources/assets/icon/success.png"
	IMG_HOME_COVER     AssetPath = "/web/resources/assets/images/home_cover.png"
	IMG_PROMPTPAY      AssetPath = "/web/resources/assets/images/promptpay.jpg"
	IMG_CHECKMARK      AssetPath = "/web/resources/assets/images/checkmark.png"
)

var (
	SVG_RING_WEDDING_STRING   = GetSVGString("assets/icon/rings-wedding.svg")
	SVG_CALENDAR_HEART_STRING = GetSVGString("assets/icon/calendar-heart.svg")
	SVG_PIGGY_BANK_STRING     = GetSVGString("assets/icon/piggy-bank.svg")
	SVG_GIFT_STRING           = GetSVGString("assets/icon/gift.svg")
	SVG_COPY_STRING           = GetSVGString("assets/icon/copy.svg")
)
