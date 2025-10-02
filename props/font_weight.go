package props

type FontWeight string

const (
	FontWeightThin       FontWeight = "100"
	FontWeight100        FontWeight = "100"
	FontWeightExtraLight FontWeight = "200"
	FontWeight200        FontWeight = "200"
	FontWeightLight      FontWeight = "300"
	FontWeight300        FontWeight = "300"
	FontWeightNormal     FontWeight = "400"
	FontWeight400        FontWeight = "400"
	FontWeightMedium     FontWeight = "500"
	FontWeight500        FontWeight = "500"
	FontWeightSemiBold   FontWeight = "600"
	FontWeight600        FontWeight = "600"
	FontWeightBold       FontWeight = "700"
	FontWeight700        FontWeight = "700"
	FontWeightExtraBold  FontWeight = "800"
	FontWeight800        FontWeight = "800"
	FontWeightBlack      FontWeight = "900"
	FontWeight900        FontWeight = "900"
)

func (f FontWeight) String() string {
	return string(f)
}
