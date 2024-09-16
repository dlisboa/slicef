# slicef

## Usage

```
slicef slices a file from a begin line to an end line.

usage:
  slicef [OPTIONS] [FILE ...]

example:
  slicef -b 'func main' -e '}' -i -p main.go
  main.go
  func main() {
    // code in func main
  }

options:
  -b delim
	  prefix of delimiter of the start of the block. The first line matching
	  this prefix will mark the start and will be printed.

  -e delim
	  prefix of delimiter of the end of the block. The first line matching this
	  prefix will mark the end of the block and will NOT be printed unless -i
	  flag is set. If special delim ':EOF:' is given it'll print the entire file
	  starting from the beginning delimiter.

  -i (default: false)
      prints the end delimiter as well

  -p (default: false)
      print the filename

```


## LICENSE

BSD 2-Clause License

Copyright (c) 2024, Diogo Lisboa

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
