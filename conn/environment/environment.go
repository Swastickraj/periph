// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package environment

import (
	"time"

	"periph.io/x/periph/conn"
	"periph.io/x/periph/conn/physic"
)

// Weather represents measurements from an environmental sensor.
type Weather struct {
	Temperature physic.Temperature
	Pressure    physic.Pressure
	Humidity    physic.RelativeHumidity
}

// SenseWeather represents an environmental weather sensor.
type SenseWeather interface {
	conn.Resource

	// Sense returns the value read from the sensor. Unsupported metrics are not
	// modified.
	Sense(w *Weather) error
	// SenseContinuous initiates a continuous sensing at the specified interval.
	//
	// It is important to call Halt() once done with the sensing, which will turn
	// the device off and will close the channel.
	SenseContinuous(interval time.Duration) (<-chan Weather, error)
	// Precision returns this sensor's precision.
	//
	// The env values are set to the number of bits that are significant for each
	// items that this sensor can measure.
	//
	// Precision is not accuracy. The sensor may have absolute and relative
	// errors in its measurement, that are likely well above the reported
	// precision. Accuracy may be improved on some sensor by using oversampling,
	// or doing oversampling in software. Refer to its datasheet if available.
	Precision(w *Weather)
}
