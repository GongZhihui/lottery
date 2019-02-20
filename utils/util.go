package utils

import(
	"net/http"
	"encoding/json"
    "bytes"
    "crypto/md5"
    "encoding/hex"
    "math/big"
    "crypto/rand"
    //"fmt"
)

//ReadJSON 将json序列化为对象
func ReadJSON(msg []byte, obj interface{})(interface{}){
    index := bytes.IndexByte(msg, 0);
    var s []byte;
    if(index == -1){
        s = msg[:];
    }else{
        s =msg[0:index];  
    }
    json.Unmarshal(s[:], obj);
    return obj;
}

// WriteJSON 将一个对象转为json，并写到response
func WriteJSON(w http.ResponseWriter, obj interface{}){
    var jv []byte;
    var err error;
    jv, err = json.Marshal(obj);
    if (err !=nil){
        jv, _ = json.Marshal("internal error!!!");
    }
	w.Header().Set("Content-Type", "application/json; charset=utf-8");
	w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
    // w.WriteHeader(1);
    w.Write(jv);
}

func checkAccount(acc, pwd string)bool{
    var data = []byte(acc + pwd);
    var h = md5.New()
    h.Write(data)
    var s = hex.EncodeToString(h.Sum(nil))
    // gzh123 admin123
    if(s == "0192023a7bbd73250516f069df18b500" ||
        s == "e0d211e5f90f76c1222f6456c11150fd"){
        return true;
    }else{
        return false;
    }
}

// Percent . 0.01 -> 100  0.0001 -> 1000
func Percent(n float64)int{
    if(n <= 0){
        return 0;
    } 
    var ret = 0.0;
    for i := 1; ;i = i * 10{
        ret = n * float64(i);
        if(ret >= 1){
            return i;
        }
    }
}

// MakeRandom . [0 , 10)
func MakeRandom(min int64, max int64) int64 {
    maxBigInt:=big.NewInt(max)
    i,_:=rand.Int(rand.Reader,maxBigInt)
    if i.Int64()<min{
        MakeRandom(min,max);
    }
    return i.Int64();
}
