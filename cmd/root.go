package cmd

import (
	"context"
	"flag"
	"fmt"
	"github.com/ProjectAthenaa/sonic-core/sonic"
	"github.com/ProjectAthenaa/sonic-core/sonic/database"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/license"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/metadata"
	"github.com/google/uuid"
	"github.com/prometheus/common/log"
	"strings"
)

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

func colorize(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

func Execute() {

	client := database.Connect("postgresql://doadmin:rh3rc0vgg1f706kz@athenadb-do-user-9223163-0.b.db.ondigitalocean.com:25060/defaultdb")

	//client.Schema.Create(context.Background())

	discordId := flag.String("discord_id", "715162234698858506", "Discord ID of the user to bind")
	licenseType := flag.String("type", "lifetime", "License type")

	flag.Parse()

	if discordId == nil || *discordId == "none" {
		log.Fatal("Discord ID is mandatory")
	}

	var ltype license.Type

	switch strings.ToLower(*licenseType) {
	case "lifetime":
		ltype = license.TypeLifetime
	case "renewal":
		ltype = license.TypeRenewal
	case "fnf":
		ltype = license.TypeFNF
	default:
		log.Fatal("Type needs be Lifetime, Renewal or Fnf")
	}

	meta, err := client.
		Metadata.
		Query().
		Where(
			metadata.DiscordID(*discordId),
		).
		WithUser().
		First(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	u := meta.Edges.User

	key := "ATH-" + uuid.NewString()

	if l, _ := u.QueryLicense().First(context.Background()); l != nil{
		client.License.DeleteOne(l).Exec(context.Background())
	}

	lic, err := client.
		License.
		Create().
		SetUser(u).
		SetType(ltype).
		SetKey(key).
		Save(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	if ltype != license.TypeFNF {
		builder := client.
			Stripe.
			Create().
			SetLicense(lic).
			SetCustomerID("cus_XXXXXXXXX")

		if ltype == license.TypeRenewal {
			builder.SetSubscriptionID("sub_XXXXXXXXX")
		}

		_, err = builder.Save(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}

	app, err := client.App.Create().SetUser(u).Save(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Settings.Create().SetCaptchaDetails(sonic.Map{}).SetApp(app).Save(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User has been created\nID: %s\nKey: %s\n", u.ID.String(), lic.Key)
}
