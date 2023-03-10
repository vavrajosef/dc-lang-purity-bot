# dc-lang-purity-bot

If, like me, you are bothered by mixing two different languages in one sentence, this discord bot is for you. Based on a request from any user of the discord server, you can ask the bot to translate a message. The bot expects a source language and a target language as input. In the channel where the translation request was registered, the bot will then send a message in which the expressions that the bot has evaluated as source language words will be translated into the target language.

## Usage

Starting the bot is trivial. All you need is a token.
The bot is then launched as follows:

`dc-lang-purity-bot -t <YOUR_TOKEN>`

The bot will run in the terminal until it registers an interrupt command. Any errors, whether they are translation errors or incorrectly typed requests, will be listed on the stdout. Discord message will only contain translated words or "Nothing to translate".

## Translation request

When you need translation for certain message just write response as followed:
`translate source-language target-language`
For languages specifications use ISO standard names or RFC3066.
Example:
`translate en cs` 
This will translate all english words in message to czech and send message starting with asterisk.


