# UFC Web Scraper

This code extracts data about all fighters from UFC Website (https://ufc.com). The Data will be exported in new .json file in current directory and will look like this:

```JSON
{
    "Fighters": [
        {
            "name": "Robert Whittaker",
            "nickName": "\"The Reaper\"",
            "division": 5,
            "status": "Active",
            "hometown": "Otahuhu, Australia",
            "trainsAt": "PMA, Padstow NSW, Australia",
            "fightingStyle": "Brazilian Jiu-Jitsu",
            "age": 32,
            "height": 72,
            "weight": 196,
            "octagonDebut": "Dec. 16, 2012",
            "reach": 73.50,
            "legReach": 43.00,
            "wins": 25,
            "loses": 7,
            "draw": 0,
            "fighterUrl": "https://www.ufc.com/athlete/danny-abbadi",
            "imageUrl": "https://some.url
            "stats": {
                "totalSigStrLandned": 1258,
                "totalSigStrAttempted": 2981,
                "strAccuracy": 42,
                "totalTkdLanded": 4,
                "totalTkdAttempted": 41,
                "tkdAccuracy": 9,
                "sigStrLanded": 4.47, // per min.
                "sigStrAbs": 3.39, // per min.
                "sigStrDefense": 60,
                "takedownDefense": 83,
                "takedownAvg": 0.85, // per 15 min.
                "submissionAvg": 0, // per 15 min.
                "knockdownAvg": 0.53,
                "avgFightTime": "14:04",
                "winByKO": 10,
                "winBySub": 5,
                "winByDec": 9
            }
        },
        // . . .
}
```

# Division ids:

| Division            | Value |
| ------------------- | ----- |
| Flyweight           | 0     |
| Bantamweight        | 1     |
| Featherweight       | 2     |
| Lightweight         | 3     |
| Welterweight        | 4     |
| Middleweight        | 5     |
| Lightheavyweight    | 6     |
| Heavyweight         | 7     |
| WomensStrawweight   | 8     |
| WomensFlyweight     | 9     |
| WomensBantamweight  | 10    |
| WomensFeatherweight | 11    |

## Usage

1. Clone repository
   `git clone https://github.com/DoRightt/ufc-scrapper.git`

2. Install dependencies
   `go mod tidy`

3. Just run the script and wait for the scraper to get data from all fighters and pages. This may take quite a long time.

Logger will create a log.json file in logger/log.json

Also, if ufc.com changes its html structure, the scraper may not work properly.

Special thanks to [eneiromatos](https://github.com/eneiromatos) for example of his [ufc-web-scrapper](https://github.com/eneiromatos/UFC-Fighters-Scraper)
