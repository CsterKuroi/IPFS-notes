'use strict'

var IPFS = require('ipfs-http-client')

var ipfs = IPFS('/ip4/127.0.0.1/tcp/5001')

const identity = ipfs.id()
console.log(identity)

function store () {
  var toStore = document.getElementById('source').value
  ipfs.add(Buffer.from(toStore), function (err, res) {
    if (err || !res) {
      return console.error('ipfs add error', err, res)
    }

    res.forEach(function (file) {
      if (file && file.hash) {
        console.log('successfully stored', file.hash)
        display(file.hash)
      }
    })
  })
}

function display (hash) {
  ipfs.cat(hash, function (err, res) {
    if (err || !res) {
      return console.error('ipfs cat error', err, res)
    }

    document.getElementById('hash').innerText = hash
    document.getElementById('content').innerText = res.toString()
  })
}

document.addEventListener('DOMContentLoaded', function () {
  document.getElementById('store').onclick = store
})
