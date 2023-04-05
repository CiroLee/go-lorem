package lorem

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/CiroLee/gear/gearmap"
	"github.com/CiroLee/go-lorem/data"
)

type AddressBase struct {
	Code string
	Name string
}

// return a random Chinese Mainland province name and code
func Province() AddressBase {
	var provinces []AddressBase
	for _, p := range data.AddressDict {
		var item AddressBase = AddressBase{
			Code: p.Code,
			Name: p.Name,
		}
		provinces = append(provinces, item)
	}
	return randomElement(provinces)
}

// return a random Chinese Mainland city name and code
func City() AddressBase {
	var province = Province()
	return randomCity(province.Code)
}

// return a random Chinese Mainland county name and code
func County() AddressBase {
	var province = Province()
	var city = randomCity(province.Code)
	return randomCounty(province.Code, city.Code)
}

// return a random Chinese Mainland address string that include province, city and county, separated by space
func Address() string {
	var province = Province()
	var city = randomCity(province.Code)
	var county = randomCounty(province.Code, city.Code)

	var address = province.Name + " " + city.Name + " " + county.Name
	return strings.TrimRight(address, " ")
}

// return a random Chinese Mainland zip code
func ZipCode() int {
	var zip, _ = IntBy(6, true)
	return zip
}

// return a random longitude and latitude array. support deg format and dms format
func LongAndLat(format string) [2]string {
	longitude, _ := randomFloat(-100, 100, 4)
	latitude, _ := randomFloat(-90, 90, 4)

	if format == "dms" {
		return [2]string{degToDms(longitude), degToDms(latitude)}
	}
	return [2]string{strconv.FormatFloat(longitude, 'f', 4, 64) + "°", strconv.FormatFloat(latitude, 'f', 4, 64) + "°"}
}

func degToDms(num float64) string {
	number := math.Abs(num)
	deg := math.Floor(number)
	minute := math.Floor((number - deg) * 60)
	second := ((number-deg)*60 - minute) * 60
	second, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", second), 64)
	str := strconv.FormatFloat(deg, 'f', 0, 64) + "°" + strconv.FormatFloat(minute, 'f', 0, 64) + "′" + strconv.FormatFloat(second, 'f', 0, 64) + "″"
	if num < 0 {
		return "-" + str
	}
	return str
}

func randomCity(provinceCode string) AddressBase {
	var _cities []data.Districts = gearmap.Values(data.AddressDict[provinceCode].Cities)
	var cities []AddressBase
	for _, c := range _cities {
		var item AddressBase = AddressBase{
			Code: c.Code,
			Name: c.Name,
		}

		cities = append(cities, item)
	}
	return randomElement(cities)
}

func randomCounty(provinceCode, cityCode string) AddressBase {
	var _counties map[string]string = data.AddressDict[provinceCode].Cities[cityCode].Districts
	// 部分省市没有区
	if len(_counties) == 0 {
		return AddressBase{
			Code: "",
			Name: "",
		}
	}
	var counties []AddressBase
	for k, v := range _counties {
		counties = append(counties, AddressBase{
			Code: k,
			Name: v,
		})
	}

	return randomElement(counties)
}
