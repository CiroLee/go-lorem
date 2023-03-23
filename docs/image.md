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