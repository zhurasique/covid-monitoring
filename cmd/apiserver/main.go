package main

import (
	"covid-monitoring/internal/app/apiserver"
	"covid-monitoring/internal/app/model"
	"covid-monitoring/internal/app/store"
	"database/sql"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"time"
)

var(
	configPath string
	storePath string
)

func init()  {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
	flag.StringVar(&storePath, "store-path", "configs/store.toml", "path to store file")
}

func getData() {
	//var countries [228]string
	//countriesLength := 0
	//
	//var records [685]string
	//recordsLength := 0
	//
	//var parsedRecords [684]int64
	//
	//var cases [228]int64
	//var deaths [228]int64
	//var recovered [228]int64
	//count := 0
	//
	//c := colly.NewCollector()
	//
	//c.OnHTML("table tr th[scope] a[href][title]", func(e *colly.HTMLElement) {
	//	if countriesLength < 1 {
	//		countries[countriesLength] = e.Text
	//		countriesLength++
	//	}else {
	//		if countries[countriesLength - 1] != "Tanzania" {
	//			countries[countriesLength] = e.Text
	//			countriesLength++
	//		}
	//	}
	//})
	//
	//c.OnHTML("table tr td", func(e *colly.HTMLElement) {
	//	match, _ := regexp.MatchString("^\\[", e.Text)
	//	if !match {
	//		if recordsLength < 1 {
	//			records[recordsLength] = e.Text
	//			recordsLength++
	//		}else {
	//			xmatch, _ := regexp.MatchString("History of deaths", records[recordsLength - 1])
	//			if !xmatch {
	//				records[recordsLength] = e.Text
	//				recordsLength++
	//			}
	//		}
	//	}
	//})
	//
	//c.Visit("https://en.wikipedia.org/wiki/Template:COVID-19_pandemic_data")
	//
	//
	//for i := 0; i < len(parsedRecords); i++{
	//	records[i] = strings.TrimSuffix(records[i], "\n")
	//
	//	st := strings.Replace(records[i], "\n", "", -1)
	//	st = strings.Replace(records[i], ",", "", -1)
	//
	//	var intBase int64
	//	var err error
	//	if st == "No data"{
	//		intBase = -1
	//	}else {
	//		intBase, err = strconv.ParseInt(st, 0, 64)
	//		if err != nil {
	//			panic(err)
	//		}
	//	}
	//
	//	parsedRecords[i] = intBase
	//}
	//
	//count = 0
	//for i := 0; i < len(parsedRecords); i++ {
	//	if i + 2 <= len(parsedRecords) {
	//		cases[count] = parsedRecords[i]
	//		i += 2
	//		count++
	//	}
	//}
	//
	//count = 0
	//for i := 1; i < len(parsedRecords); i++ {
	//	if i + 2 <= len(parsedRecords) {
	//		deaths[count] = parsedRecords[i]
	//		i += 2
	//		count++
	//	}
	//}
	//
	//count = 0
	//for i := 2; i < len(parsedRecords); i++ {
	//	if i + 2 <= len(parsedRecords) {
	//		recovered[count] = parsedRecords[i]
	//		i += 2
	//		count++
	//	}else{
	//		recovered[len(recovered) - 1] = parsedRecords[len(parsedRecords) - 1]
	//	}
	//}
	//
	////for i := 0; i < len(deaths); i++ {
	////	fmt.Println(countries[i], " ", cases[i], " ", deaths[i], " ", recovered[i])
	////}
	//
	//data := model.Data{
	//	ID:        0,
	//	Country:   countries[0],
	//	Cases:     cases[0],
	//	Deaths:    deaths[0],
	//	Recovered: recovered[0],
	//}
	//
	//fmt.Println(data.ID, data.Country, data.Cases, data.Deaths, data.Recovered)



	config := store.NewConfig()

	_, err := toml.DecodeFile(storePath, config)
	if err != nil{
		log.Fatal(err)
	}

	st := store.New(config)
	if err := st.Open(); err != nil{
		panic(err)
	}

	data := model.Data{}

	sqlStatement := `SELECT * FROM data WHERE id=$1;`
	row := st.Db.QueryRow(sqlStatement, 1)
	err = row.Scan(&data.ID, &data.Country, &data.Cases, &data.Deaths, &data.Recovered)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println(data)
	default:
		panic(err)
	}
}

func updateData() {
	for {
		time.Sleep(2 * time.Second)
		go getData()
	}
}


func main()  {
	go updateData()

	flag.Parse()

	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil{
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}