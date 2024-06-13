# ZeroWidth

텍스트에 폭이 없는 문자를 찾고 제거합니다.

## 설치

```bash
go get github.com/i31pc17/zerowidth
```

## 사용 방법
sample/sample.go 파일 참고
```go
import (
    "fmt"
    "github.com/i31pc17/zerowidth"
)

func main() {
    zw := zerowidth.NewZeroWidth()
    
    zeroText := "(​) 폭이 없는 문자 제거 테스트 입니다."
    zeroRemove, err := zw.Remove(zeroText)
    
    fmt.Println("제거전 : ", zeroText)
    if err == nil {
        fmt.Println("제거후 : ", zeroRemove)
    }
    
    spaceText := "( ) 공백 제거 테스트 입니다."
    spaceRemove, err := zw.RemoveSpace(spaceText)
    
    fmt.Println("제거전 : ", spaceText)
    if err == nil {
        fmt.Println("제거후 : ", spaceRemove)
    }

    vaText := "(❤️) 변형 셀렉터 제거 테스트 입니다."
    vaRemove, err := zw.RemoveVarSelector(vaText)

    fmt.Println("제거전 : ", vaText)
    if err == nil {
        fmt.Println("제거후 : ", vaRemove)
    }
}
```
```
제거전 :  (​) 폭이 없는 문자 제거 테스트 입니다.
제거후 :  () 폭이 없는 문자 제거 테스트 입니다.
제거전 :  ( ) 공백 제거 테스트 입니다.
제거후 :  ()공백제거테스트입니다.
제거전 :  (❤️) 변형 셀렉터 제거 테스트 입니다.
제거후 :  (❤) 변형 셀렉터 제거 테스트 입니다.
```