package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smallnest/rpcx/client"
	"net/url"
	"strings"
	"sync"
	"time"
)

var zks = &zkS{}

type zkS struct {
	sync.RWMutex
	data map[string]*client.ZookeeperDiscovery
}
type OneClient struct {
	mu       sync.RWMutex
	xClients map[string]client.XClient
}

var o OneClient
var oOnce sync.Once

func GetZD(basePath, servicePath string) (*client.ZookeeperDiscovery, error) {
	//if zks.data == nil {
	//	zks = &zkS{}
	//	zks.data = make(map[string]*client.ZookeeperDiscovery)
	//}
	//zks.RLock()
	//if v, ok := zks.data[basePath+"/"+servicePath]; ok {
	//	zks.RUnlock()
	//	return v, nil
	//}
	//zks.RUnlock()
	zd, err := client.NewZookeeperDiscovery("smartPms", "property", []string{"127.0.0.1:2181"}, nil)
	if err != nil {
		return zd, err
	}
	//zks.Lock()
	//	zks.data[basePath+"/"+servicePath] = zd
	//zks.Unlock()
	return zd, err
}
func main() {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	v1 := r.Group("v1", func(context *gin.Context) {})
	pc := v1.Group("/pc", func(context *gin.Context) {})
	pc.Any("/:srvName/:method", func(c *gin.Context) {
		var resp interface{}
		c.Request.ParseForm()
		type Param struct {
			Url  url.Values
			Post *url.Values
		}
		method := strings.ToUpper(c.Param("method")[:1]) + c.Param("method")[1:]
		err := InvokeService(c, c.Param("srvName"), method, Param{Url: c.Request.URL.Query(), Post: &c.Request.PostForm}, &resp)
		fmt.Println("ccc", err)
		if err != nil && strings.Contains(err.Error(), "can not found any server") {
			//_, err := refreshRpcXClient(c.Param("srvName"))
			fmt.Println("ccc refresh error", err)
		}
		if err != nil {
			c.JSON(500, "errorr"+err.Error())
		} else {
			c.JSON(200, resp)
		}
		return
	})
	r.Run(":8068")
}

func InvokeService(ctx context.Context, srvName, method string, args, resp interface{}) error {
	xc, err := getRpcClient(srvName)
	if err != nil {
		return err
	}
	return xc.Call(ctx, method, args, resp)
}

func getRpcClient(srvName string) (client.XClient, error) {
	if o.xClients == nil {
		o.xClients = make(map[string]client.XClient)
	}
	o.mu.RLock()
	if v, ok := o.xClients[srvName]; ok {
		o.mu.RUnlock()
		return v, nil
	}
	o.mu.RUnlock()
	zd, err := GetZD("smartPms", srvName)
	if err != nil {
		return nil, err
	}
	ops := client.DefaultOption
	ops.HeartbeatInterval = time.Second
	ops.Heartbeat = true
	xc := client.NewXClient(srvName, client.Failover, client.RandomSelect, zd, ops)
	o.mu.Lock()
	o.xClients[srvName] = xc
	o.mu.Unlock()
	return xc, nil
}

func refreshRpcXClient(srvName string) (client.XClient, error) {
	zd, err := GetZD("smartPms", srvName)
	if err != nil {
		return nil, err
	}
	ops := client.DefaultOption
	ops.HeartbeatInterval = time.Second
	ops.Heartbeat = true
	xc := client.NewXClient(srvName, client.Failover, client.RandomSelect, zd, ops)
	return xc, nil
}
