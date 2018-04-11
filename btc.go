package main

//================分割线===============
 比特币  非对称加密  公钥私钥
 钱不是支付给个人的  而是支付给一把私钥的
说到底，比特币只是区块链的一条记录，是凭空生成的
//================分割线===============
代码中大量重复着if err!= nil { return err} 这段snippet。但是如果你全面浏览过Go标准库中的代码，
你会发现像上面这样的代码并不多见。Rob Pike曾经在《errors are values》一文中针对这个问题做过解释，
并给了stdlib中的一些消除重复的方法：那就是将error作为一个内部状态：

  //bufio/bufio.go
  type Writer struct {
      err error
      buf []byte
      n   int
      wr  io.Writer
  }
  func (b *Writer) Write(p []byte) (nn int, err error) {
      if b.err != nil {
          return nn, b.err
      }
      ... ...
  }

  //writer_demo.go
  buf := bufio.NewWriter(fd)
  buf.WriteString("hello, ")
  buf.WriteString("gopherchina ")
  buf.WriteString("2017")
  if err := buf.Flush() ; err != nil {
      return err
  }
//================分割线===============
nats对网络编程的处理
//================分割线===============
1<<7-1
//================分割线===============

Don.t lock around I/O
// BAD: Don.t do this.
func root(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	count++

	msg := []byte(strings.Repeat(fmt.Sprintf("%d", count), payloadBytes))
	w.Write(msg)
}
//================分割线===============
//httpclient 是并发的  MaxIdleConnsPerHost参数
要关闭body   每次请求
func createHTTPClient() *http.Client {
    client := &http.Client{
        Transport: &http.Transport{
            MaxIdleConnsPerHost: MaxIdleConnections,
        },
        Timeout: time.Duration(RequestTimeout) * time.Second,
    }

    return client
}
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
//================分割线===============
