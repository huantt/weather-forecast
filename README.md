## Today's Weather
<div align="center">

`Hanoi, Vietnam - 03/12/2023`

<img src="https://cdn.weatherapi.com/weather/64x64/day/122.png"/>

Overcast

</div>


<table>
    <tr>
        <th>Hour</th>
        
    </tr>
    <tr>
        <th>Weather</th>
        
    </tr>
    <tr>
        <th>Condition</th>
        
    </tr>
    <tr>
        <th>Temperature</th>
        
    </tr>
    <tr>
        <th>Wind</th>
        
    </tr>
</table>


## Weather For Next 3 days


<table>
    <tr>
        <th>Date</th>
        <td>03/12/2023</td><td>04/12/2023</td><td>05/12/2023</td>
    </tr>
    <tr>
        <th>Weather</th>
        <td><img src="https://cdn.weatherapi.com/weather/64x64/day/122.png"/></td><td><img src="https://cdn.weatherapi.com/weather/64x64/day/122.png"/></td><td><img src="https://cdn.weatherapi.com/weather/64x64/day/122.png"/></td>
    </tr>
    <tr>
        <th>Condition</th>
        <td width="200px">Overcast</td><td width="200px">Overcast</td><td width="200px">Overcast</td>
    </tr>
    <tr>
        <th>Temperature</th>
        <td>17.9 -  20.2 °C</td><td>18.8 -  21.1 °C</td><td>18.5 -  21.9 °C</td>
    </tr>
    <tr>
        <th>Wind</th>
        <td>8.3 kph</td><td>13 kph</td><td>7.9 kph</td>
    </tr>
</table>


*Updated at: 2023-12-02T19:17:27Z*

## GitHub Actions: Embed up-to-date Weather in your README
<details>
<summary>
    View
</summary>

You can easily embed tables in your README.md using GitHub Actions by following these simple steps:

**Step 1:** In your repository, create a file named `README.md.template`.

**Step 2:** Write anything you want within the `README.md.template` file.

**Step 3:** Embed one of the following entities within your `README.md.template`:

- **Today's Weather Table:**
```shell
{{ template "hourly-table" $.TodayWeather.HourlyWeathers }}
```

- **Daily Weather Table:**
```shell
{{ template "daily-table" .Weathers }}
```

- **Updated at:**
```shell
{{ formatTime .UpdatedAt }}
```

If you are familiar with Go templates, you have access to the `root` variable, which includes the following fields:

- `Weathers`: An array of daily Weather. You can view the Weather struct definition in [model/weather.go](model/weather.go).
- `UpdatedAt`: This field contains the timestamp in the format of `time.Date`.

**Step 4**: Register Github Action
- Create a file `.github/workflows/update-weather.yml` in your repository.
```yml
name: "Cronjob"
on:
schedule:
- cron: '15 * * * *'

jobs:
    update-weather:
        permissions: write-all
        runs-on: ubuntu-latest
        steps:
            - uses: actions/checkout@v3
            - name: Generate README
              uses: huantt/weather-forecast@v1.0.5
              with:
                city: HaNoi
                days: 7
                weather-api-key: ${{ secrets.WEATHER_API_KEY }}
                template-file: 'README.md.template'
                out-file: 'README.md'
            - name: Commit
              run: |
                if git diff --exit-code; then
                  echo "No changes to commit."
                  exit 0
                else
                  git config user.name github-actions
                  git config user.email github-actions@github.com
                  git add .
                  git commit -m "update"
                  git push origin main
                fi
```
- Update some variable in this file:
    - city: The city that you want to forecast weather
    - days: number of forecast days
    - template-file: Path to the above template file. Eg. `template/README.md.template`
    - out-file: your README.md file name
    - weather-api-key:
        - Register free API key in [https://weatherapi.com](https://weatherapi.com)
        - Setup secrets with name `WEATHER_API_KEY` in `Your repo > settings > Secrets and variables > Actions > New repository secret`

**Step 5**: Commit your change, then Github actions will run as your specificed cron to update Weather into your README.md file
</details>


## Usage
<details>
<summary>View</summary>

#### Install
```shell
go install https://github.com/huantt/weather-forecast
```

#### Run

```shell
Usage:
weather-forecast update-weather [flags]

Flags:
--city string              City
--days int                 Days of forecast (default 7)
-h, --help                     help for update-weather
-o, --out-file string          Output file path
-f, --template-file string     Readme template file path
-k, --weather-api-key string   weatherapi.com API key

```

**Sample**
```shell
weather-forecast update-weather \
--days=7 \
--weather-api-key="$WEATHER_API_KEY" \
--template-file='template/README.md.template' \
--city=HaNoi \
--out-file='README.md'
```

### Docker
```shell
docker build -t weather-forecast .
```

```shell
docker run --rm \
-v ./:/app/data \
weather-forecast \
--weather-api-key='XXXX' \
--city=HaNoi \
--out-file=data/README.md \
--template-file=data/README.md.template
```

</details>
