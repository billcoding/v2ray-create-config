package main

import (
	"embed"
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type vnode struct {
	address string
	uuid    string
	wspath  string
	port    int
}

var (
	help          = flag.Bool("h", false, "Help info")
	nodeFile      = flag.String("n", "nodes.txt", "The v2ray nodes spec file")
	nodeTplFile   = flag.String("np", "node.tpl", "The v2ray node template file")
	configTplFile = flag.String("cp", "config.tpl", "The v2ray client config file")
	outputFile    = flag.String("o", "config.json", "The v2ray output config file")
	uploadFile    = flag.Bool("u", false, "The v2ray upload to Aliyun OSS")
	nodeText      = ""
	nodeTplText   = ""
	configTplText = ""

	//go:embed tpls/config.tpl tpls/node.tpl
	FS embed.FS
)

func main() {
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	nodeBytes, err := ioutil.ReadFile(*nodeFile)
	if err != nil {
		panic(err)
	}
	nodeText = string(nodeBytes)

	nodeTplBytes, err := ioutil.ReadFile(*nodeTplFile)
	if err != nil {
		nodeTplBytes, err = FS.ReadFile("tpls/node.tpl")
	}
	nodeTplText = string(nodeTplBytes)

	configTplBytes, err := ioutil.ReadFile(*configTplFile)
	if err != nil {
		configTplBytes, err = FS.ReadFile("tpls/config.tpl")
	}
	configTplText = string(configTplBytes)

	nodes := getNodes()
	nodeConfigs := getNodeConfig(nodes)
	config := getConfig(strings.Join(nodeConfigs, ","))

	if *uploadFile {
		uploadConfig(config)
		return
	}

	createConfig(config)
}

func getNodes() []vnode {
	nodeList := strings.Split(nodeText, "\n")
	if nil == nodeList || len(nodeList) <= 0 {
		panic("no found nodes spec")
	}
	nodes := make([]vnode, 0)
	for _, nodeLine := range nodeList {
		nodeSpec := strings.TrimSpace(nodeLine)
		if nodeSpec == "" {
			continue
		}
		nodeSpecs := strings.Split(nodeSpec, ",")

		if len(nodeSpecs) < 4 {
			continue
		}
		address := strings.TrimSpace(nodeSpecs[0])
		portStr := strings.TrimSpace(nodeSpecs[1])
		uuid := strings.TrimSpace(nodeSpecs[2])
		wspath := strings.TrimSpace(nodeSpecs[3])
		if address == "" {
			continue
		}
		if portStr == "" {
			continue
		}
		if uuid == "" {
			continue
		}
		if wspath == "" {
			continue
		}
		port, err := strconv.Atoi(portStr)
		if err != nil {
			fmt.Println("invalid port:" + portStr)
			continue
		}
		// try ping
		if telnet(address, port) {
			nodeInfo := vnode{
				address: address,
				uuid:    uuid,
				wspath:  wspath,
				port:    port,
			}
			nodes = append(nodes, nodeInfo)
			fmt.Println("Add : " + nodeSpec)
		}
	}
	return nodes
}

func getNodeConfig(nodes []vnode) []string {
	nodeConfigs := make([]string, 0)
	if nodes != nil {
		for _, node := range nodes {
			nodeConfigs = append(nodeConfigs, fmt.Sprintf(nodeTplText, node.address, node.port, node.uuid, node.wspath))
		}
	}
	return nodeConfigs
}

func getConfig(nodeConfigs string) string {
	return fmt.Sprintf(configTplText, nodeConfigs)
}

func createConfig(config string) {
	_ = ioutil.WriteFile(*outputFile, []byte(config), 0766)
}
