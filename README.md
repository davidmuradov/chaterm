# Falcon

A CLI XMPP client with OMEMO encryption written in GO

## Contacts Info

To login in the main page:
- Email: test
- Password: test

### xmlcoms moved

The xmlcoms directory has been moved to:
https://github.com/Vacheprime/xmlcoms

### OMEMO (XEP-0384 v0.8.3 and XEP-0454 v0.1.0)

The omemo library is in:
https://github.com/davidmuradov/omemo

This library is, of course, a work in progress.

### Default colors

The default colors used are the ones from the "Nord" color palette.
Eventually, there will be options to change the colors from within
Falcon.

To get 24bit colors (directcolor aka truecolor) in the terminal, users should
set the COLORTERM environment variable to "truecolor" or "24-bit" or "24bit"
according to the tcell documentation.

### Next Features
- Create max sending messages box height so that the sending messages box
does not take over the box where messages are received (TESTING)

- Calculate max sending messages width on screen resizes
