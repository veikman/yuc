package main

// This file constitutes the source code of YUC.
// 
// YUC is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
// 
// YUC is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
// 
// You should have received a copy of the GNU General Public License
// along with YUC.  If not, see <http://www.gnu.org/licenses/>.
// 
// Copyright 2015 Viktor Eikman

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

var Version = "0.1.0"

func print_usage() {
	fmt.Printf("YAGNI Universal Compiler (YUC) v%s\n", Version)
	fmt.Println("usage:")
	fmt.Printf("  %s [-h] FILE [FILE...]\n", os.Args[0])
	fmt.Println("arguments:")
	fmt.Println("  -h    display this help message")
	fmt.Println("  FILE  source code filepath, any language")
}

func compile(filepaths []string) error {
	for _, path := range filepaths {
		file, err := os.Open(path)
		if os.IsNotExist(err) {
			s := "no such file or directory: %s"
			return errors.New(fmt.Sprintf(s, path))
		}
		fmt.Printf("skipping (not needed): %s\n", path)
		file.Close()
	}
	fmt.Println("finished without errors")
	return nil
}

func main() {
	var help = flag.Bool("h", false, "print help message")
	flag.Parse()

	if *help {
		print_usage()
		return
	}

	filepaths := os.Args[1:]
	if len(filepaths) == 0 {
		print_usage()
		os.Exit(1)
	}

	err := compile(filepaths)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
