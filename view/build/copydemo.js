var ora = require('ora')
var rm = require('rimraf')
var fs = require('fs')
var path = require('path')
var chalk = require('chalk')

function copy(src, dst) {
  fs.createReadStream(src).pipe(fs.createWriteStream(dst))
}

function travel(dir, todir, level) {
  fs.stat(todir, function (err2, stats) {
    if (err2) fs.mkdirSync(todir)
    fs.readdir(dir, function (err, files) {
      if (err) {
        throw err
      } else {
        if (files.length > 0) {
          files.forEach(o => {
            var pathname = path.join(dir, o)
            var topath = path.join(todir, o)

            fs.stat(pathname, function (err1, stats) {
              if (stats.isDirectory()) {
                travel(pathname, topath, level + 1)
              } else {
                copy(pathname, topath)
              }
            })
          })
          if (level === 1) {

          }
        }
      }
    })
  })
}
