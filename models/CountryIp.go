package models

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

const IPV4_URL = "https://ipv4.fetus.jp/"

var COUNTRIES = []map[string]string{
	{
		"name": "laos",
		"code": "la",
	},
	// {
	// 	"name": "ghana",
	// 	"code": "gh",
	// },
	// {
	// 	"name": "philippines",
	// 	"code": "ph",
	// },
	// {
	// 	"name": "malaysia",
	// 	"code": "my",
	// },
	// {
	// 	"name": "nigeria",
	// 	"code": "ng",
	// },
	// {
	// 	"name": "turkey",
	// 	"code": "tr",
	// },
	// {
	// 	"name": "singapore",
	// 	"code": "sg",
	// },
	// {
	// 	"name": "cambodia",
	// 	"code": "kh",
	// },
	// {
	// 	"name": "unitedstates",
	// 	"code": "us",
	// },
	// {
	// 	"name": "unitedarabemirates",
	// 	"code": "ae",
	// },
	// {
	// 	"name": "canada",
	// 	"code": "ca",
	// },
	// {
	// 	"name": "thailand",
	// 	"code": "th",
	// },
	// {
	// 	"name": "pakistan",
	// 	"code": "pk",
	// },
	// {
	// 	"name": "southafrica",
	// 	"code": "za",
	// },
	// {
	// 	"name": "vanuatu",
	// 	"code": "vu",
	// },
	// {
	// 	"name": "seychelles",
	// 	"code": "sc",
	// },
	// {
	// 	"name": "brazil",
	// 	"code": "br",
	// },
	// {
	// 	"name": "mexico",
	// 	"code": "mx",
	// },
	// {
	// 	"name": "ci",
	// 	"code": "ci",
	// },
}

/* struct
type country struct {
    name string
    code string
}

var COUNTRIES = []country{
    {
        name: "laos",
        code: "la",
    },
*/

func SubProcess() {

	log.Println("start country ip")

	// if opts["code"] != "" {
	// 	if opts["code"] != country["code"] {
	// 		continue
	// 	}
	// }

	for _, country := range COUNTRIES {
		fmt.Println(country)
		count := deleteOldIp(country)
		count = RegisterIp(country)
		fmt.Println(count)
	}

}

func ReadIp() {
	countryCode := "ph"
	res := readFile(countryCode)
	fmt.Println(res)
	// filePath := fmt.Sprintf("%s%s.txt", IPV4_URL, countryCode)
	// fmt.Println(filePath)

	// resp, err := http.Get(filePath)
	// if err != nil {
	// 	fmt.Printf("[ERROR] unread %s", filePath)
	// 	// return nil
	// }
	// fmt.Println(resp)

	// r := csv.NewReader(f)

	// for {
	// 	record, err := r.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(record)
	// }
}

func RegisterIp(country map[string]string) int {
	countryCode := country["code"]
	countryName := country["name"]

	list := readFile(countryCode)

	loop := 0

	for _, cidr := range list {
		fmt.Println(cidr)
		fmt.Println(countryCode, countryName)

		loop++
	}
	// YoubrideDataCountryIP.UpdateOrCreate(map[string]string{
	// 	"cidr": cidr,
	// }, map[string]string{
	// 	"country_code": countryCode,
	// 	"country":      countryName,
	// })

	// if loop%100 != 0 {
	// 	time.Sleep(time.Second)
	// }

	return loop
}

func deleteOldIp(country map[string]string) int {
	sql := fmt.Sprintf("DELETE FROM country_ip WHERE country_code = %s", country["code"])
	test := fmt.Sprintf(sql, "test")
	fmt.Println(test)
	// dbh = Youbride::Data::CountryIp->driver->rw_handle;
	// my $sth = $dbh->prepare($sql);
	// my $ret = $sth->execute($country->{code});

	// return ret if $ret =~ /^\d+$/
	return 0
}

func readFile(countryCode string) []string {
	filePath := fmt.Sprintf("%s%s.txt", IPV4_URL, countryCode)
	// fmt.Println(filePath)
	resp, err := http.Get(filePath)
	if err != nil {
		fmt.Printf("[ERROR] unread %s", filePath)
		return nil
	}
	// fmt.Println(resp)
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	response := strings.Split(string(content), "\n")
	// fmt.Printf("response: %v", response)
	var list []string
	for _, line := range response {
		// fmt.Printf("Lines of contents: %v", line)
		r := regexp.MustCompile("^[0-9]")
		matched := r.MatchString(line)
		if matched {
			list = append(list, line)
		}
	}
	return list
}

/* 構造体

type country struct {
    name string
    code string
}

type countryIp struct {
    countryCode string
    countryName string
}

var COUNTRIES = []country{
    {
        name: "laos",
        code: "la",
    },
    {
        name: "ghana",
        code: "gh",
    },
    {
        name: "philippines",
        code: "ph",
    },
    {
        name: "malaysia",
        code: "my",
    },
    {
        name: "nigeria",
        code: "ng",
    },
*/
