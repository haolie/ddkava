
let libUtils={
    // 从数组中移除元素
    removeFromArray:function(ary,cb){
       var tempList=[]
       for(var i=ary.length-1;i>=0;i--){
          var item=ary.pop()
          var r=cb(i,item)
          if (r==0){
            tempList.push(ary[i])
            continue
          }

          if(r==1){
            break;
          }
       }

       for(var i=tempList.length-1;i>=0;i--){
          ary.push(tempList[i])
       }
    }, 
}

 

export default libUtils