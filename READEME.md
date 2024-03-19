url for NOAA tides:
https://api.tidesandcurrents.noaa.gov/api/prod/datagetter?begin_date=20240313&end_date=20240319&station=9415339&product=predictions&datum=MLLW&time_zone=lst_ldt&interval=hilo&units=english&application=DataAPI_Sample&format=json

url for NOAA weather forcast:
https://api.weather.gov

Marshall, CA GEOLOCATION lat/long
38.16213,-122.89357

api "points" endpoint for Inverness based on Marshall Geolocale
https://api.weather.gov/points/38.1621,-122.8935

  forecastGridData url (https://api.weather.gov/gridpoints/MTR/72,126) gives you

api forecast grid data, hourly
https://api.weather.gov/gridpoints/MTR/72,126/forecast/hourly

api endpoint for non-hourly forecast (i.e. "this afternoon", "tonight", "thursday", etc...)
https://api.weather.gov/gridpoints/MTR/72,126/forecast

so for Marshall, the Inverness station is office "MTR" and GridX/GridY are 72, 126