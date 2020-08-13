# searchcli

A command line tool that, given a set of tokens, can quickly search a large document to find the location of matches. For example, given a JSON file that has characters from the Lord of the Rings novels:

    searchcli -f characters.json < /path/to/text/lotr.txt

where `tokens.json` is in the form:
    
    [
      {
        "id": 1,
        "label": "aragorn",
        "category": "character"
      },
      {
        "id": 2,
        "label": "frodo",
        "category": "character"
      },
      .
      .
      .
    ]

could return something like:

    [
      {
        "token": {
        "id": 1,
        "label": "aragorn",
        "category": "character"
      },
      "start_at": 918,
      "end_at": 925
    },
    {
      "token": {
        "id": 2,
        "label": "frodo",
        "category": "character"
      },
      "start_at": 1218,
      "end_at": 1223
    },
    .
    .
    .
  }
