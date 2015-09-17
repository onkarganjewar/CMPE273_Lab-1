package inter

import "testing"
import "math"

func TestPerimeter (t *testing.T) {

  k:= 0 // instance variables
  m:= 0

  var cd = []Circle {{2},{4},{5},{2},{1},{3}}   // predefined shape inputs
  var rc = []Rectangle {
                          {2,3},{1,2},{4,5},{7,8},{4,4},
                        }

type Rstr struct {
 num float64
  rc Rectangle
} // structure to define the output of Rectangle inputs


type Cstr struct {
  cnum float64
  cr  Circle
} // structure to define the output of Circle inputs


var Rarr = []Rstr {
{10, rc[0]}, {6, rc[1]}, {18, rc[2]}, {30, rc[3]}, {16, rc[4]},
} // array of structure consisting output values for predefined Rectangle inputs


var Carr = []Cstr {
{4*math.Pi, cd[0]}, {8*math.Pi, cd[1]}, {10*math.Pi, cd[2]}, {4*math.Pi, cd[3]}, {2*math.Pi, cd[4]}, {6*math.Pi, cd[5]},
}  // array of structure consisting output values for predefined Circle inputs


  for _, i := range cd {
    val := Shape.Perimeter(i);
    if val != Carr[k].cnum {
           t.Errorf("Input value = %f, Expected output = %f, Actual output = %f", i, Carr[k].cnum, val)
          }
          k++
        } //loop to test the Perimeter of circle


  for _, j := range rc {
    rval := Shape.Perimeter(j);
    if rval != Rarr[m].num {
      t.Errorf("Expected value %f Actual value %f Input Value %f",Rarr[m].num,rval,j)
    }
    m++
  } //loop to test the Perimeter of Rectangle
}
