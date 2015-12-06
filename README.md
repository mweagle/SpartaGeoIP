# SpartaGeoIP
Simple [Sparta](http://gosparta.io) based Request IP -> Geo mapping example

Inspired by https://github.com/tmaiaroto/go-lambda-geoip

Instructions:

  1. `make get`
  1. `S3_BUCKET=<MY_S3_BUCKET_NAME> make provision`
  1. Query curl -vs https://**REST_API_ID**.execute-api.**AWS_REGION**.amazonaws.com/ipgeo/info

Sample pretty printed HTTP response body:

```json
{
  "code": 200,
  "status": "OK",
  "headers": {
    "content-type": "application/json",
    "date": "Sun, 06 Dec 2015 17:50:15 GMT",
    "content-length": "984"
  },
  "results": {
    "info": {
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
}
```

## TODO
  - Add to [API Gateway Docs](http://gosparta.io/docs/apigateway/)
