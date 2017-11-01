package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/pinzolo/casee"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}

func ToGo(s string) string {
	return casee.ToPascalCase(s)
}
func run() error {
	topLevel, err := FetchAPIInfo("")
	if err != nil {
		return err
	}
	var keys []string
	for key := range topLevel {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	//egFnames := map[string]string{}
	fmt.Printf("type Services struct {\n")
	for _, k := range keys {
		fmt.Printf("	%s struct {\n", ToGo(k))
		mounts := topLevel[k].(map[string]interface{})["mounts"].(map[string]interface{})
		var endpointGroups []string
		for m, mountDetails := range mounts {
			_ = mountDetails
			endpointGroups = append(endpointGroups, m)
			//egFnames[m] = mountDetails.(map[string]interface{})["friendlyName"].(string)
		}
		sort.Strings(endpointGroups)
		for _, eg := range endpointGroups {
			fmt.Printf("		%s %sService\n", ToGo(eg), ToGo(eg))
		}
		fmt.Printf("	}\n")
	}
	fmt.Printf("}\n")
	for _, k := range keys {
		mounts := topLevel[k].(map[string]interface{})["mounts"].(map[string]interface{})
		var endpointGroups []string
		for m := range mounts {
			endpointGroups = append(endpointGroups, m)
		}
		sort.Strings(endpointGroups)
		for _, eg := range endpointGroups {
			mount := mounts[eg].(map[string]interface{})
			mountPath := fmt.Sprintf("/%s", mount["mountName"])
			groupDetails, err := FetchEndpointGroupInfo(mountPath)
			if err != nil {
				return err
			}
			_ = groupDetails

			var endpoints []string
			fnames := map[string]string{}
			for _, ep := range groupDetails {
				epname := ep.(map[string]interface{})["name"].(string)
				endpoints = append(endpoints, epname)
				fnames[epname] = ep.(map[string]interface{})["friendlyName"].(string)

			}
			sort.Strings(endpoints)
			//fmt.Printf("// %sService provides access to %s\n", egFnames[eg])
			fmt.Printf("type %sService interface {\n", ToGo(eg))
			for _, e := range endpoints {
				endpointDetail, err := FetchAPIInfo(fmt.Sprintf("/%s/%s", eg, e))
				//spew.Dump(endpointDetail, err)
				if err != nil {
					return err
				}
				_ = endpointDetail
				fqen := fmt.Sprintf("%s%s", ToGo(eg), ToGo(e))
				fmt.Printf("	// %s - %s\n", ToGo(e), fnames[e])
				fmt.Printf("	%s(%sRequest) (%sResponse, error)\n", ToGo(e), fqen, fqen)
			}
			fmt.Printf("}\n")
		}
	}
	return nil
}
