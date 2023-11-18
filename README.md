# Tekken Backend

This is an API which allows you to request Tekken 8 character and move data.

## Implemented:

✅ `/api/help` directs to the help page.

✅ `/api/characters` returns data for all characters.

```json
[
  {
    "Id": 1,
    "ShortName": "Jin",
    "LongName": "Jin Kazama",
    "FightingStyle": "Karate",
    "Nationality": "Japan",
    "Height": 183,
    "Weight": 87,
    "Gender": "m"
  },
  {
    "Id": 2,
    "ShortName": "Asuka",
    "LongName": "Asuka Kazama",
    "FightingStyle": "Kazama Style Traditional Martial Arts",
    "Nationality": "Japan",
    "Height": 167,
    "Weight": 57,
    "Gender": "f"
  },
  {
    "etc": "etc"
  }
]
```

## To Do:

❌ `/api/character?=bryan` returns data for the specified character.

```json
[
  {
    "Id": 13,
    "ShortName": "Bryan",
    "LongName": "Bryan Fury",
    "FightingStyle": "Kickboxing",
    "Nationality": "United States",
    "Height": 186,
    "Weight": 87,
    "Gender": "m"
  }
]
```

❌ `/api/movelist` returns moves for all characters.

❌ `/api/movelist?=brayn` returns moves for the specified character.
