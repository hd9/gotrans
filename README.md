# gotrans
A simple #golang tool to translate text using [Azure Translation
Services](https://docs.microsoft.com/en-us/azure/cognitive-services/translator/translator-info-overview).

## Introduction
Microsoft offers an excellent translation API on Azure than can be used (with limitation) for free that 
offers more capacity than [soimort's translate-shell tool](https://github.com/soimort/translate-shell).

I'm a really fan of [soimort's translate-shell tool](https://github.com/soimort/translate-shell)
but recently the limits imposed by the free Apis were very limiting. Also, because
this tool is written in [Golang](https://golang.org) which is more familiar than [AWK](https://en.wikipedia.org/wiki/AWK).

## About Microsoft Translator
Microsoft Translator is a cloud-based machine translation service. The core
service is the Translator Text API, which powers a number of Microsoft products
and services, and is used by thousands of businesses worldwide in their
applications and workflows, which allows their content to reach a global
audience.

### What's offered by the Translator Text API
The Translator Text API is easy to integrate in your applications, websites,
tools, and solutions. It allows you to add multi-language user experiences in
more than 60 languages, and can be used on any hardware platform with any
operating system for text-to-text language translation.

The Translator Text API is part of the Azure Cognitive Services API collection
of machine learning and AI algorithms in the cloud, and is readily consumable
in your development projects

### Pricing
Azure's Cognitive Api provides up to 2M chars of standard translation per month
which should be more than enough for regular users.

### Creating an API Key
This tool requires an Azure Account. Creating an account in Azure is free and
can be done at [https://azure.microsoft.com/free/](https://azure.microsoft.com/free/).

This document describes [how to create a Cognitive Api Key on
Azure](https://docs.microsoft.com/en-us/azure/cognitive-services/translator/translator-text-how-to-signup).

## Installation
Assuming you have Go installed, you can use `go get` (if you're using Go 1.7 or newer):
```
go get -u github.com/hd9/gotrans
```

Make sure the `~/go/bin` path is present in your `$PATH` env var with:
```
PATH=$PATH:~/go/bin
```

Then run `gotrans` with:
```s
gotrans -h
```

## Usage
The tool can be used from the command line and accepts redirections and files
as inputs.

```s
-------------------------------------------------------
gotrans - Translate text using Azure Cognitive Services
-------------------------------------------------------

Usage:
    ./gotrans -t nl "<text-to-translate>"
    ./gotrans -t es -file <file>
    echo "text" | gotrans [OPTIONS]
    cat <file> | gotrans [OPTIONS]

Options:
  -h, --help
        Show help
  - string
        Text to translate
  -f string
        Source Language (optional). Set empty to auto-detect
  -file string
        File name
  -t string
        Target language. See list here: https://bit.ly/37o3PFX
  -v    Run in verbose mode
```

## Vim
The tool can also be called from inside [Vim](//vim.org). Simply select some
text in visual mode and, assuming it's available on your `$PATH`, run:

```
'<,'> !gotrans -t es
```

## Issues
This is alpha software and was developed in less than one day to cover my specific needs
so expect bugs. Feel free to open an issue and contribute. 

Thanks!

## Related Resources
* [What is the Translator Text API?](https://docs.microsoft.com/en-us/azure/cognitive-services/translator/translator-info-overview)
* [Translator Text API 3.0](https://docs.microsoft.com/en-us/azure/cognitive-services/translator/reference/v3-0-translate)
* [Translator Text API 3.0: Languages](https://docs.microsoft.com/en-us/azure/cognitive-services/translator/reference/v3-0-languages)
* [Translator Text documentation](https://docs.microsoft.com/en-us/azure/cognitive-services/translator/)
* [How to sign up for the Translator Text API](https://docs.microsoft.com/en-us/azure/cognitive-services/translator/translator-text-how-to-signup)
* [Quickstart: Use the Translator Text API to translate text](https://docs.microsoft.com/en-us/azure/cognitive-services/translator/quickstart-translate?pivots=programming-language-go)
* [MicrosoftTranslator / Text-Translation-API-V3-Go | GitHub](https://github.com/MicrosoftTranslator/Text-Translation-API-V3-Go)
* [Azure Cognitive Services](https://docs.microsoft.com/en-us/azure/cognitive-services/)
* [translate-shell tool](https://github.com/soimort/translate-shell)

