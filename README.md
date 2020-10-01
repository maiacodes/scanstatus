ScanStatus lets you make sure members of your Discord guild aren't putting bad things in their statuses such as offensive things or other things against the rules of the guild.

ScanStatus checks that users aren't using 'bad words' in their statuses, and if they are, an alert will be sent into a channel of your choice.

`.env` template:
```
TOKEN=
BADWORDS=
ALERTCHANNEL=
```
Notes: `TOKEN` must be a Discord Bot token. `BADWORDS` is a comma seperated list of badwords (in lower case!), for example: `bleppers,example,oink`. `ALERTCHANNEL` is the Channel ID of the channel to post alerts.