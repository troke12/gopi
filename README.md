<p align="center">
  <a href="https://github.com/troke12/gopi" target="blank"><img src="https://socialify.git.ci/troke12/gopi/image?description=1&font=KoHo&forks=1&issues=1&logo=https%3A%2F%2Fcdn.discordapp.com%2Fattachments%2F874251888357441537%2F910394407277182976%2F36606a05322a3f71c8c500a03e297fe703d6e647.png&name=1&pattern=Floating%20Cogs&pulls=1&stargazers=1&theme=Dark" alt="Go" /></a>

</p>

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
