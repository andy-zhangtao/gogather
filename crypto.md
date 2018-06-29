# crypto
--
    import "github.com/andy-zhangtao/gogather/crypto"


## Usage

#### func  GeneratePasswd

```go
func GeneratePasswd(password, userSalt string, iterations int) (cryptoPasswd string)
```
Write by zhangtao<ztao8607@gmail.com> . In 2018/5/7. GeneratePasswd
生成与Django兼容的加密字符串

#### func  GeneratePrivateKey

```go
func GeneratePrivateKey(bitSize int) (*rsa.PrivateKey, []byte, error)
```
GeneratePrivateKey 生成私钥 返回原生私钥数据， 和文本化后的私钥数据

#### func  GeneratePublicKey

```go
func GeneratePublicKey(privatekey *rsa.PublicKey) ([]byte, error)
```
GeneratePublicKey 基于私钥生成相对应的公钥
