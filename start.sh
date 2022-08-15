#!/bin/sh
# This script is function to download maxminddb
. config/maxmind.config

# DOWNLOAD DIRECTLY TO MAXMIND
wget -O geoip.tar.gz "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-City&license_key=$licensekey&suffix=tar.gz"

# EXTRACT THE DB
tar -zxvf geoip.tar.gz

# MOVE TO DATA
cp GeoLite2-City_*/GeoLite2-City.mmdb data/

# START GOPI
./gopi