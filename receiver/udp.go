/*
 * skogul, udp message receiver
 *
 * Copyright (c) 2019 Telenor Norge AS
 * Author(s):
 *  - Kristian Lyngstøl <kly@kly.no>
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with this library; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA
 * 02110-1301  USA
 */

package receiver

import (
	"log"
	"net"

	"github.com/KristianLyng/skogul"
)

// UDP contains the configuration for the receiver
type UDP struct {
	Address string            `doc:"Address and port to listen to." example:"[::1]:3306"`
	Handler skogul.HandlerRef `doc:"Handler used to parse, transform and send data."`
}

// Start starts listening for incoming UDP messages on the configured address
func (ud *UDP) Start() error {
	udpip, err := net.ResolveUDPAddr("udp", ud.Address)
	if err != nil {
		log.Printf("Can't resolve %s: %v", ud.Address, err)
		return err
	}
	ln, err := net.ListenUDP("udp", udpip)
	if err != nil {
		log.Printf("Can't listen on %s: %v", ud.Address, err)
		return err
	}
	for {
		bytes := make([]byte, 9000)
		n, err := ln.Read(bytes)
		if err != nil || n == 0 {
			log.Printf("Unable to read UDP message. Got %d bytes. Error: %v", n, err)
			continue
		}
		go func() {
			if err := ud.Handler.H.Handle(bytes[0:n]); err != nil {
				log.Printf("Unable to handle UDP message: %s", err)
			}
		}()
	}
}
