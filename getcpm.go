/*
* Copyright (C) 2017 - Alexey V. Voronin @ FoxyLab
* Email:    support@foxylab.com
* Website:  https://acdc.foxylab.com
*
* This program is free software; you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation; either version 3 of the License, or
* (at your option) any later version.
*
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License
* along with this program; if not, write to the Free Software
* Foundation, Inc., 59 Temple Place, Suite 330, Boston, MA  02111-1307 USA
*/

package main

import (
    "io/ioutil"
    "fmt"
    "os"
    "strconv"
    "log"
    "encoding/hex"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    	var length_00 uint32
	var length_01 uint32
	var i uint32
	var b_00 uint8
	var b_01 uint8
	var count uint32
	var offset uint8
	var ram int

	ram_str := os.Args[1]
	ram, err := strconv.Atoi(ram_str)
	check(err)

	serial := os.Args[2]

	if (len(serial)!=12) {
	   log.Fatal("Serial number length is incorrect!");
	}

	serial_hex:=[]byte(serial)

	serial_bytes := make([]byte, hex.DecodedLen(len(serial_hex)))
	n, err := hex.Decode(serial_bytes, serial_hex)
	check(err)

	if (n!=6) {
	   log.Fatal("Serial number length is incorrect!");
	}

	data_00, err := ioutil.ReadFile("CPM00K.SYS")
	check(err)
	length_00 = uint32(len(data_00))	
	
	data_01, err := ioutil.ReadFile("CPMDIFF.SYS")
	check(err)
	length_01 = uint32(len(data_01))
	

	if length_00!=length_01 {
		fmt.Println("Different file lengths!")
	}else{
		
		offset = 52+(uint8(ram)-20)*4

		fmt.Print("Target RAM: ")
		fmt.Print(ram)
		fmt.Println(" KBytes")		

		count = 0
		
		tgt := make([]byte, length_00)
		//adresses changes
		for i = 0; i < length_00; i++ {
			b_00 = data_00[i]
			b_01 = data_01[i]
			if (b_01!=0) {
					tgt[i] = b_00 + offset
					count++
			}else{
				tgt[i] = b_00
			}
		}
		//serial number write

		//report
		tgtname := "CPM"
		tgtname = tgtname + ram_str
		tgtname = tgtname + "K.SYS"
		err = ioutil.WriteFile(tgtname, tgt, 664)
    		check(err)
		fmt.Print("Output file: ")
		fmt.Println(tgtname)
		fmt.Print(count)
		fmt.Println(" changes are made")
	}
}

