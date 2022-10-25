package fuzz

import "strconv"
import "github.com/awnumar/rosen/config"
import "github.com/awnumar/rosen/crypto"

func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) > 1 {
        num, _ = strconv.Atoi(string(bytes[0]))

        switch num {
    
        case 0:
            content := string(bytes)
            config.DecodeKeyString(content)
            return 0

        case 1:
            content := string(bytes)
            config.LoadConfig(content)
            return 0

        default:
            crypto.NewCipher(bytes)
            return 0

        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}