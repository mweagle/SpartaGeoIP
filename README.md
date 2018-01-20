# SpartaGeoIP
Simple [Sparta](http://gosparta.io) based Request IP -> Geo mapping example

Original version: https://github.com/tmaiaroto/go-lambda-geoip

Instructions:

  1. `make get`
  1. `S3_BUCKET=<MY_S3_BUCKET_NAME> make provision`
  1. `curl -vs https://**REST_API_ID**.execute-api.**AWS_REGION**.amazonaws.com/ipgeo/info`

Sample pretty-printed (via [jq](https://stedolan.github.io/jq/) response:

```json
{
  "ip": "73.118.138.121",
  "record": {
    "City": {
      "GeoNameID": 0,
      "Names": null
    },
    "Continent": {
      "Code": "NA",
      "GeoNameID": 6255149,
      "Names": {
        "de": "Nordamerika",
        "en": "North America",
        "es": "Norteamérica",
        "fr": "Amérique du Nord",
        "ja": "北アメリカ",
        "pt-BR": "América do Norte",
        "ru": "Северная Америка",
        "zh-CN": "北美洲"
      }
    },
    "Country": {
      "GeoNameID": 6252001,
      "IsoCode": "US",
      "Names": {
        "de": "USA",
        "en": "United States",
        "es": "Estados Unidos",
        "fr": "États-Unis",
        "ja": "アメリカ合衆国",
        "pt-BR": "Estados Unidos",
        "ru": "США",
        "zh-CN": "美国"
      }
    },
    "Location": {
      "AccuracyRadius": 0,
      "Latitude": 0,
      "Longitude": 0,
      "MetroCode": 0,
      "TimeZone": ""
    },
    "Postal": {
      "Code": ""
    },
    "RegisteredCountry": {
      "GeoNameID": 6252001,
      "IsoCode": "US",
      "Names": {
        "de": "USA",
        "en": "United States",
        "es": "Estados Unidos",
        "fr": "États-Unis",
        "ja": "アメリカ合衆国",
        "pt-BR": "Estados Unidos",
        "ru": "США",
        "zh-CN": "美国"
      }
    },
    "RepresentedCountry": {
      "GeoNameID": 0,
      "IsoCode": "",
      "Names": null,
      "Type": ""
    },
    "Subdivisions": null,
    "Traits": {
      "IsAnonymousProxy": false,
      "IsSatelliteProvider": false
    }
  }
}
```
