// TODO: Make these into real *testing.T tests
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"gopkg.in/yext/yext-go.v2"
)

func main() {
	client := yext.NewClient(yext.NewDefaultConfig().WithEnvCredentials())

	llr, _, err := client.LocationService.List(nil)
	fatalif(err)
	log.Println(printllr(llr))

	llr, _, err = client.LocationService.List(&yext.ListOptions{})
	fatalif(err)
	log.Println(printllr(llr))

	llr, _, err = client.LocationService.List(&yext.ListOptions{Limit: 2})
	fatalif(err)
	log.Println(printllr(llr))

	llr, _, err = client.LocationService.List(&yext.ListOptions{Offset: 1})
	fatalif(err)
	log.Println(printllr(llr))

	llr, _, err = client.LocationService.List(&yext.ListOptions{Offset: 5})
	fatalif(err)
	log.Println(printllr(llr))

	locs, err := client.LocationService.ListAll()
	fatalif(err)
	log.Println(printlocs(locs))

	cfs, err := client.CustomFieldService.ListAll()
	fatalif(err)
	log.Println(printjson(cfs))

	fs, err := client.FolderService.ListAll()
	fatalif(err)
	log.Println(printjson(fs))

	plr, _, err := client.ListService.ListProductLists(nil)
	fatalif(err)
	log.Println(printjson(plr.ProductLists))

	ps, err := client.ListService.ListAllProductLists()
	fatalif(err)
	log.Println(printjson(ps))

	blr, _, err := client.ListService.ListBioLists(nil)
	fatalif(err)
	log.Println(printjson(blr.BioLists))

	bs, err := client.ListService.ListAllBioLists()
	fatalif(err)
	log.Println(printjson(bs))

	elr, _, err := client.ListService.ListEventLists(nil)
	fatalif(err)
	log.Println(printjson(elr.EventLists))

	es, err := client.ListService.ListAllEventLists()
	fatalif(err)
	log.Println(printjson(es))

	us, err := client.UserService.ListAll()
	fatalif(err)
	log.Println(printjson(us))

	rs, _, err := client.UserService.ListRoles()
	fatalif(err)
	log.Println(printjson(rs))
}

func fatalif(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func printllr(llr *yext.LocationListResponse) string {
	return fmt.Sprintf("count=%d, locs=%s", llr.Count, printlocs(llr.Locations))
}

func printlocs(locs []*yext.Location) string {
	var ids []string
	for _, loc := range locs {
		ids = append(ids, loc.GetId())
	}
	return strings.Join(ids, ",")
}

func printjson(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}
