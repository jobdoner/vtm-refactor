package migrations

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
	"github.com/gocarina/gocsv"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func Migrate(uri string, migrationStatus int, redisConnURI string) {
	CustomAffinity := make(map[string]string) //make([]string,6748)
	ContentCategory := make(map[string]string)
	BidStrategy := make(map[string]string)  //make([]string,10)
	BudgetType := make(map[string]string)   //make([]string,2)
	CampaignType := make(map[string]string) //make([]string,12)
	Ages := make(map[string]string)         //make([]string,7)
	Gender := make(map[string]string)       // make([]string,2)
	location := make(map[string]string)
	Placement := make(map[string]string)
	Audience := make(map[string]string)
	Topics := make(map[string]string)
	Language := make(map[string]string)

	ReadMigrationsFileCC(Audience, "vtm-refactor/migrations/redis/Audience")

	ReadMigrationsFileCC(Placement, "vtm-refactor/migrations/redis/Placement")

	ReadMigrationsFileCC(Topics, "vtm-refactor/migrations/redis/Topics")

	ReadMigrationsFileCC(Language, "vtm-refactor/migrations/redis/Language")

	ReadMigrationsFile(CustomAffinity, "vtm-refactor/migrations/redis/CustomAffinity")
	ReadMigrationsFileCC(ContentCategory, "vtm-refactor/migrations/redis/ContentCategory")
	//ReadMigrationsFileCC(location, "vtm-refactor/migrations/redis/Location")

	ReadMigrationsFile(BidStrategy, "vtm-refactor/migrations/redis/BidStrategy")
	ReadMigrationsFile(BudgetType, "vtm-refactor/migrations/redis/BudgetType")
	ReadMigrationsFile(CampaignType, "vtm-refactor/migrations/redis/CampaignType")
	ReadMigrationsFile(Ages, "vtm-refactor/migrations/redis/Ages")
	ReadMigrationsFile(Gender, "vtm-refactor/migrations/redis/Gender")
	ReadMigrationsCSV(location, "vtm-refactor/migrations/redis/Location.csv")

	db, err := sql.Open("postgres", uri)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "vtm-refactor/migrations/redis",
	}

	i, err := migrate.Exec(db, "postgres", migrations, migrate.MigrationDirection(migrationStatus))

	fmt.Println(i, err)

	opt, err := redis.ParseURL(redisConnURI)
	redisConn := redis.NewClient(opt)
	if redisConn.Ping(context.Background()).Err() != nil {
		panic(redisConn.Ping(context.Background()).Err())
	}

	for i := 0; i < 75; i++ {
		err := redisConn.HSet(context.Background(), "videos:limit", strconv.Itoa(i)+"_1", "1").Err()
		err = redisConn.HSet(context.Background(), "videos:limit", strconv.Itoa(i)+"_2", "2").Err()
		err = redisConn.HSet(context.Background(), "videos:limit", strconv.Itoa(i)+"_3", "3").Err()
		err = redisConn.HSet(context.Background(), "videos:limit", strconv.Itoa(i)+"_4", "4").Err()
		err = redisConn.HSet(context.Background(), "videos:limit", strconv.Itoa(i)+"_5", "5").Err()
		err = redisConn.HSet(context.Background(), "videos:limit", strconv.Itoa(i)+"_6", "6").Err()
		if err != nil {
			fmt.Println(err)
		}
	}

	defer redisConn.Close()

	//  Topic
	//  Placement
	//  Location
	//  Languages
	//	BudgetType /d/
	//	CampaignType /d/
	//	Networks /d/
	//  Frequency /c/ /d/
	//	BLS SPLIT /D/

	redisConn.Del(context.Background(), "adgroup:CustomAffinity",
		"adgroup:AudienceInterests",
		"adgroup:BidStrategy",
		"adgroup:CampaignType",
		"adgroup:BudgetType",
		"adgroup:Age",
		"adgroup:Gender",
		"adgroup:location",
		"adgroup:settings",
		"adgroup:audience",
		"adgroup:placement",
		"adgroup:topics",
		"adgroup:language",
		"adgroup:languages",
	)

	redisConn.Del(context.Background(), "adgroup:customAffinity",
		"adgroup:audienceInterests",
		"adgroup:bidStrategy",
		"adgroup:campaignType",
		"adgroup:budgetType",
		"adgroup:age",
		"adgroup:gender",
		"adgroup:location",
		"adgroup:settings",
		"adgroup:audience",
		"adgroup:placement",
		"adgroup:topics",
		"adgroup:language",
		"adgroup:languages",
	)

	if ok, err := redisConn.HGetAll(context.Background(), "adgroup:customAffinity").Result(); len(ok) == 0 && err == nil {
		for key, value := range CustomAffinity {
			redisConn.HSet(context.Background(), "adgroup:customAf`finity", key, value)
		}
	}
	if ok, err := redisConn.HGetAll(context.Background(), "adgroup:audienceInterests").Result(); len(ok) < 1 && err == nil {
		for key, value := range ContentCategory {
			redisConn.HSet(context.Background(), "adgroup:audienceInterests", key, value)
		}
	}

	if ok, err := redisConn.HGetAll(context.Background(), "adgroup:bidStrategy").Result(); len(ok) == 0 && err == nil {
		for key, value := range BidStrategy {

			redisConn.HSet(context.Background(), "adgroup:bidStrategy", key, value)
		}
	}

	if ok, err := redisConn.HGetAll(context.Background(), "adgroup:campaignType").Result(); len(ok) == 0 && err == nil {
		for key, value := range CampaignType {
			redisConn.HSet(context.Background(), "adgroup:campaignType", key, value)
		}
	}

	if ok, err := redisConn.HGetAll(context.Background(), "adgroup:budgetType").Result(); len(ok) == 0 && err == nil {
		for key, value := range BudgetType {
			redisConn.HSet(context.Background(), "adgroup:budgetType", key, value)
		}

	}

	if ok, err := redisConn.HGetAll(context.Background(), "adgroup:age").Result(); len(ok) == 0 && err == nil {
		for key, value := range Ages {
			redisConn.HSet(context.Background(), "adgroup:age", key, value)
		}

	}

	if ok, err := redisConn.HGetAll(context.Background(), "adgroup:gender").Result(); len(ok) == 0 && err == nil {
		for key, value := range Gender {
			redisConn.HSet(context.Background(), "adgroup:gender", key, value)
		}

	}

	if ok, err := redisConn.HGetAll(context.Background(), "adgroup:settings").Result(); len(ok) == 0 && err == nil {
		for key, value := range AdGroupsParameters {
			redisConn.HSet(context.Background(), "adgroup:settings", key, value)
		}
	}

	if ok, err := redisConn.HGetAll(context.Background(), "adgroup:location").Result(); len(ok) < 1 && err == nil {
		for key, value := range location {
			redisConn.HSet(context.Background(), "adgroup:location", key, value)
		}
	}

	if ok, err := redisConn.HGetAll(context.Background(), "adgroup:audience").Result(); len(ok) < 1 && err == nil {
		for key, value := range Audience {
			redisConn.HSet(context.Background(), "adgroup:audience", key, value)
		}
	}
	if ok, err := redisConn.HGetAll(context.Background(), "adgroup:placement").Result(); len(ok) < 1 && err == nil {
		for key, value := range Placement {
			redisConn.HSet(context.Background(), "adgroup:placement", key, value)
		}
	}
	if ok, err := redisConn.HGetAll(context.Background(), "adgroup:topics").Result(); len(ok) < 1 && err == nil {
		for key, value := range Topics {
			redisConn.HSet(context.Background(), "adgroup:topics", key, value)
		}
	}
	if ok, err := redisConn.HGetAll(context.Background(), "adgroup:languages").Result(); len(ok) < 1 && err == nil {
		for key, value := range Language {
			redisConn.HSet(context.Background(), "adgroup:languages", key, value)
		}
	}

}
func ReadMigrationsFile(str map[string]string, filename string) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		panic(err)
	}
	fileReader := bufio.NewReader(file)
	defer file.Close()
	i := 1
	for {
		s, err := fileReader.ReadString('\n')
		if len(s) > 2 {
			str[strconv.Itoa(i)] = s[:len(s)-1]
		}
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		i++

	}
	fmt.Println(i)

}
func ReadMigrationsFileCC(str map[string]string, filename string) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		panic(err)
	}
	fileReader := bufio.NewReader(file)
	defer file.Close()
	i := 0
	for {
		s, err := fileReader.ReadString('\n')
		newS := strings.Split(s, "\t")

		if len(s) > 2 {
			str[newS[1][0:len(newS[1])-1]] = newS[0]
		}
		//else {
		//	str[newS[1][0:len(newS[1])]] = newS[0]
		//}
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		i++

	}
	fmt.Println(i)

}

type LocationCSV struct {
	CriteriaID    string `csv:"Criteria ID"`
	Name          string `csv:"Name"`
	CanonicalName string `csv:"Canonical Name"`
	ParentID      string `csv:"Parent ID"`
	CountryCode   string `csv:"Country Code"`
	TargetType    string `csv:"Target Type"`
}

func ReadMigrationsCSV(str map[string]string, filename string) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		panic(err)
	}

	defer file.Close()
	newLocationCSV := []LocationCSV{}
	err = gocsv.UnmarshalFile(file, &newLocationCSV)
	if err != nil {
		fmt.Println(err)
	}
	for _, value := range newLocationCSV {

		if (value.CountryCode == "RU" || value.CountryCode == "IND") && value.TargetType == "City" {
			str[value.CriteriaID] = value.Name
		}

	}
	fmt.Println(len(newLocationCSV))

}
