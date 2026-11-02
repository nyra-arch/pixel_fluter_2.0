package structs

import(
	"flag"
	"fmt"
	"image"
	"net"
	"os"
)

type Arg struct {
	Adress string
	X, Y   int
	XLen   int
	YLen   int
	Scale  float64
	Path   string
	Image  image.Image
	Pre    bool
}

func (a *Arg) Su(){
	flag.StringVar(&a.Adress, "host", "", "the adress of the host")
	flag.StringVar(&a.Path, "path", "", "the path to the image/gif")
	flag.IntVar(&a.X, "x", 0, "the x cordinat of the top left corner of the image/gif")
	flag.IntVar(&a.Y, "y", 0, "the y cordinat of the top left corner if the image/gif")
	flag.IntVar(&a.XLen, "xlen", -1, "the length of the x axis of the image/gif")
	flag.IntVar(&a.YLen, "ylen", -1, "the length of the y axis of the image/gif")
	flag.Float64Var(&a.Scale, "s", 1, "the scale factor of the image/gif")
	flag.BoolVar(&a.Pre, "pre", false, "show previe image")
	flag.Parse()

	        conn, err := net.Dial("tcp", a.Adress)
        if err != nil {
                 fmt.Println("Connection error:", err)
								os.Exit(1)
        } 
				conn.Close()

	if a.Path == "" {
		fmt.Fprintln(os.Stderr, "error: no path given use -Path [Path] to set Path")
		os.Exit(1)
	}
	if _, err := os.Stat(a.Path); os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr, "error: path does not exist:", a.Path)
		os.Exit(1)
	}

		file, err := os.Open(a.Path)
	if nil != err {
		fmt.Println("not a valide image")
		os.Exit(1)
	}
	img, name, err := image.Decode(file)
	if nil != err {
		fmt.Println("somthing is WRONG with you iamge")
		os.Exit(1)
	}
	a.Image = img
	defer file.Close()
	fmt.Println(name, "is loded")

	b := a.Image.Bounds()
	if a.XLen == -1 {
		a.XLen = b.Dx()
	}
	if a.YLen == -1 {
		a.YLen = b.Dy()
	}

}
