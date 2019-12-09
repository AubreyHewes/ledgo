package ledgo

import (
	"encoding/hex"
	"log"

	"github.com/google/gousb"
)

type SupportedDevice struct {
	pid           gousb.ID
	description   string
	url           string
	supportsLogo  bool
	supportsWheel bool
}

func NewClient() {
	// Initialize a new Context.
	ctx := gousb.NewContext()
	defer ctx.Close()

	// Iterate through available Devices, finding all that match a known VID/PID.
	vid, products := gousb.ID(0x046d), []SupportedDevice{
		{
			pid:           gousb.ID(0xc52b),
			description:   "Logitech Wireless test",
			supportsWheel: false,
			supportsLogo:  false,
		},
		{
			pid:           gousb.ID(0xc084),
			description:   "Logitech G203 Prodigy Gaming Mouse",
			url:           "https://www.logitechg.com/en-roeu/products/gaming-mice/g203-prodigy-gaming-mouse.html#910-004845",
			supportsWheel: false,
			supportsLogo:  true,
		},
		{
			pid:           gousb.ID(0xc083),
			description:   "Logitech G403 Prodigy Gaming Mouse",
			url:           "https://www.logitechg.com/en-us/products/gaming-mice/g403-prodigy-wired-gaming-mouse.html",
			supportsWheel: true,
			supportsLogo:  true,
		},
		{
			pid:           gousb.ID(0xc08f),
			description:   "Logitech G403 HERO Gaming Mouse",
			url:           "https://www.logitechg.com/en-roeu/products/gaming-mice/g403-hero-gaming-mouse.html",
			supportsWheel: true,
			supportsLogo:  true,
		},
	}

	//// Open our jsonFile
	//jsonFile, err := os.Open("devices.json")
	//// if we os.Open returns an error then handle it
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("Successfully Opened users.json")
	//// defer the closing of our jsonFile so that we can parse it later on
	//defer jsonFile.Close()

	foundling := SupportedDevice{}

	devices, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		// this function is called for every device present.
		// Returning true means the device should be opened.
		if desc.Vendor != vid {
			return false
		}
		//log.Printf("Found vendor \"%s\"", desc.Vendor)
		for _, product := range products {
			if desc.Product == product.pid {
				log.Printf("Found \"%s\"", product.description)
				foundling = product
				return true
			}
		}
		//log.Printf("Skipping unknown vendor product \"%s\"", desc.Product)
		return false
	})
	if err != nil {
		log.Print("ERROR: Possibly not run as root?")
		log.Fatalf("FATAL: %v", err)
	}
	// All returned devices are now open and will need to be closed.
	for _, d := range devices {
		defer d.Close()
	}
	if len(devices) == 0 {
		log.Fatal("no compatible devices found")
	}

	// Open any device with a given VID/PID using a convenience function.
	device, err := ctx.OpenDeviceWithVIDPID(vid, foundling.pid)
	if err != nil {
		log.Fatalf("Could not open a device: %v", err)
	}
	defer device.Close()

	//type + mode + color + misc + suffix
	data, err := hex.DecodeString("11ff0e3b00" + "01" + "000000" + "0000000000" + "000000000000")
	if err != nil {
		log.Fatalf("FATAL: %v", err)
	}
	res, err := device.Control(0x21, 0x09, 0x0211, 0x01, data)
	if err != nil {
		log.Fatalf("Could not control a device: %v", err)
	}
	if res != 0 {
		log.Fatalf("Could not control a device(2): %v", err)
	}

	err = device.Close()
	if err != nil {
		log.Fatalf("Could not close a device: %v", err)
	}
}
