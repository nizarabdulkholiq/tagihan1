package controller

import (
	"net/http"

	"github.com/aiteung/musik"
	gege "github.com/gocroot/kampus/model"
	gaga "github.com/gocroot/kampus/module"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/harisriyoni3/iteung/config"
	"github.com/whatsauth/whatsauth"
)

var suratdek = "surat"
var user = "users"

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
func GetSurat(c *fiber.Ctx) error {
	getstats := gaga.GetSurat("Kamu", config.MongoConn, user)
	return c.JSON(getstats)
}

func GetUserDataNama(c *fiber.Ctx) error {
	getstats := gaga.GetNamaUser("Haris Riyoni", config.MongoConn, user)
	return c.JSON(getstats)
}

func InsertDataSuratChat(c *fiber.Ctx) error {
	database := config.MongoConn
	var srt gege.Surat
	if err := c.BodyParser(&srt); err != nil {
		return err
	}
	Inserted := gaga.InsertSuratChat(database,
		suratdek,
		srt.Isisurat,
		srt.Subject,
	)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": Inserted,
	})
}
