package main

import (
	"gopkg.in/codegangsta/cli.v1"
	"gopkg.in/ini.v1"
	"log"
	"os"
)

func main() {

	log.Printf("[remap started]")

	app := cli.NewApp()
	app.Name = "remap"
	app.Usage = "remap EIPs or DNS names to instance"

	eipFlags := []cli.Flag{
		cli.StringFlag{
			Name:   "elastic-ip, e",
			Value:  "",
			Usage:  "The Elastic IP",
			EnvVar: "REMAP_ELASTIC_IP",
		},
		cli.StringFlag{
			Name:   "eip-allocation-id, a",
			Value:  "",
			Usage:  "Allocation id of IP",
			EnvVar: "REMAP_EIP_ALLOCATION_ID",
		},
		cli.IntFlag{
			Name:   "interval",
			Value:  15,
			Usage:  "Interval at which we check our IP",
			EnvVar: "REMAP_INTERVAL",
		},
		cli.StringFlag{
			Name:   "from-userdata",
			Value:  "false",
			Usage:  "Get values from user data",
			EnvVar: "REMAP_FROM_USERDATA",
		},
		cli.StringFlag{
			Name:   "run-once",
			Value:  "false",
			Usage:  "use this if you want to run from a cronjob",
			EnvVar: "REMAP_RUN_ONCE",
		},
	}

	dnsFlags := []cli.Flag{
		cli.StringFlag{
			Name:   "hosted-zone-id",
			Value:  "",
			Usage:  "route53 hosted zone id",
			EnvVar: "REMAP_HOSTED_ZONE_ID",
		},
		cli.StringFlag{
			Name:   "dns-name",
			Value:  "",
			Usage:  "dns name to remap",
			EnvVar: "REMAP_DNS_NAME",
		},
		cli.IntFlag{
			Name:   "ttl",
			Value:  300,
			Usage:  "TTL for the DNS record",
			EnvVar: "REMAP_TTL",
		},
		cli.StringFlag{
			Name:   "use-public-ip",
			Value:  "false",
			Usage:  "Map to Public IP or VPC IP",
			EnvVar: "REMAP_USE_PUBLIC_IP",
		},
		cli.IntFlag{
			Name:   "interval",
			Value:  15,
			Usage:  "Interval at which we check our IP",
			EnvVar: "REMAP_INTERVAL",
		},
		cli.StringFlag{
			Name:   "from-userdata",
			Value:  "false",
			Usage:  "Get values from user data",
			EnvVar: "REMAP_FROM_USERDATA",
		},
		cli.StringFlag{
			Name:   "run-once",
			Value:  "false",
			Usage:  "use this if you want to run from a cronjob",
			EnvVar: "REMAP_RUN_ONCE",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "eip-mode",
			Usage: "Mode to map an EIP to an instance",
			Flags: eipFlags,
			Action: func(c *cli.Context) {

				settings := RemapSettings{}

				from_userdata := c.Bool("from-userdata")
				settings.Eip = c.String("elastic-ip")
				settings.EipAllocationId = c.String("eip-allocation-id")
				settings.Interval = int64(c.Int("interval"))
				settings.RunOnce = c.Bool("run-once")

				if from_userdata {
					settings = initFromUserData(settings)
				}

				if settings.Eip == "" || settings.EipAllocationId == "" {
					log.Fatal("elastic-ip and eip-allocation-id are required arguments.")
				}

				EIPMode(settings)
			},
		},
		{
			Name:  "dns-mode",
			Usage: "Mode to map an dns name to an instance public IP",
			Flags: dnsFlags,
			Action: func(c *cli.Context) {

				settings := RemapSettings{}

				from_userdata := c.Bool("from-userdata")

				settings.HostedZoneID = c.String("hosted-zone-id")
				settings.DNSName = c.String("dns-name")
				settings.TTL = int64(c.Int("ttl"))
				settings.Interval = int64(c.Int("interval"))
				settings.UsePublicIP = c.Bool("use-public-ip")
				settings.RunOnce = c.Bool("run-once")

				if from_userdata {
					settings = initFromUserData(settings)
				}

				log.Printf("%+v", settings)

				DNSMode(settings)
			},
		},
	}

	app.Run(os.Args)
}

type RemapSettings struct {
	Mode     string
	Interval int64
	RunOnce  bool

	// eip-mode
	Eip             string
	EipAllocationId string
	IdentURL        string

	// dns-mode
	HostedZoneID string
	DNSName      string
	TTL          int64
	UsePublicIP  bool
}

func initFromUserData(remapSettings RemapSettings) RemapSettings {

	a, e := getUserData()
	if e != nil {
		log.Fatal(e)
	}

	cfg, err := ini.Load([]byte(a))
	if err != nil {
		log.Fatal(err)
	}

	section, err := cfg.GetSection("")
	if err != nil {
		log.Fatal(err)
	}

	remapSettings.Mode = section.Key("REMAP_MODE").String()
	remapSettings.Eip = section.Key("REMAP_ELASTIC_IP").String()
	remapSettings.EipAllocationId = section.Key("REMAP_EIP_ALLOCATION_ID").String()

	remapSettings.HostedZoneID = section.Key("REMAP_HOSTED_ZONE_ID").String()
	remapSettings.DNSName = section.Key("REMAP_DNS_NAME").String()

	usePublicIP, err := section.Key("REMAP_USE_PUBLIC_IP").Bool()
	if err != nil {
		usePublicIP = true
		log.Printf("Error parsing REMAP_USE_PUBLIC_IP, settings to %v", usePublicIP)
	}
	remapSettings.UsePublicIP = usePublicIP

	ttl, err := section.Key("REMAP_TTL").Int()
	if err != nil {
		ttl = 300
		log.Printf("Error parsing REMAP_TTL, settings to %v", ttl)
	}
	remapSettings.TTL = int64(ttl)

	interval, err := section.Key("REMAP_INTERVAL").Int()
	if err != nil {
		interval = 15
		log.Printf("Error parsing REMAP_INTERVAL, settings to %v", interval)
	}
	remapSettings.Interval = int64(interval)

	return remapSettings

}
