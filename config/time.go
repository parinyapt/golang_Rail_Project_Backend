package config

import (
	"log"
	"os"
	"time"
)

func TimezoneSetup() {
	loc, err := time.LoadLocation(os.Getenv("TZ")) 
	if err != nil {
		log.Fatalf("[Error]->Failed to LoadLocation Timezone Config on TimezoneSetup() : %s", err)
	}
	time.Local = loc // -> this is setting the global timezone
}