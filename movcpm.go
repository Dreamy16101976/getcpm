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
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    	var length_00 uint32
	var length_01 uint32
	const buf_size = 256
	var i uint32
	var b_00 uint8
	var b_01 uint8
	var count uint32

	filename_00 := os.Args[1]
	fmt.Printf("File: %s\n", filename_00)		
	data_00, err := ioutil.ReadFile(filename_00)
	check(err)
	length_00 = uint32(len(data_00))	

	filename_01 := os.Args[2]
	fmt.Printf("File: %s\n", filename_01)		
	data_01, err := ioutil.ReadFile(filename_01)
	check(err)
	length_01 = uint32(len(data_01))	
	
	if length_00!=length_01 {
		fmt.Println("Different file lengths!")
	}else{
		count = 0;
		
		diff := make([]byte, length_00)

		for i = 0; i < length_00; i++ {
			b_00 = data_00[i]
			b_01 = data_01[i]
			if (b_00 != b_01) {
				//upcase/lower case
				//2 serial numbers
				//starting cursor position
				if ((b_01-b_00)!=32) && ((b_01-b_00)!=224) && ((i<808) || (i>813)) && ((i<2048) || (i>2053)) && (i!=2827) {
					diff[i] = b_01 - b_00
					count++
				}
			}else{
				diff[i] = 0;
			}
		}
		err = ioutil.WriteFile("CPMDIFF.SYS", diff, 664)
    		check(err)
	}
	fmt.Print(count)
	fmt.Println(" differences")
}

