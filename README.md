# gopi
Simple API to get information from your IP Address

# Idea
This idea come from [IP zxq](https://ip.zxq.co) and literaly i clone it

# How to download GeoIP2 ?
Remember to change `YOUR_LICENSE_KEY` , you can obtain [here](https://www.maxmind.com/en/account)

```bash
wget -O geoip.tar.gz "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-City&license_key=YOUR_LICENSE_KEY&suffix=tar.gz"
tar -zxvf geoip.tar.gz
mv geoip.tar.gz data
```

# Routes
- `/` - get your current information
- `/8.8.8.8/country` - get the country iso code
- `/8.8.8.8` - get the information another ip

# Result

```json
{
    "ip":"127.0.0.1",
    "city":"Jakarta",
    "region":"Jakarta",
    "country":"ID",
    "country_full":"Indonesia",
    "continent":"AS",
    "continent_full":"Asia",
    "loc":"-6.2092,106.8200",
    "postal":""
}
```