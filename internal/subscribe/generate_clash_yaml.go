package subscribe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"

	"gopkg.in/yaml.v2"
)

var (
	OutputFile string
)

func (c *Clash) LoadTemplate(path string, proxies []interface{}) []byte {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		log.Printf("[%s] template doesn't exist.", path)
		return nil
	}
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("[%s] template open the failure.", path)
		return nil
	}
	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		log.Printf("[%s] Template format error.", path)
	}

	c.Proxy = nil

	var proxy []map[string]interface{}
	var proxiesStr []string
	names := map[string]int{}

	for _, proto := range proxies {
		o := reflect.ValueOf(proto)
		nameField := o.FieldByName("Name")
		proxyItem := make(map[string]interface{})
		j, _ := json.Marshal(proto)
		_ = json.Unmarshal(j, &proxyItem)

		name := nameField.String()
		if index, ok := names[name]; ok {
			names[name] = index + 1
			name = fmt.Sprintf("%s%d", name, index+1)
		} else {
			names[name] = 0
		}

		proxyItem["name"] = name
		proxy = append(proxy, proxyItem)
		c.Proxy = append(c.Proxy, proxyItem)
		proxiesStr = append(proxiesStr, name)
	}

	c.Proxy = proxy

	for _, group := range c.ProxyGroup {
		groupProxies := group["proxies"].([]interface{})
		for i, p := range groupProxies {
			if p == "1" {
				groupProxies = groupProxies[:i]
				var tmpGroupProxies []string
				for _, s := range groupProxies {
					tmpGroupProxies = append(tmpGroupProxies, s.(string))
				}
				tmpGroupProxies = append(tmpGroupProxies, proxiesStr...)
				group["proxies"] = tmpGroupProxies
				break
			}
		}

	}

	d, err := yaml.Marshal(c)
	if err != nil {
		return nil
	}

	return d
}

func GenerateClashConfig(proxies []interface{}) ([]byte, error) {
	clash := Clash{}

	r := clash.LoadTemplate(OutputFile, proxies)
	if r == nil {
		return nil, fmt.Errorf("sublink 返回数据格式不对")
	}
	return r, nil
}
