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

    	var length_00 uint16
	var length_01 uint16
	var i uint16
	var b_00 uint8
	var b_01 uint8
	var count uint16
	var offset uint8
	var ram int
	const SERIAL_1 uint16 = 808
	const SERIAL_2 uint16 = 2048
	const SERIAL_LEN uint16 = 16
	const TEMPLATE_FILE string = "CPM00K.SYS"
	const DIFF_FILE string = "CPMDIFF.SYS"

	//RAM size (KBytes) reading
	ram_str := os.Args[1]
	ram, err := strconv.Atoi(ram_str)
	check(err)

	//serial number reading
	serial := os.Args[2]
	if (uint16(len(serial))!=(SERIAL_LEN*2)) {
	   log.Fatal("Serial number length is incorrect!")
	}
	serial_hex:=[]byte(serial)
	serial_bytes := make([]byte, hex.DecodedLen(len(serial_hex)))
	n, err := hex.Decode(serial_bytes, serial_hex)
	check(err)
	if (uint16(n)!=SERIAL_LEN) {
	   log.Fatal("Serial number length is incorrect!")
	}

	//CPM.SYS template reading
	data_00, err := ioutil.ReadFile(TEMPLATE_FILE)
	check(err)
	length_00 = uint16(len(data_00))	
	//diff file reading
	data_01, err := ioutil.ReadFile(DIFF_FILE)
	check(err)
	length_01 = uint16(len(data_01))
	//files lengths compare
	if length_00!=length_01 {
		fmt.Println("Different file lengths!")
	}else{
		//offset calculation
		offset = 52+(uint8(ram)-20)*4
		fmt.Print("Target RAM: ")
		fmt.Print(ram)
		fmt.Println(" KBytes")		
		count = 0 //reset changes counter
		tgt := make([]byte, length_00) //target bytes array init
		//adresses changes
		for i = 0; i < length_00; i++ {
			b_00 = data_00[i]
			b_01 = data_01[i]
			if (b_01!=0) {
					//changes made
					tgt[i] = b_00 + offset
					count++
			}else{
				//no changes
				tgt[i] = b_00
			}
		}
		//serial number write
		//serial 1	
		for i = 0; i < SERIAL_LEN; i++ {
			tgt[i+SERIAL_1] = serial_bytes[i]
		}
		//serial 2
		for i = 0; i < SERIAL_LEN; i++ {
			tgt[i+SERIAL_2] = serial_bytes[i]
		}
		//target file name
		tgtname := "CPM"
		tgtname = tgtname + ram_str
		tgtname = tgtname + "K.SYS"
		//target CPM*.SYS write
		err = ioutil.WriteFile(tgtname, tgt, 664)
    		check(err)
		//report
		fmt.Print("Output file: ")
		fmt.Println(tgtname)
		fmt.Print(count)
		fmt.Println(" changes are made")
	}
}

