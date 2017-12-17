package yext

import (
	"encoding/json"
	"reflect"
	"testing"
)

func jsonString(l *Location) (error, string) {
	buf, err := json.Marshal(l)
	if err != nil {
		return err, ""
	}

	return nil, string(buf)
}

func TestJSONSerialization(t *testing.T) {
	type test struct {
		l    *Location
		want string
	}

	tests := []test{
		{&Location{}, `{}`},
		{&Location{City: nil}, `{}`},
		{&Location{City: String("")}, `{"city":""}`},
		{&Location{Languages: nil}, `{}`},
		{&Location{Languages: nil}, `{}`},
		{&Location{Languages: &[]string{}}, `{"languages":[]}`},
		{&Location{Languages: &[]string{"English"}}, `{"languages":["English"]}`},
		{&Location{HolidayHours: nil}, `{}`},
		{&Location{HolidayHours: &[]HolidayHours{}}, `{"holidayHours":[]}`},
	}

	for _, test := range tests {
		if err, got := jsonString(test.l); err != nil {
			t.Error("Unable to convert", test.l, "to JSON:", err)
		} else if got != test.want {
			t.Errorf("json.Marshal(%#v) = %s; expected %s", test.l, got, test.want)
		}
	}
}

func TestJSONDeserialization(t *testing.T) {
	type test struct {
		json string
		want *Location
	}

	tests := []test{
		{`{}`, &Location{}},
		{`{"emails": []}`, &Location{Emails: Strings([]string{})}},
		{`{"emails": ["mhupman@yext.com", "bmcginnis@yext.com"]}`, &Location{Emails: Strings([]string{"mhupman@yext.com", "bmcginnis@yext.com"})}},
	}

	for _, test := range tests {
		v := &Location{}

		if err := json.Unmarshal([]byte(test.json), v); err != nil {
			t.Error("Unable to deserialize", test.json, "from JSON:", err)
		} else if !reflect.DeepEqual(v, test.want) {
			t.Errorf("json.Unmarshal(%#v) = %s; expected %s", test.json, v, test.want)
		}
	}
}

func TestSampleJSONResponseDeserialization(t *testing.T) {
	v := &Location{}
	if err := json.Unmarshal([]byte(sampleLocationJSON), v); err != nil {
		t.Error("Unable to deserialize", sampleLocationJSON, "from JSON:", err)
	}
}

var sampleLocationJSON = `{
	 "id":"1138",
	 "timestamp":1483891079283,
	 "accountId":"479390",
	 "locationName":"Best Buy",
	 "address":"105 Topsham Fair Mall Rd",
	 "city":"Topsham",
	 "state":"ME",
	 "zip":"04086",
	 "countryCode":"US",
	 "language":"en",
	 "phone":"8882378289",
	 "isPhoneTracked":false,
	 "localPhone":"2077985780",
	 "categoryIds":[
			"77",
			"67",
			"262",
			"1376",
			"333"
	 ],
	 "googleAttributes":[
			{
				 "id":"has_gift_wrapping"
			},
			{
				 "id":"has_in_store_pickup",
				 "optionIds":[
						"true"
				 ]
			},
			{
				 "id":"has_service_installation",
				 "optionIds":[
						"true"
				 ]
			},
			{
				 "id":"has_service_repair",
				 "optionIds":[
						"true"
				 ]
			},
			{
				 "id":"has_wheelchair_accessible_elevator"
			},
			{
				 "id":"has_wheelchair_accessible_entrance"
			},
			{
				 "id":"pay_check"
			},
			{
				 "id":"pay_credit_card"
			},
			{
				 "id":"pay_debit_card"
			},
			{
				 "id":"pay_mobile_nfc"
			},
			{
				 "id":"requires_cash_only"
			},
			{
				 "id":"sells_goods_name_brand",
				 "optionIds":[
						"true"
				 ]
			}
	 ],
	 "featuredMessage":"Black Friday 2016",
	 "featuredMessageUrl":"http://www.bestbuy.com/site/electronics/black-friday/pcmcat225600050002.c?id\u003dpcmcat225600050002\u0026ref\u003dNS\u0026loc\u003dns100",
	 "websiteUrl":"[[BEST BUY LOCAL PAGES URL]]/?ref\u003dNS\u0026loc\u003dns100",
	 "displayWebsiteUrl":"[[BEST BUY LOCAL PAGES URL]]/?ref\u003dNS\u0026loc\u003dns100",
	 "logo":{
			"url":"http://www.yext-static.com/cms/79fb6255-2e84-435f-9864-aa0572fe4cbd.png"
	 },
	 "photos":[
			{
				 "url":"http://www.yext-static.com/cms/d99bf85e-9104-47f6-9c45-bbdc619650af.jpg"
			},
			{
				 "url":"http://www.yext-static.com/cms/f875ecbb-5890-4a69-807f-ec3d34e888f0.jpg"
			},
			{
				 "url":"http://www.yext-static.com/cms/dd278927-7749-4291-a85a-2881dfd626ea.jpg"
			},
			{
				 "url":"http://www.yext-static.com/cms/0b675609-2b18-4c16-8dd6-e7bd435bbc47.jpg"
			}
	 ],
	 "videoUrls":[
			"http://www.youtube.com/watch?v\u003dGXJWmvbn7PI"
	 ],
	 "twitterHandle":"BestBuy",
	 "facebookPageUrl":"https://www.facebook.com/bestbuy",
	 "displayLat":43.93842,
	 "displayLng":-69.98002,
	 "yextDisplayLat":43.93604,
	 "yextDisplayLng":-69.984222,
	 "yextRoutableLat":43.9352,
	 "yextRoutableLng":-69.983363,
	 "folderId":"54748",
	 "customFields":{
			"7880":"https://www.timetrade.com/app/bestbuy/workflows/PREC1234/find",
			"6031":"At Best Buy [[Geomodifier]], we specialize in helping you find the best technology to fit the way you live. Together, we can transform your living space with the latest HDTVs, computers, smart home technology, and gaming consoles like Xbox One, PlayStation 4 and Wii U. We can walk you through updating your appliances with cutting-edge refrigerators, ovens, washers and dryers. We’ll also show you how to make the most of your active lifestyle with our huge selection of smartphones, tablets and wearable technology. At Best Buy [[Geomodifier]], we’ll keep your devices running smoothly with the full range of expert services from Geek Squad®. We’re here to help, so visit us at [[ADDRESS]] in [[CITY]], [[STATE]] to find the perfect new camera, laptop, Blu-ray player, smart lighting or activity tracker today.",
			"7882":"We repair almost every major brand – no matter where you bought it. TV repair, refrigerator repair, dishwasher repair, washing machine repair or dryer repair, we’ll come to your home to diagnose the problem.",
			"7881":"Appliance Repair \u0026 TV Repair",
			"5584":{
				 "url":"http://a.mktgcdn.com/p/mnkdMgSNOcvL036kWByOGAx-2bmQx7ZhNapg9x5rUF8/520x303.jpg",
				 "description":"SLOW COMPUTER? HAVE A VIRUS? GEEK SQUAD CAN HELP.",
				 "clickthroughUrl":"http://www.bestbuy.com/247support/"
			},
			"6036":{
				 "url":"http://a.mktgcdn.com/p/viDc_593ilLjIZv8kVIUjoVotYpg646lageeCbbPLPM/552x368.jpg"
			},
			"7884":"Our Agents provide repair, installation and setup services on all kinds of tech – including computer repair, setup and support, TV repair, home theater installation, car stereo installation and home appliance repair. We fix most makes and models, no matter where you bought them, and can show you how to get the most out of your technology.",
			"5585":"SLOW COMPUTER? HAVE A VIRUS? GEEK SQUAD CAN HELP.",
			"7883":"http://www.bestbuy.com/site/clp/tv-appliance-repair/pcmcat339300050004.c?id\u003dpcmcat339300050004",
			"6114":[

			],
			"7877":"Professional Car Electronics \u0026 Car Stereo Installation",
			"7876":"http://www.bestbuy.com/site/electronics/geek-squad/pcmcat138100050018.c?id\u003dpcmcat138100050018",
			"6625":"Shop the Deals",
			"7879":"https://autotech-scheduling.geeksquad.com/",
			"7878":"Let our expert Autotechs install your car audio, video, security and more.",
			"5591":"LOST YOUR DATA?",
			"5592":"http://www.geeksquad.com/scheduling",
			"5593":"Get Help Now",
			"5594":"Our Agents are standing by to help you. Starting at $39.99.",
			"5595":"https://www.timetrade.com/app/bestbuy/workflows/ABCD1234/schedule/",
			"5596":"Device Setup",
			"9558":"true",
			"5597":"Leave with confidence and control over your new technology. We perform a personalized setup based on your needs and interests. Then, we show you how your new device works and point out the features you\u0027ll wonder how you ever lived without. Let us help you set up your new computer (PC or Mac), tablet, smartphone, digital camera, gaming console or handheld system.",
			"5598":"https://www.timetrade.com/app/bestbuy/workflows/ABCD1234/schedule/",
			"6160":"At Best Buy, we make technology work for you at with a full selection of TVs, computers, cameras, gaming consoles, appliances, cell phones, tablets, Geek Squad® services and more.",
			"5590":{
				 "url":"http://a.mktgcdn.com/p/ixPrMgn5ZLVM687aEcGa2edY6bSU6PZdiXF0Ik-sZDY/520x303.jpg",
				 "description":"LOST YOUR DATA?",
				 "clickthroughUrl":"http://www.bestbuy.com/site/geek-squad/geek-squad-data-recovery/pcmcat748300502324.c?id\u003dpcmcat748300502324"
			},
			"6010":"The Best Buy Topsham store is permanently closing at the end of the day on October 29th. We apologize for the inconvenience. While this store is closing, Best Buy is in your community to stay like our Best Buy Portland store at 364 Maine Mall Rd., #301, South Portland, ME 04106. To find another location, use the store locator or shop 24/7 on BestBuy.com.",
			"6176":"http://stores.bestbuy.com/me/topsham/105-topsham-fair-mall-rd-1138.html",
			"5882":[
				 "2197",
				 "2198",
				 "2199",
				 "2200",
				 "2201",
				 "2202"
			],
			"5565":"https://www-ssl.bestbuy.com/usw/covertops",
			"6250":[
				 "Appliance Store",
				 "Computer Store",
				 "Home Theater Store",
				 "Music Store",
				 "Video Game Store",
				 "Video Store"
			],
			"5599":"1-on-1 Tutorial with an Agent",
			"6249":[
				 "Electronics Store"
			],
			"5754":"Topsham",
			"6009":"Best Buy Topsham is Closing October 29",
			"5570":"In Store",
			"5571":"Bring your product to a Best Buy store and speak with an Agent in person.",
			"5890":"No one stands behind you like Geek Squad. The Geek Squad Agents at Best Buy [[Geomodifier]] are ready to help. More than just a computer store, we’ve got the tools, knowledge and experience to turn questions into answers and issues into fixes. Whether you’re in need of cell phone repair or you’re wondering “Where is there reliable appliance repair near me?” visit us at [[ADDRESS]] to see how we can help and learn more about our services including appliance repair, tablet and computer repair and setup, TV repair and home theater setup, WiFi/networking setup, car stereo installation and GPS setup, and technology consultation. We’ll give you advice on how you can get the most out of your technology and help troubleshoot it when it’s not working properly. ",
			"5572":"https://www-ssl.bestbuy.com/services/triage/home",
			"5891":"Geek Squad Agents across the U.S. are trained to work on a full range of technology, including computers, tablets, TVs, home theater, car audio, home appliances and more. We’ll install, set up, protect, support or repair your product. Because bad luck deserves good coverage, Geek Squad Protection Plans are also available at your local Best Buy [[Geomodifier]] to ensure your technology is working like new. We support and fix most brands, makes and models – no matter where you bought it, unlike many computer stores and appliance stores. So for those looking for more than just “computer stores near me,” Best Buy [[Geomodifier]] gives you expert cell phone repair, computer repair, washer repair, refrigerator repair, dryer repair and dishwasher repair.",
			"5573":"https://www-ssl.bestbuy.com/services/scheduling/tts/precinct/modify",
			"5892":"Can’t make it into the Best Buy [[Geomodifier]] store? Geek Squad has over 20,000 Agents available 24/7/365* to help online and over the phone. Plus, we can always schedule an Agent to come out to your home or office.",
			"7950":"false",
			"5574":"In Home",
			"5893":"*24/7/365 support availability is limited to remote online chat and over the phone.",
			"5575":"We can set up your products, help you use them, and diagnose any issues. Call 1-800-433-5778.",
			"7875":"https://www-ssl.bestbuy.com/usw/repairstatus/home",
			"7874":{
				 "url":"http://a.mktgcdn.com/p/cPr99O8i_judFZ5JAyj1xH6Xxr5hkDh3ZOXHAC8KgzM/276x152.jpg"
			},
			"5566":"Online \u0026 Phone",
			"5567":"Need something fixed ASAP? Have questions? Agents are available 24/7. Call 1-800-433-5778",
			"5600":"Questions answered. Technology simplified. A Geek Squad Agent can teach you the basic and advanced features of products like iPhone, iPad, Windows 10 device, Mac OS X computer, Android computer and DSLR camera.",
			"5568":"https://www-ssl.bestbuy.com/usw/covertops",
			"5601":"https://www.timetrade.com/app/bestbuy/workflows/ABCD1234/schedule/",
			"5569":"https://www-ssl.bestbuy.com/usw/repairstatus/home",
			"6019":{
				 "url":"http://a.mktgcdn.com/p/ZTJ9jdKRiI085sotLoDziEU7AvgZkref2sv7WqS4JBg/544x300.jpg",
				 "description":"Great deals happening now",
				 "clickthroughUrl":"http://www.bestbuy.com/site/global-promotions/special-sale/pcmcat185700050011.c?id\u003dpcmcat185700050011"
			},
			"5925":{
				 "url":"http://a.mktgcdn.com/p/T4bbCO2_y82H7PpHtjEoh7kK9Qc5autgd2Ko0xfuZEc/642x357.jpg"
			}
	 },
	 "closed":{
			"isClosed":true
	 },
	 "labelIds":[
			"12908",
			"19092"
	 ],
	 "productListIds":[
			"445835",
			"445834"
	 ],
	 "eventListsLabel":"",
	 "eventListIds":[
			"Event Calendars-1138"
	 ],
	 "locationType":"LOCATION"
}`
