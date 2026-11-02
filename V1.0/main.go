package main

import(
"fmt"
"os"
"flag"
"image"
_ "image/png"
_ "image/jpeg"
_ "image/gif"
"net"
"bufio"
"math"
"golang.org/x/image/draw"
"github.com/reiver/go-telnet"
)

type floodtyp struct{
	pat string
	d *details
}

func (f floodtyp) ltr(){
	w:= int(math.Round(float64(f.d.x_len) * f.d.scale)) 
	h:= int(math.Round(float64(f.d.y_len) * f.d.scale))
 	dst := image.NewRGBA(image.Rect(f.d.X, f.d.Y, w, h))
	draw.NearestNeighbor.Scale(dst, dst.Rect, f.d.image, f.d.image.Bounds(), draw.Over, nil)
	con , err := telnet.DialTo(f.d.adress)
	if err != nil{
		fmt.Println("eooro:", err)
		os.Exit(1)
	}
	out :=bufio.NewWriter(con)
	defer out.Flush()
   for y := range h {
       for x := range w {
           c := dst.RGBAAt(x, y)
           fmt.Fprintf(out, "PX %d %d %02x%02x%02x\n", f.d.X+x, f.d.Y+y, c.R, c.G, c.B)
       }
   }
}

type details struct{
	adress string
	X int
	Y int
	x_len int
	y_len int
	scale float64
	path string
	image image.Image
	pre bool
	paterm floodtyp
}

func (d details) preshow(){
	w := int(math.Round(float64(d.x_len) * d.scale))
	h := int(math.Round(float64(d.y_len) * d.scale))
dst := image.NewRGBA(image.Rect(0, 0, w, h))
draw.NearestNeighbor.Scale(dst, dst.Rect, d.image, d.image.Bounds(), draw.Over, nil)


	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

for y:= range h {
	for x := range w {
					top := dst.RGBAAt(x, y)
			bot := dst.RGBAAt(x, y+1)
			fmt.Fprintf(out, "\x1b[38;2;%d;%d;%dm\x1b[48;2;%d;%d;%dm▀",
				top.R, top.G, top.B, bot.R, bot.G, bot.B)
		}
		fmt.Fprint(out, "\x1b[0m\n")
	}
}



func (d *details)  imgstart(){
	file, err := os.Open(d.path)
	if nil != err {
		fmt.Println("not a valide image")
		os.Exit(1)
	}
	img, name, err := image.Decode(file)
	if nil != err {
		fmt.Println("somthing is WRONG with you iamge")
		os.Exit(1)
	}
	d.image = img
	defer file.Close()
	fmt.Println(name, "is loded")

	b := d.image.Bounds()
	if d.x_len == -1 {
		d.x_len = b.Dx()
	}
	if d.y_len == -1 {
		d.y_len = b.Dy()
	}
}

func (d *details) argu(){
	flag.StringVar(&d.adress, "ip", "", "the adress of the pixel flut server (with port)")
	flag.IntVar(&d.X, "X", 0, "X Cord")
	flag.IntVar(&d.Y, "Y", 0, "Y Cord")
	flag.IntVar(&d.x_len, "xlen", -1, "the length of the X side")
	flag.IntVar(&d.y_len, "ylen", -1, "the length of the y side")
	flag.Float64Var(&d.scale, "s", 1, "the scaling factori of the image")
	flag.StringVar(&d.path, "path", "", "the path to the image")
	flag.BoolVar(&d.pre, "pre", false, "show the previe in terminal")
	flag.Parse() 

        conn, err := net.Dial("tcp", d.adress)
        if err != nil {
                 fmt.Println("Connection error:", err)
								os.Exit(1)
        } 
				conn.Close()

	if d.path == "" {
		fmt.Fprintln(os.Stderr, "error: no path given use -Path [Path] to set Path")
		os.Exit(1)
	}
	if _, err := os.Stat(d.path); os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr, "error: path does not exist:", d.path)
		os.Exit(1)
	}
}


func main() {

	d := details{X: 0, Y: 0, x_len: 100, y_len: 100, path: "."}
	d.argu()
	d.imgstart()
	d.paterm.d = &d
	if d.pre == true {
		d.preshow()

	}
	for i := range 100{
			i--
	d.paterm.ltr()
}
}
