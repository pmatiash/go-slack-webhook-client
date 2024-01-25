# go-slack-webhook-client
Package could be useful for sending messages to a Slack webhook

### Example
```
{
	"blocks": [
		{
			"type": "section",
			"text": {
				"type": "plain_text",
				"text": "This is a plain text section block.",
				"emoji": true
			}
		},
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": "This is a mrkdwn section block :ghost: *this is bold*, and ~this is crossed out~, and <https://google.com|this is a link>"
			}
		},
		{
			"type": "section",
			"fields": [
				{
					"type": "plain_text",
					"text": "*this is plain_text text*",
					"emoji": true
				},
				{
					"type": "plain_text",
					"text": "*this is plain_text text*",
					"emoji": true
				},
				{
					"type": "plain_text",
					"text": "*this is plain_text text*",
					"emoji": true
				},
				{
					"type": "plain_text",
					"text": "*this is plain_text text*",
					"emoji": true
				},
				{
					"type": "plain_text",
					"text": "*this is plain_text text*",
					"emoji": true
				}
			]
		}
	]
}
```