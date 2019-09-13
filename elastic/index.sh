#!/usr/bin/env bash
echo '=== DELETE INDEX ==='
curl -XDELETE "http://localhost:9200/noonde_local"
echo ''


echo '=== CREATE INDEX ==='
curl -XPUT "http://localhost:9200/noonde_local" -H 'Content-Type: application/json' -d'
{
  "mappings": {
    "_doc": {
      "properties": {
        "type": {
          "type": "keyword"
        },
        "search": {
          "type": "text",
          "analyzer": "my_analyzer"
        },
        "tags": {
          "type": "keyword"
        },
        "id": {
          "type": "long"
        },
        "date": {
          "type": "keyword"
        },
        "review": {
           "type": "float"
        },
        "event_types": {
           "type": "keyword"
        },
        "amenities": {
           "type": "keyword"
        },
        "latidude": {
            "type": "double"
        },
        "longitude": {
            "type": "double"
        },
        "capacity": {
            "type": "integer"
        }
      }
    }
  },
  "settings": {
    "index": {
      "number_of_shards": 1,
      "max_ngram_diff": 10
    },
    "analysis": {
      "analyzer": {
        "my_analyzer": {
          "tokenizer": "my_tokenizer",
          "filter": [
            "lowercase"
          ]
        }
      },
      "tokenizer": {
        "my_tokenizer": {
          "type": "ngram",
          "min_gram": 1,
          "max_gram": 10,
          "token_chars": [
            "letter",
            "digit",
            "symbol",
            "punctuation"
          ]
        }
      }
    }
  }
}'
echo ''