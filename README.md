# linear – linear least squares with Go

linear is a Go program to take in a csv file with data only (as numbers,) perform least squares linear regression using batch gradient descent, and return the parameter vector θ of the line of best fit. linear works for data in N dimensions.

#installation

```bash
# from any $GOPATH/src
go get github.com/cdipaolo/linear

# install binaries
go install ./...

# run linear regression:
linear -f='~/data.csv'

# example output
CSV File Path: /Users/cdipaolo/Data/LSD_Math/lsd.csvRecord: [1.17 78.93]Record: [2.97 58.20]Record: [3.26 67.47]Record: [4.69 37.47]Record: [5.83 45.65]Record: [6.00 32.92]Record: [6.41 29.97]Data: [[1.17] [2.97] [3.26] [4.69] [5.83] [6] [6.41]]
y: [78.93 58.2 67.47 37.47 45.65 32.92 29.97]

j(θ) = +Inf
Theta with <α = 10>: [-5.910366272001731e+155 -2.9796712511438646e+156]

j(θ) = +Inf
Theta with <α = 3.3333333333333335>: [8.295525651153176e+155 4.1821332483166224e+156]

j(θ) = +Inf
Theta with <α = 1.1111111111111112>: [1.5253031771464177e+154 7.689712984035258e+154]

j(θ) = +Inf
Theta with <α = 0.3703703703703704>: [-6.722575966189079e+153 -3.3891412846906586e+154]

j(θ) = +Inf
Theta with <α = 0.1234567901234568>: [-4.484024965425517e+152 -2.260590911629691e+153]

j(θ) = 18.134465785768114
Theta with <α = 0.0411522633744856>: [89.0897287015997 -9.002693525546308]

j(θ) = 18.134638299009588
Theta with <α = 0.013717421124828532>: [89.06447864846355 -8.997685017900878]

j(θ) = 18.135157462761608
Theta with <α = 0.004572473708276177>: [89.02074424297835 -8.989010022127218]

j(θ) = 18.13671294352282
Theta with <α = 0.0015241579027587256>: [88.94517869481557 -8.974021117998445]

j(θ) = 18.141381429497574
Theta with <α = 0.0005080526342529085>: [88.81427951693965 -8.948056438651562]

Final theta with j(θ) = 18.134465785768114
[89.0897287015997 -9.002693525546308]
```


# license – MIT

Copyright 2015 William Conner DiPaolo and other contributors

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.