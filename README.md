# CanvasGio arcLine bug

|| NO LONGER RELEVANT ||

|| BUG CORRECTED IN A LATTER VERSION OF GIOCANVAS ||


## Notes

some example code and thoughts about the ArcLine method of golang giocanvas package

issue closed: https://github.com/ajstarks/giocanvas/issues/12

## Examples

- `go run .\cmd\arcLine01\`
    - Original arcLine method has a memory leak when angle1 and angle2 have the same value or is their difference is too small.
    - Needs to be studied what "too small" means

- `go run .\cmd\arcLine02\`
    - show the anticlockwise issues that limit the 2 angles of arcline


- `go run .\cmd\arcLine03\`
    - shows a possible solution to all those issues by rewriting the arcLine method
    - it also uses a animation to test the new arcLine method

- `go run .\cmd\arcLine04\`
    - Not relevant for the issue here.
    - similar to the arcLine03, but with more fluff.
    - Just some quick teaking with the animation and stuff.

- `go run .\cmd\arcLine05\`
    - Just a fun visualization made by the author of giocanvas in feedback to the issue.
    - copied here for archive purposes.

##  links 

https://github.com/ajstarks/giocanvas

https://pkg.go.dev/github.com/ajstarks/giocanvas

