package zones

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"testing"
	"time"
)

func TestZones(t *testing.T) {
	for _, zone := range []string{
		UTC,
		Local,

		Etc_GMT_P12,
		Etc_GMT_P11,
		Pacific_Honolulu,
		America_Anchorage,
		America_Santa_Isabel,
		America_Los_Angeles,
		America_Chihuahua,
		America_Phoenix,
		America_Denver,
		America_Guatemala,
		America_Chicago,
		America_Regina,
		America_Mexico_City,
		America_Bogota,
		America_Indiana_Indianapolis,
		America_New_York,
		America_Caracas,
		America_Halifax,
		America_Asuncion,
		America_La_Paz,
		America_Cuiaba,
		America_Santiago,
		America_St_Johns,
		America_Sao_Paulo,
		America_Godthab,
		America_Cayenne,
		America_Argentina_Buenos_Aires,
		America_Montevideo,
		Etc_GMT_P2,
		Atlantic_Cape_Verde,
		Atlantic_Azores,
		Africa_Casablanca,
		Atlantic_Reykjavik,
		Europe_London,
		Etc_GMT,
		Europe_Berlin,
		Europe_Paris,
		Africa_Lagos,
		Europe_Budapest,
		Europe_Warsaw,
		Africa_Windhoek,
		Europe_Istanbul,
		Europe_Kiev,
		Africa_Cairo,
		Asia_Damascus,
		Asia_Amman,
		Africa_Johannesburg,
		Asia_Jerusalem,
		Asia_Beirut,
		Asia_Baghdad,
		Europe_Minsk,
		Asia_Riyadh,
		Africa_Nairobi,
		Asia_Tehran,
		Europe_Moscow,
		Asia_Tbilisi,
		Asia_Yerevan,
		Asia_Dubai,
		Asia_Baku,
		Indian_Mauritius,
		Asia_Kabul,
		Asia_Tashkent,
		Asia_Karachi,
		Asia_Colombo,
		Asia_Kolkata,
		Asia_Kathmandu,
		Asia_Almaty,
		Asia_Dhaka,
		Asia_Yekaterinburg,
		Asia_Yangon,
		Asia_Bangkok,
		Asia_Novosibirsk,
		Asia_Krasnoyarsk,
		Asia_Ulaanbaatar,
		Asia_Shanghai,
		Australia_Perth,
		Asia_Singapore,
		Asia_Taipei,
		Asia_Irkutsk,
		Asia_Seoul,
		Asia_Tokyo,
		Australia_Darwin,
		Australia_Adelaide,
		Australia_Hobart,
		Asia_Yakutsk,
		Australia_Brisbane,
		Pacific_Port_Moresby,
		Australia_Sydney,
		Asia_Vladivostok,
		Pacific_Guadalcanal,
		Etc_GMT_M12,
		Pacific_Fiji,
		Asia_Magadan,
		Pacific_Auckland,
		Pacific_Tongatapu,
		Pacific_Apia,
	} {
		_, err := time.LoadLocation(zone)
		xtesting.Nil(t, err)
	}
}
