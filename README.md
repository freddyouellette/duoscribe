# DuoScribe [![Donate](https://img.shields.io/badge/Donate-fec133)](https://www.paypal.com/donate/?hosted_button_id=3PJ9XD363CC5E)
Extract text from Duolingo exported exercise images

### Input Image
![Duolingo Exported Exercise Image](https://raw.githubusercontent.com/freddyouellette/duolingo-transcriber/master/test/2_lines.jpg)
### Exported Text
```
Lo può ripetere per favore?
Can you repeat it please?
```
### Or in JSON format:
```
[
	{
		"text": "Lo può ripetere per favore?",
		"language": "it"
	},
	{
		"text": "Can you repeat it please?",
		"language": "en"
	}
]
```

## How to Use
> WARNING: This program uses AWS Rekognition and AWS Comprehend to extract text from images and get the language. You must have an AWS account set up with an access key. The Free Tier of AWS allows 5000 processed images per month on Rekognition, and $0.0001 per unit on Comprehend.
1. Build the program
```
make build
```
2. Run the program and pass your input image:
```
./bin/duoscribe <input_filepath>
```
3. If you want json, add the flag:
```
./bin/duoscribe --json <input_filepath>
```

## DuoScribe Program Structure
* [Flow Diagram](https://mermaid.live/edit#pako:eNqFU81u2zAMfhVC19XY3RgCFHUOHoLFaBLs4gtjM44aW_IkqmtR9N1HWXGTFWunm6TvT6T4ohrbksqVp1-BTEOFxs7hUBuQhQ1bBztADztPLh2O6Fg3ekTDUMWrypEnw8jaGljh87-ARQQWdkD9IaSMkNIcHHp2oeHg6CPoMkKXT0zOYA9ba3tfmwTbZYvFLpfL0ToGPWBHcHB2gCLYXpvOXmBfqhzusO_hzg4DmhZW2tCk9m3vvi7gt-bjWUHs0HVhkGcmfhX5hfCteSQxmiVuXecTWxu28_EcroisMqZjJ7WFLT1xSldGmwQqI2iZQ9otZZcJ5fbnBu7pZDujY52Tx8Q_i8nhhZ8Vkf-3aUFM4rlC04X4JnuYRAibI0hl_usubxkdHcm0E-9NJ-l-ai_bexrsI6WwhmzwU_h3IKlmS8bT1AkPbCcnK12JAWEkB_3Z9pqZSSM3kgv22JxSUdb7Bwk1g7BnyLIHP2ecOygfZR14DHxNmjyl4d836x8JTrGDFx2egl_rZJ8KVX389ds31iSnbtRATuahleF7iRe14iMNVCupnWrRnWpVm1fBYWC7eTaNymUu6EaFsUWeB1XlB-w9vf4B5gY1dg)
* [Class UML Diagram](https://mermaid.live/edit#pako:eNrNVk1P3DAQ_SuWT4vE9gdEq5VQKdIiJCTojXBwkyFYbOytM2lZwf732nHS2ItjZ1EPzSWJ572ZNx928kYLWQLNaLFlTXPJWaVYnYtcEH11a6RmXJC3Ye1gb93icrkm315RsQK_wys2d_CzhQZ9egBgvJH-2ty2uGvxSqqaIYK6B_WLF3ClKVLtiYxZRy99kMUZaVBxUXlaQwqM9I1r-CrrmonS154Q56Rx3UhhwSNMqwnzR5oJfSItldRtumJT3Kmwbkn-34okmrV6Xy5Jojhe7wPT4eY3jtzDo8H4WkLkTkFy5hJhec0qMMYeJ4ckyGbK4hf3homq1chLQHDpEdvMpEPKu5GMC5uiRQW5FTNAt0Tm_egkMNfg6-Mh4bU9Lna12gg9OE-sgHUys0lz9reQXH65A1aCso21z6a6ViJBUxVf4cVvvWlfZCU4cikc9_rNCz8DeLKQaFp2wv-lvm4QfFysGEOPtB9Qgm3XcXJy5P1uJ-Gx7ZURu7Jw6jkMItn2JOvmKD-9J3YKnkGUnw4-9GWmqznIE_KZ425o9Aj80GfXFGrzFDX26fAbHEOGbVm_vrDHodkvP_Z4dD6FvoTWHrJMupxWYfsbDnM872EhIcunhYSc0XNagwbzUv9rdmd1TvEZashpph9Lpl5ymouDxrEW5f1eFDRD1cI5bXclQ-h_TWn2xLYNHP4A4wzVDQ)

# Links
* [Github Repository](https://github.com/freddyouellette/duolingo-transcriber)
* [MIT License](https://github.com/freddyouellette/duolingo-transcriber/blob/master/LICENSE.md)

## Support Me
[![Donate](https://img.shields.io/badge/Donate-fec133?logo=paypal)](https://www.paypal.com/donate/?hosted_button_id=3PJ9XD363CC5E)

Bitcoin: `bc1qs39glh9cwsef0qv40dny6ajnweqe2le7ynfgr2`

Ethereum: `0x5Baba8708b8676afBFF2974b4af4894Fc12aE242`