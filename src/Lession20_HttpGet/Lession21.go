package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

/*
// 回传数据样例
{"kind": "Listing",
	"data": {
		"modhash": "",
		"children": [
			{ "kind": "t3", "data":
				{
					"domain": "austingwalters.com", "banned_by": null, "media_embed": {}, "subreddit": "golang", "selftext_html": null, "selftext": "", "likes": null, "secure_media": null, "link_flair_text": null, "id": "26l1by", "gilded": 0, "secure_media_embed": {}, "clicked": false, "stickied": false, "author": "dgryski", "media": null, "score": 2, "approved_by": null, "over_18": false, "hidden": false, "thumbnail": "", "subreddit_id": "t5_2rc7j", "edited": false, "link_flair_css_class": null, "author_flair_css_class": null, "downs": 8, "saved": false, "is_self": false,
					"permalink": "/r/golang/comments/26l1by/building_a_web_server_in_go_database_accesses/", "name": "t3_26l1by", "created": 1401198126.0,
					"url": "http://austingwalters.com/building-a-web-server-in-go-database-accesses/", "author_flair_text": null,
					"title": "Building a Web Server in Go: Database Accesses", "created_utc": 1401169326.0, "ups": 10, "num_comments": 1, "visited": false, "num_reports": null, "distinguished": null
				}
			},
			{ "kind": "t3", "data":
				{
					"domain": "self.golang", "banned_by": null, "media_embed": {}, "subreddit": "golang", "selftext_html": "&lt;!-- SC_OFF --&gt;&lt;div class=\"md\"&gt;&lt;p&gt;Hello champs!&lt;/p&gt;\n\n&lt;p&gt;Again with a short-question&lt;/p&gt;\n\n&lt;p&gt;As I&amp;#39;m a ruby/RoR/sinatra api developer I would like to convinced my CTO to move some of the stack to golang.&lt;/p&gt;\n\n&lt;p&gt;I know some of the webframework in golang&lt;/p&gt;\n\n&lt;ul&gt;\n&lt;li&gt;revel&lt;/li&gt;\n&lt;li&gt;martini&lt;/li&gt;\n&lt;li&gt;negroni&lt;/li&gt;\n&lt;li&gt;goji&lt;/li&gt;\n&lt;/ul&gt;\n\n&lt;p&gt;and libraries like gorila.&lt;/p&gt;\n\n&lt;p&gt;But, always all the people seems to say: start with net/http&lt;/p&gt;\n\n&lt;p&gt;Totally agree, but I found documentation a little bit random :)&lt;/p&gt;\n\n&lt;p&gt;So, my CTO ask me:&lt;/p&gt;\n\n&lt;ul&gt;\n&lt;li&gt;how stable are the api&lt;/li&gt;\n&lt;li&gt;how hot and active are the frameworks&lt;/li&gt;\n&lt;li&gt;if I use net/http do I have all the rack, sinatra, API libraries&lt;/li&gt;\n&lt;/ul&gt;\n\n&lt;p&gt;There any long save tutorial for net/http (json response, json parse, auth, REST, etc)?&lt;/p&gt;\n\n&lt;p&gt;all the best&lt;/p&gt;\n&lt;/div&gt;&lt;!-- SC_ON --&gt;", "selftext": "Hello champs!\n\nAgain with a short-question\n\nAs I'm a ruby/RoR/sinatra api developer I would like to convinced my CTO to move some of the stack to golang.\n\nI know some of the webframework in golang\n\n* revel\n* martini\n* negroni\n* goji\n\nand libraries like gorila.\n\nBut, always all the people seems to say: start with net/http\n\nTotally agree, but I found documentation a little bit random :)\n\nSo, my CTO ask me:\n\n* how stable are the api\n* how hot and active are the frameworks\n* if I use net/http do I have all the rack, sinatra, API libraries\n\n\nThere any long save tutorial for net/http (json response, json parse, auth, REST, etc)?\n\nall the best", "likes": null, "secure_media": null, "link_flair_text": null, "id": "26jiyc", "gilded": 0, "secure_media_embed": {}, "clicked": false, "stickied": false, "author": "claudiug1", "media": null, "score": 10, "approved_by": null, "over_18": false, "hidden": false, "thumbnail": "", "subreddit_id": "t5_2rc7j", "edited": false, "link_flair_css_class": null, "author_flair_css_class": null, "downs": 5, "saved": false, "is_self": true,
					"permalink": "/r/golang/comments/26jiyc/where_to_look_to_start_with_golang_web/", "name": "t3_26jiyc", "created": 1401160182.0,
					"url": "http://www.reddit.com/r/golang/comments/26jiyc/where_to_look_to_start_with_golang_web/", "author_flair_text": null,
					"title": "where to look to start with golang web", "created_utc": 1401131382.0, "ups": 15, "num_comments": 9, "visited": false, "num_reports": null, "distinguished": null
				}
			},
			...
*/
// 保存回归JSON数据的对象,请注意这个结构需要和回传数据结构相同
type Entry struct {
	Title     string
	Author    string
	URL       string
	Permalink string
}

// 返回的结构组
type Feed struct {
	Data struct {
		Children []struct {
			Data Entry
		}
	}
}

func main() {
	// 请求的网站
	url := "http://www.reddit.com/r/golang/new.json"
	// 获取Url
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("Error fetching:", err)
	}
	// 延迟关闭响应
	defer resp.Body.Close()

	// 检查我们是否得到OK
	if resp.StatusCode != http.StatusOK {
		log.Fatalln("Error Status not OK:", resp.StatusCode)
	}

	// 若OK，则读取响应数据Body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error reading body:", err)
	}

	// 创建Feed对象，解析Json并写入
	var entries Feed
	if err := json.Unmarshal(body, &entries); err != nil {
		log.Fatalln("Error decoing JSON", err)
	}

	// 循环填充到子单元中
	for _, ed := range entries.Data.Children {
		entry := ed.Data
		log.Println(">>>")
		log.Println("Title   :", entry.Title)
		log.Println("Author  :", entry.Author)
		log.Println("URL     :", entry.URL)
		log.Printf("Comments: http://reddit.com%s \n", entry.Permalink)
	}
}
