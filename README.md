# Phising Checker API

This is made specifically for discord scam link, use case would be in discord bot

## Endpoints:

- Check:
  - path: `/api/v1/check`
  - query: `url`
  - example success response:
    - status code: `200`
    - response:
      ```json
      {
        "status": 200,
        "message": "Found it, its scam",
        "data": {
          "isFound": true,
          "isPhising": true,
          "domain": "discred.gift",
          "date": "2022-03-01T05:55:05.065343466Z"
        }
      }
      ```

## Refrences:

- [Discord phising links](https://github.com/nikolaischunk/discord-phishing-links)
- [SinkingYachts Phishing Domain API](https://phish.sinking.yachts/)
- [Anti-Fish API](https://anti-fish.bitflow.dev/)
- [Phising Database](https://github.com/mitchellkrogza/Phishing.Database)

## License:

@2022 Reynadi. [MIT License](https://opensource.org/licenses/MIT)
