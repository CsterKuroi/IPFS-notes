var IPFS = require('ipfs-http-client');

var ipfs = IPFS();

hash = "QmXyfZv1wW1HxnLQT4ZdJXWqwCRrqpnbLBNT4vQanXfCrY";

ipfs.cat(hash, function (err, res) {
    if (err || !res) {
        return console.error('ipfs cat error', err, res)
    }
    console.log(hash);
    console.log(res.toString())
});

var toStore = "hello ipfs from js jihao";
ipfs.add(Buffer.from(toStore), function (err, res) {
    if (err || !res) {
        return console.error('ipfs add error', err, res)
    }

    res.forEach(function (file) {
        if (file && file.hash) {
            console.log('successfully stored', file.hash)
        }
    })
});
