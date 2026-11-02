
package structs

import(
"image"
_"image/png"
_"image/jpeg"
_"image/gif"
"math"
"golang.org/x/image/draw"
)


type flood struct{
	a *Arg
}

func (f  *flood) imgtopx(){
	w:= int(math.Round(float64(f.a.XLen) * f.a.Scale)) 
	h:= int(math.Round(float64(f.a.YLen) * f.a.Scale))
 	dst := image.NewRGBA(image.Rect(f.a.X, f.a.Y, w, h))
	draw.NearestNeighbor.Scale(dst, dst.Rect, f.a.Image, f.a.Image.Bounds(), draw.Over, nil)
}
