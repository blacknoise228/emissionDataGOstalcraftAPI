package getData

import (
	"fmt"
	"stalcraftbot/internal/logs"

	"strings"

	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

var PromoText string

// Parsing URL and returned actual promocodes
func ParseFunc() string {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{"https://steamcommunity.com/sharedfiles/filedetails/?id=3286541385"},
		ParseFunc: parsePromo,
	}).Start()
	fmt.Println(PromoText)
	logs.Logger.Debug().Msg("ParsePromo done")
	return PromoText
}

func parsePromo(g *geziyor.Geziyor, r *client.Response) {
	promo, _ := r.HTMLDoc.Find("div.subSectionDesc").Html()
	text := strings.ReplaceAll(promo, "<br/>", "\n")
	text = strings.ReplaceAll(text, "\t", "")
	text = strings.ReplaceAll(text, `<div style="clear: both"></div>`, "")
	PromoText = strings.TrimSpace(text)

}
