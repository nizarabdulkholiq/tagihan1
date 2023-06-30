package controller

import (
	"github.com/aiteung/musik"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/nizarabdulkholiq/nizar"
	"github.com/nizarabdulkholiq/tagihan1/config"
	"github.com/whatsauth/whatsauth"
)

var DataTagihanRegistrasi = "tagihanregis"
var DataTagihanSPP = "tagihanspp"


func WsWhatsAuthQR(c *websocket.Conn) {
	whatsauth.RunSocket(c, config.PublicKey, config.Usertables[:], config.Ulbimariaconn)
}

func PostWhatsAuthRequest(c *fiber.Ctx) error {
	if string(c.Request().Host()) == config.Internalhost {
		var req whatsauth.WhatsauthRequest
		err := c.BodyParser(&req)
		if err != nil {
			return err
		}
		ntfbtn := whatsauth.RunModuleLegacy(req, config.PrivateKey, config.Usertables[:], config.Ulbimariaconn)
		return c.JSON(ntfbtn)
	} else {
		var ws whatsauth.WhatsauthStatus
		ws.Status = string(c.Request().Host())
		return c.JSON(ws)
	}

}

func GetHome(c *fiber.Ctx) error {
	getip := musik.GetIPaddress()
	return c.JSON(getip)
}

//	func GetUserDataNomer(c *fiber.Ctx) error {
//		getstats := gaga.GetUserData("081234567890", config.MongoConn, user)
//		return c.JSON(getstats)
//	}

func GetTagihanSPP(c*fiber.Ctx) error {
	getstatus := nizar.GetTagihanSPP("biaya_spp")
	return c.JSON(getstatus)
}