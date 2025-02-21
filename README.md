## 上传文件到minio
1. 启动minio, 创建对应的`accessKey` 和 `secretKey`和 `bucket`, 本地启动的 minio 一般没有证书, 而且可能是HTTP协议, 需要跳过证书检验, `secure` 为 `true` 时, 会跳过证书检验
```shell
docker compose -f minio/compose.yaml up -d
```

2. 配置 Minio配置
```yaml
data:
  minio:
    endpoint: 192.168.3.125:9000
    accessKey: LH1I4a83WdghqEmTnZmw
    secretKey: 3xRjjBTJ3ZAcOIvgcDAmkOZQvY8NYAIefNSFTpkK
    secure: true
    token:

```

3. 启动后端服务:
```bash
kratos run
```

4. 请求预签名URL

![img.png](img.png)

根据实际情况来修改:
- method: POST 和 PUT 都行, 建议使用 PUT,它是冥等的
- contentType: 上传文件的类型
- bucketName: 桶名
- filePath: 上传文件的路径
- fileName: 上传文件的名称
```json
{
  "method": "PUT",
  "contentType": "video/mp4",
  "bucketName": "tiktok",
  "filePath":"/Users/sumery/Downloads/example_video.mp4",
  "fileName": "example.mp4"
}
```

- curl:
```bnsh
curl --location --request POST 'http://localhost:8000/admin/v1/file:upload-url' \
--header 'Content-Type: application/json' \
--data-raw '{
  "method": "POST",
  "contentType": "video/mp4",
  "bucketName": "tiktok",
  "filePath":"/Users/sumery/Downloads/example_video.mp4",
  "fileName": "example.mp4"
}'

```

- powershell:
```powershell
$headers = New-Object "System.Collections.Generic.Dictionary[[String],[String]]"
$headers.Add("Content-Type", "application/json")

$body = "{`n  `"method`": `"POST`",`n  `"contentType`": `"video/mp4`",`n  `"bucketName`": `"tiktok`",`n  `"filePath`":`"/Users/sumery/Downloads/example_video.mp4`",`n  `"fileName`": `"example.mp4`"`n}"

$response = Invoke-RestMethod 'http://localhost:8000/admin/v1/file:upload-url' -Method 'POST' -Headers $headers -Body $body
$response | ConvertTo-Json
```

- js fetch:
```ts
var myHeaders = new Headers();
myHeaders.append("Content-Type", "application/json");

var raw = JSON.stringify({
   "method": "POST",
   "contentType": "video/mp4",
   "bucketName": "tiktok",
   "filePath": "/Users/sumery/Downloads/example_video.mp4",
   "fileName": "example.mp4"
});

var requestOptions = {
   method: 'POST',
   headers: myHeaders,
   body: raw,
   redirect: 'follow'
};

fetch("http://localhost:8000/admin/v1/file:upload-url", requestOptions)
   .then(response => response.text())
   .then(result => console.log(result))
   .catch(error => console.log('error', error));
```

go:
```go
package main

import (
   "fmt"
   "strings"
   "net/http"
   "io/ioutil"
)

func main() {

   url := "http://localhost:8000/admin/v1/file:upload-url"
   method := "POST"

   payload := strings.NewReader(`{
  "method": "POST",
  "contentType": "video/mp4",
  "bucketName": "tiktok",
  "filePath":"/Users/sumery/Downloads/example_video.mp4",
  "fileName": "example.mp4"
}`)

   client := &http.Client {
   }
   req, err := http.NewRequest(method, url, payload)

   if err != nil {
      fmt.Println(err)
      return
   }
   req.Header.Add("Content-Type", "application/json")

   res, err := client.Do(req)
   if err != nil {
      fmt.Println(err)
      return
   }
   defer res.Body.Close()

   body, err := ioutil.ReadAll(res.Body)
   if err != nil {
      fmt.Println(err)
      return
   }
   fmt.Println(string(body))
}
```

5. 上传文件到minio
将第四步生成的`downloadUrl` 作为 `URL PATH`, 请求方法为 `PUT`
![img_1.png](img_1.png)

```bash
var https = require('follow-redirects').https;
var fs = require('fs');

var options = {
   'method': 'PUT',
   'hostname': '192.168.3.125',
   'port': 9000,
   'path': '/tiktok/example.mp4?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=LH1I4a83WdghqEmTnZmw%252F20250221%252Fus-east-1%252Fs3%252Faws4_request&X-Amz-Date=20250221T172240Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&X-Amz-Signature=7b2a63f7e139c3d4a6756fa72771f8f5c230565d826f0121f2d39263e79613ca',
   'headers': {
      'Content-Type': 'application/octet-stream'
   },
   'maxRedirects': 20
};

var req = https.request(options, function (res) {
   var chunks = [];

   res.on("data", function (chunk) {
      chunks.push(chunk);
   });

   res.on("end", function (chunk) {
      var body = Buffer.concat(chunks);
      console.log(body.toString());
   });

   res.on("error", function (error) {
      console.error(error);
   });
});

req.end();
```

- powershell:
```powershell
$headers = New-Object "System.Collections.Generic.Dictionary[[String],[String]]"
$headers.Add("Content-Type", "application/octet-stream")

$body = "<file-contents-here>"

$response = Invoke-RestMethod 'https://192.168.3.125:9000/tiktok/example.mp4?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=LH1I4a83WdghqEmTnZmw%2F20250221%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20250221T172240Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&X-Amz-Signature=7b2a63f7e139c3d4a6756fa72771f8f5c230565d826f0121f2d39263e79613ca' -Method 'PUT' -Headers $headers -Body $body
$response | ConvertTo-Json
```

js fetch:
```js
var myHeaders = new Headers();
myHeaders.append("Content-Type", "application/octet-stream");

var file = "<file contents here>";

var requestOptions = {
   method: 'PUT',
   headers: myHeaders,
   body: file,
   redirect: 'follow'
};

fetch("https://192.168.3.125:9000/tiktok/example.mp4?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=LH1I4a83WdghqEmTnZmw%2F20250221%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20250221T172240Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&X-Amz-Signature=7b2a63f7e139c3d4a6756fa72771f8f5c230565d826f0121f2d39263e79613ca", requestOptions)
   .then(response => response.text())
   .then(result => console.log(result))
   .catch(error => console.log('error', error));

```

go:
```go
package main

import (
   "fmt"
   "strings"
   "net/http"
   "io/ioutil"
)

func main() {

   url := "https://192.168.3.125:9000/tiktok/example.mp4?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=LH1I4a83WdghqEmTnZmw%252F20250221%252Fus-east-1%252Fs3%252Faws4_request&X-Amz-Date=20250221T172240Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&X-Amz-Signature=7b2a63f7e139c3d4a6756fa72771f8f5c230565d826f0121f2d39263e79613ca"
   method := "PUT"

   payload := strings.NewReader("<file contents here>")


   client := &http.Client {
   }
   req, err := http.NewRequest(method, url, payload)

   if err != nil {
      fmt.Println(err)
      return
   }
   req.Header.Add("Content-Type", "application/octet-stream")

   res, err := client.Do(req)
   if err != nil {
      fmt.Println(err)
      return
   }
   defer res.Body.Close()

   body, err := ioutil.ReadAll(res.Body)
   if err != nil {
      fmt.Println(err)
      return
   }
   fmt.Println(string(body))
}
```

检验最终结果:
![img_2.png](img_2.png)
