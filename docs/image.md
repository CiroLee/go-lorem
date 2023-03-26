## Placeholder       
return a random placeholder image with pure color background. if only set Width or Height, the other value will
be the same, that means it will render a square image.  if you do not set Width nor Height, the values will between 320 and 1024.                    
signature:      
```go
func Placeholder(option PlaceholderOption) (string, error)      

type PlaceholderOption struct {
	Width     int      // width of the image. default between 320 and 1024
	Height    int      // height of the image. default between 320 and 1024
	Text      string   // text shown in the image
	BgColor   string   // background of the image. default is a random hex color
	FontColor string   // font color of text, only valid if text set and value will be #fffff, if bgColor is dark, or #000000 if bgColor is light
}
```
example:     
```go
lorem.Placeholder(lorem.PlaceholderOption{})  // https://dummyimage.com/563x688/ffffff
lorem.Placeholder(lorem.PlaceholderOption{Width: 100}) // https://dummyimage.com/100x100/31b6c6
lorem.Placeholder(lorem.PlaceholderOption{Text: "hahaha"}) // https://dummyimage.com/449x772/67320c/ffffff&text=hahaha
```

## SimplePlaceholder        
return a random placeholder image without params        
signature:     
```go
func SimplePlaceholder() string
```
example:     
```go
lorem.SimplePlaceholder() // https://dummyimage.com/563x688/ffffff&text=image 
```

## Picsum      
return a random image from unSplash by special param struct        
signature:      
```go
func Picsum(option PicsumOption) string

type PicsumOption struct {
	Width     uint  // width of the image. default between 320 and 1024
	Height    uint  // Height of the image. default between 320 and 1024
	Grayscale bool  // is grayscale image
	Blur      uint  // blur value. min is 1 and max is 10
}
```
example:      
```go
lorem.Picsum(lorem.PicsumOption{Grayscale: true, Width: 100, Height: 200, Blur: 2})
// https://picsum.photos/100/200?grayscale&blur=2
lorem.Picsum(lorem.PicsumOption{Width: 100})
// https://picsum.photos/100/100
```

## SimplePicsum       
simple use of Picsum. return a random image from unSplash without params            
signature:      
```go
func SimplePicsum() string
```
example:     
```go
lorem.SimplePicsum() // https://picsum.photos/581/797
```

## Classify       
return a random image classified image       
signature:      
```go
func Classify(option ClassifyOption) string

type ClassifyOption struct {
	Type   string    // image type. you can use built-in type consts, such as consts.ImageAnimals
	Width  uint      // width of the image. default between 320 and 1024
	Height uint      // Height of the image. default between 320 and 1024
	Lock   bool      // set Lock true, the server will return the cache image
}
```
example:      
```go
import (
	"github.com/CiroLee/go-lorem/consts"
)

lorem.Classify(lorem.ClassifyOption{ Type: consts.ImageAnimals })
// https://loremflickr.com/463/611/animals
```

## SimpleClassify      
simple use of Classify without params          
signature:      
```go
func SimpleClassify() string
```
example:     
```go
lorem.SimpleClassify() // https://loremflickr.com/463/611/animals
```