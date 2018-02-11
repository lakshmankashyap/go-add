package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "encoding/json"
  "flag"
)

// Declare a Point type.
type Point struct {
  X float64
  Y float64
}

func gracefulExit (str string) {
  fmt.Println(str)
  os.Exit(1)
}

func main() {
  
  // Declare a var to hold the input JSON.
  var input string

  // Declare a var to hold the decoded JSON.
  var point = Point{}

  // Declare a "verbose" flag to be set by the -v flag.
  var verbose = flag.Bool("v", false, "Print verbose expression")  

  // Parse all command line flags.
  flag.Parse()
  
  //   ██████╗ ███████╗████████╗    ██╗███╗   ██╗██████╗ ██╗   ██╗████████╗
  //  ██╔════╝ ██╔════╝╚══██╔══╝    ██║████╗  ██║██╔══██╗██║   ██║╚══██╔══╝
  //  ██║  ███╗█████╗     ██║       ██║██╔██╗ ██║██████╔╝██║   ██║   ██║   
  //  ██║   ██║██╔══╝     ██║       ██║██║╚██╗██║██╔═══╝ ██║   ██║   ██║   
  //  ╚██████╔╝███████╗   ██║       ██║██║ ╚████║██║     ╚██████╔╝   ██║   
  //   ╚═════╝ ╚══════╝   ╚═╝       ╚═╝╚═╝  ╚═══╝╚═╝      ╚═════╝    ╚═╝   
  //                                                                       

  if len(os.Args) > 1 {
    // If a filename was provided, attempt to read from that file.
    if content, err := ioutil.ReadFile(os.Args[len(os.Args)-1]); err == nil {
      // If successful, assign content to "input".
      input = string(content)
    } else {
      // Otherwise, if the file couldn't be found, say so.
      if (os.IsNotExist(err)) {
        gracefulExit(fmt.Sprintf("Sorry, could not find the file `%v`.\n", os.Args[1]))
      }
      // If an unknown error occurred, bail out as gracefully as possible.
      gracefulExit(fmt.Sprintf("Sorry, an error occurred reading the input file:\n\n%v.\n", err))
      return
    }


  } else {
    if info, err := os.Stdin.Stat(); err == nil {
      // Otherwise check if there's any data in stdin.
      if info.Size() > 0 {
        // If so, attempt to read from that.
        if data, err := ioutil.ReadAll(os.Stdin); err == nil {
          // If successful, assign the data to "input".
          input = string(data)
        } else {
          // Otherwise bail out as gracefully as possible.
          gracefulExit(fmt.Sprintf("Sorry, an error occurred reading the input:\n\n%v\n", err))
          return
        }
      } else {
        gracefulExit("Usage: calc <input filename>")
      }
    } else {
      // If an unknown error occurred, bail out as gracefully as possible.
      gracefulExit(fmt.Sprintf("Sorry, an unknown error occurred:\n\n%v\n", err))

    }
  }
  
  //  ██████╗ ███████╗ ██████╗ ██████╗ ██████╗ ███████╗    ██╗███╗   ██╗██████╗ ██╗   ██╗████████╗
  //  ██╔══██╗██╔════╝██╔════╝██╔═══██╗██╔══██╗██╔════╝    ██║████╗  ██║██╔══██╗██║   ██║╚══██╔══╝
  //  ██║  ██║█████╗  ██║     ██║   ██║██║  ██║█████╗      ██║██╔██╗ ██║██████╔╝██║   ██║   ██║   
  //  ██║  ██║██╔══╝  ██║     ██║   ██║██║  ██║██╔══╝      ██║██║╚██╗██║██╔═══╝ ██║   ██║   ██║   
  //  ██████╔╝███████╗╚██████╗╚██████╔╝██████╔╝███████╗    ██║██║ ╚████║██║     ╚██████╔╝   ██║   
  //  ╚═════╝ ╚══════╝ ╚═════╝ ╚═════╝ ╚═════╝ ╚══════╝    ╚═╝╚═╝  ╚═══╝╚═╝      ╚═════╝    ╚═╝   
  //                                                                                                

  // Attempt to parse the JSON.
  err := json.Unmarshal([]byte(input), &point)

  // If unsuccessful, bail out as gracefully as possible.
  if err != nil {
    gracefulExit(fmt.Sprintf(`Sorry, an error occurred processing the input.
Please ensure input is a JSON string representing an object with 'x' and 'y' keys, whose values are numbers.
For the record, your input was %v`, input))
  }

  if *verbose == true {
    // If verbose mode is true, output the whole expression.
    fmt.Printf("%v + %v = %v", point.X, point.Y, point.X + point.Y);  
  } else {
    // Otherwise output the result only.
    fmt.Printf("%v", point.X + point.Y);
  }

}
