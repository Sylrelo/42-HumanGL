package mat4

import (
	"fmt"
	"math"
)

type Mat4 [4][4]float32

func Cos(f float32) float32 {
    return float32(math.Cos(float64(f)))
}

func Sin(f float32) float32 {
    return float32(math.Sin(float64(f)))
}

func Tan(f float32) float32 {
    return float32(math.Tan(float64(f)))
}

func Identity() Mat4 {
    var mat Mat4

    mat[0][0] = 1
    mat[0][1] = 0
    mat[0][2] = 0
    mat[0][3] = 0
    mat[1][0] = 0
    mat[1][1] = 1
    mat[1][2] = 0
    mat[1][3] = 0
    mat[2][0] = 0
    mat[2][1] = 0
    mat[2][2] = 1
    mat[2][3] = 0
    mat[3][0] = 0
    mat[3][1] = 0
    mat[3][2] = 0
    mat[3][3] = 1

    return mat
}

func Mult(a, b Mat4) Mat4 {
    var res Mat4 
    
    for i := 0; i < 4; i++ { 
        for j := 0; j < 4; j++ { 
            res[i][j] = 0 
            for k := 0; k < 4; k++ {
                res[i][j] += a[i][k] * b[k][j]
            }
        } 
    }

    return res
}

func (curr Mat4) Mult(other Mat4) Mat4 {    
    var res Mat4 
    
    for i := 0; i < 4; i++ { 
        for j := 0; j < 4; j++ { 
            res[i][j] = 0 
            for k := 0; k < 4; k++ {
                res[i][j] += other[i][k] * curr[k][j]
            }
        } 
    }

    return res
}

func Translate(x, y, z float32) Mat4 {
    var translation Mat4 = Identity()

    translation[3][0] = x
    translation[3][1] = y
    translation[3][2] = z

    return translation
}

func Scale(x, y, z float32) Mat4 {
    var scale Mat4 = Identity()

    scale[0][0] = x
    scale[1][1] = y
    scale[2][2] = z

    return scale
}

func Rotation(x, y, z float32) Mat4 {
    var rotationX Mat4 = Identity()
    var rotationY Mat4 = Identity()
    var rotationZ Mat4 = Identity()

    rotationX[1][1] = Cos(x)
    rotationX[1][2] = -Sin(x)
    rotationX[2][1] = Sin(x)
    rotationX[2][2] = Cos(x)

    rotationY[0][0] = Cos(y)
    rotationY[0][2] = Sin(y)
    rotationY[2][0] = -Sin(y)
    rotationY[2][2] = Cos(y)

    rotationZ[0][0] = Cos(z)
    rotationZ[0][1] = -Sin(z)
    rotationZ[1][0] = Sin(z)
    rotationZ[1][1] = Cos(z)

    return Mult(Mult(rotationX, rotationY), rotationZ)
}

func Perspective(fovradian, ratio, near, far float32) Mat4 {
    var perspective Mat4 = Identity()
    var e float32 = Tan(fovradian / 2.0)

    perspective[0][0] = 1.0 / (ratio * e)
    perspective[1][1] = 1.0 / (e)
    perspective[2][2] = (far + near) / (far - near) * -1
    perspective[2][3] = -1.0
    perspective[3][2] = (2.0 * far * near) / (far - near) * -1

    return perspective
}

func (mat Mat4) Print() {
    for i := 0; i < 4; i++ {
        fmt.Printf("%1.2f %1.2f %1.2f %1.2f\n", mat[i][0], mat[i][1], mat[i][2], mat[i][3])
    } 
}

func DegToRad(angle float32) float32 {
    return angle / 180.0 * math.Pi
}