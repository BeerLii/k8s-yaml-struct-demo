package main

import (
	"bwei.com/test/parser"
	"fmt"
	"k8s.io/client-go/kubernetes/scheme"
	"strings"
)

var deployment = `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-nginx
spec:
  replicas: 2
  template:
    metadata:
      labels:
        run: my-nginx
    spec:
      containers:
      - name: my-nginx
        image: nginx
        ports:
        - containerPort: 80
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-nginx2
spec:
  replicas: 3
  template:
    metadata:
      labels:
        run: my-nginx2
    spec:
      containers:
      - name: my-nginx2
        image: nginx
        ports:
        - containerPort: 80


`

func run() error {
	const templateFile = "template/template.yaml"
	const dataFile = "template/values.yaml"
	const outputFile = "parsed/parsed.yaml"
	if err := parser.Parse(templateFile, dataFile, outputFile); err != nil {
		return err
	}
	fmt.Printf("file %s was generated.\n", outputFile)
	return nil
}

func main() {
	//if err := run(); err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//deploys := &v1.Deployment{}
	//data, err := yaml.YAMLToJSON([]byte(deployment))
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(data))
	//err = json.Unmarshal(data, deploys)
	//if err != nil {
	//	fmt.Println(err)
	//}
	sepYamlfiles := strings.Split(deployment, "---")
	for _, sepYamlfile := range sepYamlfiles {
		decode := scheme.Codecs.UniversalDeserializer()
		obj, groupVersionKind, err := decode.Decode([]byte(sepYamlfile), nil, nil)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(groupVersionKind)
		fmt.Println(obj)
	}

}
