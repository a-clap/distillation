/*
 * Copyright (c) 2023 a-clap. All rights reserved.
 * Use of this source code is governed by a MIT-style license that can be found in the LICENSE file.
 */

package main

import (
	"log"
	"time"

	"github.com/a-clap/iot/internal/distillation"
	"github.com/a-clap/iot/internal/embedded"
)

const EmbeddedAddr = "http://localhost:8080"

func main() {
	distil, err := distillation.New(
		distillation.WithPT(embedded.NewPTClient(EmbeddedAddr, 1*time.Second)),
		distillation.WithDS(embedded.NewDS18B20Client(EmbeddedAddr, 1*time.Second)),
		distillation.WithHeaters(embedded.NewHeaterClient(EmbeddedAddr, 1*time.Second)),
		distillation.WithGPIO(embedded.NewGPIOClient(EmbeddedAddr, 1*time.Second)),
	)
	if err != nil {
		log.Fatalln(err)
	}
	pts := distil.PTHandler.GetSensors()
	for _, pt := range pts {
		pt.Enabled = true
		pt.Samples = 10
		if _, err := distil.PTHandler.Configure(pt); err != nil {
			log.Println(err)
		}
	}

	err = distil.Run("localhost:8081")
	log.Println(err)
}