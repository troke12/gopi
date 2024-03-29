
![gopi](https://github.com/troke12/gopi/assets/10250068/89ceea57-f91d-4aee-8eb7-b92acc804727)

<h2 align="center"><a href="https://ip.datenshi.pw">Demo</a></h1>

## Idea
This idea come from [IP zxq](https://ip.zxq.co) and literaly i clone it, also already used for my project.

## Download GeoIP2Lite Database
First of all, we need to download the database from Maxmind and below you can use this to download the db and remember to change `YOUR_LICENSE_KEY` 

You can obtain LICENSE KEY [here](https://www.maxmind.com/en/account)

```bash
wget -O geoip.tar.gz "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-City&license_key=YOUR_LICENSE_KEY&suffix=tar.gz"
tar -zxvf geoip.tar.gz
mv geoip.tar.gz data
```

## Development

```
git clone https://github.com/troke12/gopi
cd gopi
cp .env.example .env
go build
./gopi
```

## Docker
If you want to use docker just add some build-arg in build docker
```
docker build --build-arg LICENSEY_KEY=LICENSEY_KEY --build-arg API_FGIP=API_FREEGEOIP --build-arg PORT="localhost:3045" --build-arg ROLLBARTOKEN=ROLLBAR_TOKEN . -t gopi

docker run -d p 3045:3045 --name gopi gopi
```
It will automatically download the latest MaxMindDB and extract it to the `data` folder

## Routes
- `/` - get your current information
- `/8.8.8.8/country` - get the country iso code
- `/8.8.8.8` - get the information another ip

## Result

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
