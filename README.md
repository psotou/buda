# Buda API requests

Makes a simple request to the Buda API in order to get the current market price and variation of the criptocurrency we are interested in.

The way we run the command is:

```bash
$ $GOBIN/./buda <criptocurrency>
```

Where  `criptocurrency` &isin; {btc, eth, ltc, bch}. This call is set so as to deliver information about the `<cryptocurrency>-clp` market.
