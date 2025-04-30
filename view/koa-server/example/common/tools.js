var Mock = require('mockjs');


const obj=function () {}
obj.prototype.mocNameValuesArray=function (names,count,valStr,valueStr,order) {
  names=[].concat(names);
  if(count){
    if(typeof count =="string"){
      count=count.indexOf("-")>0?roolNum(count):count
    }
  }else count=names.length;

  if(!valStr)valStr="20-520";
  if(!valueStr)valueStr="value"

  var list=[];
  while (list.length<count&&names.length){
    var i=0
    if(names.length>1&&!order)i=roolNum(0,names.length-1)
    var temp={name:names[i]};
    names.splice(i,1);
    temp[valueStr+"|"+valStr ]=10
    list.push(temp)
  }

  return Mock.mock(list);
}

obj.prototype.NameObj={
  productNames:["三星波轮式洗衣机","三星 Galaxy Note7 智能手机","三星笔记本电脑","三星智能电视","三星硬盘","朗逸","途观","迈腾","麦芒3","荣耀7","华为P8","嘉陵","特步","李宁","李维斯","捷安特","美利达"],
  brandNames:["乐高","飞利浦","佳能","三星","大众","宝马","LG","华为","莎普爱思","大众","宝马","雪佛兰","法拉利","兰博基尼","保时捷","阿斯顿马丁","劳斯拉斯"],
  tradeNames:["汽车","电器","玩具","摩托车","手机","农机","自行车","服装"],
  infoSource:["社交媒体","新闻","综合","纸媒","博客","社交媒体","论坛","其他"],
  partsNames:["刹车","车门","空调","大灯","安全带","耳机","显示器","屏幕","键盘"]
}

function roolNum(min,max,str) {
  if(!min)min=0
  if(!str)str=min+"-"+max;


  var temp={};temp["num|"+str]=0
  return Mock.mock(temp).num;
}

obj.prototype.mapNameValuePer=function(){
  var list=["湖南","河南","上海","天津","浙江","四川","重庆","其他"],result=[];
  list.forEach(function (l) {
    result.push({
      name:l,
      "value|10-5000":10,
      "persent|1-99":10
    })
  })

  return Mock.mock(result);
}

module.exports=new obj();
