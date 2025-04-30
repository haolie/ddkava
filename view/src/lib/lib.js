let lib = {
  isEmptyObject (o) {
    for (let i in o) {
      return false
    }
    return true
  },
  isEqual (obj, obj2) {
    let status = true
    if (this.isEmptyObject(obj)) {
      if (this.isEmptyObject(obj2)) {
        return true
      } else {
        return false
      }
    } else {
      if (this.isEmptyObject(obj2)) {
        return false
      }
    }
    for (let i in obj) {
      if (obj[i] !== obj2[i]) {
        status = false
        return
      }
    }
    return status
  },
  request ($vue, api, method, params, callback) {
    if (typeof method !== 'string') {
      callback = params
      params = method
      method = 'get'
    }

    if (typeof params !== 'function') {
      callback = params
      params = {}
    }
    $vue.$http[method]({
      api: api,
      params: params,
      callback: function () {
        debugger
      }
    })
  },
  toThousands (num) {
    return (num || 0).toString().replace(/(\d)(?=(?:\d{3})+$)/g, '$1,')
  },
  transNum (num, point, textValue) {
    num = Number(num) * 10000
    var o = {
      value: Math.abs(Number(num)),
      unit: ''
    }
    var syb = Number(num) >= 0 ? 1 : -1
    if (o.value >= 1000000) {
      o.value = o.value/ 100000000
      o.unit = '亿'
    } else if (o.value >= 10000) {
      o.value = o.value / 10000
      o.unit = '万'
    }
    if (typeof point === 'number' && point) {
      var temp = Math.pow(10, point)
      o.value = Math.floor(o.value * temp) / temp
    }
    o.value = syb * o.value
    if (textValue) return o.value + o.unit
    return o
  },
  isInArray (arr, target) {
    let hash = false
    arr.map((o) => {
      if (o === target) hash = true
    })
    return hash
  }
}

export default lib
