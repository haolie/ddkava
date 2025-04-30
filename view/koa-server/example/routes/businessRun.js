var Mock = require('mockjs')
import {request, summary, body, tags, middlewares, path, description, okMsg, lib, errMsg} from '../../lib'

const tag = tags(['大屏'])

export default class businessRouter {
  @request('get', '/business/businessRun/totalItem')
  @summary('大屏态势-四宫格指标')
  @tag
  @description('大屏态势-四宫格指标')
  static async totalItem (ctx) {
    okMsg.data = Mock.mock([
      {
        'itemId|+1': 1,
        itemName: '人均营收',
        'itemValue|80000-900000000': 200000,
        'itemPer|1-100.1-10': 1.2,
        sort: /(0|1)/,
      },
      {
        'itemId|+1': 2,
        itemName: '人均利润',
        'itemValue|80000-900000000': 200000,
        'itemPer|1-100.1-10': 1.2,
        sort: /(0|1)/,
      },
      {
        'itemId|+1': 3,
        itemName: '资产总额',
        'itemValue|80000-900000000': 200000,
        'itemPer|1-100.1-10': 1.2,
        sort: /(0|1)/,
      },
      {
        'itemId|+1': 4,
        itemName: '科技指出及技术投入',
        'itemValue|80000-900000000': 200000,
        'itemPer|1-100.1-10': 1.2,
        sort: /(0|1)/,
      }
    ])
    ctx.body = okMsg
  }

  @request('get', '/business/businessRun/unitMemSort')
  @summary('成员单位指标排名')
  @tag
  @description('成员单位指标排名')
  static async unitMenSort (ctx) {
    okMsg.data = []
    var names = ['卫士通', '三十所', '三十三所', '网安本部', '凯天', '神风', '新瑞','凯瑞','卫视都','凯峰','新田','凯氏同','卫视峰','科技创新','国际业务','军工电子']
    for (var i = 0; i < names.length; i++) {
      okMsg.data.push(
        Mock.mock({
          'unitMemId': i,
          unitMem: names[i],
          quotaName: '营业收入',
          quotaNum: '01',
          'targetValue|80000-900000000': 200000,
          'finishedValue|-200000000-900000000': 200000,
          'noFinishedValue|80000-900000000': 200000,
          statMonth: '201910',
          'finishedPer|0-0.2': 0.2,
          'index|+1': 0,
        })
      )
      okMsg.data=[
        {
          "unitMemId": 0,
          "unitMem": "卫士通",
          "quotaName": "营业收入",
          "quotaNum": "01",
          "targetValue": 5250,
          "finishedValue": 4478.68,
          "noFinishedValue": 771.3198,
          "statMonth": "201910",
          "finishedPer": 85.31,
          "index": 0
        },
        {
          "unitMemId": 1,
          "unitMem": "三十所",
          "quotaName": "营业收入",
          "quotaNum": "01",
          "targetValue": 13636.1,
          "finishedValue": 10760.04,
          "noFinishedValue": 2872.0896,
          "statMonth": "201910",
          "finishedPer": 78.93,
          "index": 0
        },
        {
          "unitMemId": 2,
          "unitMem": "三十三所",
          "quotaName": "营业收入",
          "quotaNum": "01",
          "targetValue": 28533.45,
          "finishedValue": 19167,
          "noFinishedValue": 9366.449,
          "statMonth": "201910",
          "finishedPer": 67.176,
          "index": 0
        },
        {
          "unitMemId": 3,
          "unitMem": "网安本部",
          "quotaName": "营业收入",
          "quotaNum": "01",
          "targetValue": 84100,
          "finishedValue": 51287.72,
          "noFinishedValue": 32812.28,
          "statMonth": "201910",
          "finishedPer": 60.98,
          "index": 0
        },
        {
          "unitMemId": 4,
          "unitMem": "凯天",
          "quotaName": "营业收入",
          "quotaNum": "01",
          "targetValue": 13896.22,
          "finishedValue": 7775.35,
          "noFinishedValue": 6120.3696,
          "statMonth": "201910",
          "finishedPer": 55.96,
          "index": 0
        },
        {
          "unitMemId": 5,
          "unitMem": "神风",
          "quotaName": "营业收入",
          "quotaNum": "01",
          "targetValue": 223000,
          "finishedValue": 117755.93,
          "noFinishedValue": 105244.07,
          "statMonth": "201910",
          "finishedPer":52.81,
          "index": 0
        },
        {
          "unitMemId": 6,
          "unitMem": "新瑞",
          "quotaName": "营业收入",
          "quotaNum": "01",
          "targetValue": 12915.84,
          "finishedValue": 6715.21,
          "noFinishedValue": 6200.63,
          "statMonth": "201910",
          "finishedPer": 51.99,
          "index": 0
        },
        {
          "unitMemId": 7,
          "unitMem": "凯瑞",
          "quotaName": "营业收入",
          "quotaNum": "01",
          "targetValue": 12915.84,
          "finishedValue": 6515.11,
          "noFinishedValue": 6400.73,
          "statMonth": "201910",
          "finishedPer": 50.44,
          "index": 0
        },
        {
          "unitMemId": 8,
          "unitMem": "卫视都",
          "quotaName": "营业收入",
          "quotaNum": "01",
          "targetValue": 17101.61,
          "finishedValue": 8568.5,
          "noFinishedValue": 8533.109,
          "statMonth": "201910",
          "finishedPer": 50.10,
          "index": 0
        },
        {
          "unitMemId": 9,
          "unitMem": "凯峰",
          "quotaName": "营业收入",
          "quotaNum": "01",
          "targetValue": 208775.31,
          "finishedValue": 103061.64,
          "noFinishedValue": 105713.67,
          "statMonth": "201910",
          "finishedPer": 49.36,
          "index": 0
        },
        {
          "unitMemId": 10,
          "unitMem": "新田",
          "quotaName": "营业收入",
          "quotaNum": "01",
          "targetValue": 16670.4,
          "finishedValue": 7871.72,
          "noFinishedValue": 8798.68,
          "statMonth": "201910",
          "finishedPer": 47.22,
          "index": 0
        },
        {
          "unitMemId": 11,
          "unitMem": "凯氏同",
          "quotaName": "营业收入",
          "quotaNum": "01",
          "targetValue": 15434.23,
          "finishedValue": 7101.45,
          "noFinishedValue": 8332.78,
          "statMonth": "201910",
          "finishedPer": 46.01,
          "index": 0
        },
        {
          "unitMemId": 12,
          "unitMem": "卫视峰",
          "quotaName": "营业收入",
          "quotaNum": "01",
          "targetValue": 3233.76,
          "finishedValue": 1140.08,
          "noFinishedValue": 2093.6802,
          "statMonth": "201910",
          "finishedPer": 35.26,
          "index": 0
        },
        {
          "unitMemId": 13,
          "unitMem": "科技创新",
          "quotaName": "营业收入",
          "quotaNum": "01",
          "targetValue": 4060,
          "finishedValue": 1373.33,
          "noFinishedValue": 2686.67,
          "statMonth": "201910",
          "finishedPer": 33.83,
          "index": 0
        },
        {
          "unitMemId": 14,
          "unitMem": "国际业务",
          "quotaName": "营业收入",
          "quotaNum": "01",
          "targetValue": 8185.34,
          "finishedValue": 2354.39,
          "noFinishedValue": 5830.96,
          "statMonth": "201910",
          "finishedPer": 28.76,
          "index": 0
        },
        {
          "unitMemId": 15,
          "unitMem": "军工电子",
          "quotaName": "营业收入",
          "quotaNum": "01",
          "targetValue": 13377,
          "finishedValue": 284.06,
          "noFinishedValue": 13092.94,
          "statMonth": "201910",
          "finishedPer": 2.12,
          "index": 0
        },
        {
          "unitMemId": 16,
          "unitMem": "zhazi",
          "quotaName": "营业收入",
          "quotaNum": "01",
          "targetValue": 0,
          "finishedValue": 356210.72,
          "noFinishedValue": -356210.72,
          "statMonth": "201910",
          "finishedPer": 0,
          "index": 0
        }
      ]
      ctx.body = okMsg
    }
  }

  @request('get', '/business/businessRun/quotaFinishedItem')
  @summary('成员单位完成率低于某个百分比')
  @tag
  @description('成员单位完成率低于某个百分比')
  static async quotaFinishedItem (ctx) {
    okMsg.data = []
    var names = ['卫士通', '三十所', '三十三所', '网安本部']
    for (var i = 0; i < names.length; i++) {
      okMsg.data.push(
        Mock.mock({
          'unitMemId': i,
          unitMem: names[i],
          quotaName: '营业收入',
          quotaNum: '01',
          'targetValue|80000-900000000': 200000,
          'finishedValue|80000-900000000': 200000,
          'nofinishedValue|80000-900000000': 200000,
          statMonth: '201910',
          'finishedPer|0-0.2': 0.2,
          'index|+1': 0,
        })
      )
      ctx.body = okMsg
    }
  }

  @request('get', '/business/businessRun/finishedUnitMemTop3')
  @summary('完成率排名前三单位')
  @tag
  @description('完成率排名前三单位')
  static async finishedUnitMemTop3 (ctx) {
    okMsg.data = []
    var names = ['卫士通', '三十所', '三十三所']
    for (var i = 0; i < names.length; i++) {
      okMsg.data.push(
        Mock.mock({
          'unitMemId': i,
          unitMem: names[i],
          quotaName: '营业收入',
          quotaNum: '01',
          'targetValue|80000-900000000': 200000,
          'finishedValue|80000-900000000': 200000,
          'noFinishedValue|80000-900000000': 200000,
          statMonth: '201910',
          'finishedPer|-0.3-100': 0.2,
          'index|+1': 0,
        })
      )
      ctx.body = okMsg
    }
  }

  @request('get', '/business/businessRun/wdytValuePie')
  @summary('五大业态统计接口')
  @tag
  @description('五大业态统计接口')
  static async wdytValuePie (ctx) {
    okMsg.data = Mock.mock([
      {
        name: '其它收入',
        'value|8000-900000000': 200000,
        'per|1-100.1-2': 1.2
      },
      {
        name: '军工电子',
        'value|8000-900000000': 200000,
        'per|1-100.1-2': 1.2
      },
      {
        name: '国际化经营',
        'value|8000-900000000': 200000,
        'per|1-100.1-2': 1.2
      },
      {
        name: '民品产业',
        'value|8000-900000000': 200000,
        'per|1-100.1-2': 1.2
      },
      {
        name: '科技创新',
        'value|8000-900000000': 200000,
        'per|1-100.1-2': 1.2
      },
      {
        name: '资产经营和资本运作',
        'value|8000-900000000': 200000,
        'per|1-100.1-2': 1.2
      },
    ])
    ctx.body = okMsg
  }

  @request('get', '/business/businessRun/quotaInfoList')
  @summary('更多指标信息')
  @tag
  @description('更多指标信息')
  static async quotaInfoList (ctx) {
    okMsg.data = Mock.mock([{
      name: '两金占比',
      valueA: '15.2%',
      valueB: '35%',
      'percent|1-100.2': 40,
      'changeValue|1-100': 12,
      'changeType|0-1': 0
    },
    {
      name: '全员劳动生产率',
      valueA: '12.6万元人年',
      valueB: '18万元人年',
      'percent|1-100.2': 40,
      'changeValue|1-100': 12,
      'changeType|0-1': 0
    }])
    ctx.body = okMsg
  }

  @request('get', '/business/businessRun/mainQuotaInfos')
  @summary('主要指标信息')
  @tag
  @description('主要指标信息')
  static async mainQuotaInfos (ctx) {
    var names=['营业收入','净利润','新签合同','EVA值']
    okMsg.data=[]
    for(var i=0;i<names.length;i++){
      okMsg.data.push( Mock.mock( {
        name:names[i],
        'targetValue|100000000-900000000': 200000,
        'finishedValue|80000-90000000': 200000,
        'changePer|0-0.2': 0.23,
        'changeType|0-1': 0
      }))
    }
    ctx.body = okMsg
  }

  @request('get', '/business/businessRun/quotaAnalysis')
  @summary('指标信息')
  @tag
  @description('指标信息')
  static async quotaAnalysis (ctx) {
    okMsg.data={
      name:'1-9月',
      'targetValue': 200000000,
      'finishedValue': 150000000,
      items:[
        {
          name:'1-3月',
          value:50000000,
        },
        {
          name:'1-6月',
          value:100000000,
        },
        {
          name:'1-9月',
          value:150000000,
        }
      ]
    }

    okMsg.data={
      name:'1-9月',
      'total': 200000000,
      'completed': 150000000,
      completedGrowthTate:1.5,
      quarterData:[
        {
          name:'1-3月',
          value:50000000,
          quarterGrowthTate:0.8
        },
        {
          name:'1-6月',
          value:100000000,
          quarterGrowthTate:0.8
        },
        {
          name:'1-9月',
          value:150000000,
          quarterGrowthTate:0.8
        }
      ]
    }
    ctx.body = okMsg
  }

  @request('get', '/business/businessRun/lowerThanAverage')
  @summary('低于平均值')
  @tag
  @description('低于平均值')
  static async lowerThanAverage (ctx) {
    okMsg.data={
      averagePer:0.35,
      items:[
        {
          name:'欣瑞',
          targetValue:50000000,
          finishedValue:10000000,
          finishedPer:0.20
        },
        {
          name:'神风',
          targetValue:40000000,
          finishedValue:10000000,
          finishedPer:0.25
        },
        {
          name:'凯天',
          targetValue:80000000,
          finishedValue:10000000,
          finishedPer:0.125
        },
      ]
    }
    ctx.body = okMsg
  }

  @request('get', '/business/businessRun/khzqValue')
  @summary('考核周期与综合评分')
  @tag
  @description('考核周期与综合评分')
  static async khzqValue (ctx) {
    okMsg.data=Mock.mock({
      'dayNum|0-70': 3,
      'index|0-100': 3,
      'statusValue|0-1': 0,
    })
    ctx.body = okMsg
  }

  @request('get', '/business/businessRun/noFinishedValue')
  @summary('单位未完成率')
  @tag
  @description('单位未完成率')
  static async noFinishedValue (ctx) {
    var profit={
      legend: ['新欣神风','鹏跃','华北网安','特装','杂志社','网安'],
      target:[15434.23,233233.76,4060,8185.34,13377,0.0],
      done:[7101.45,-1140.08,1373.33,-2354.39,286.06,356210.72],
      percent:[
        {percent:'46.01'},{percent:'35.26'},{percent:'33.83'},{percent:'28.76'},{percent:'2.12'},{percent:'0'},
      ],
      info:{percent:'38%',num:'6'}
    }


    okMsg.data={
      legend: ['神风','凯乐','欣瑞'],
      target:[32000,50000,6000],
      done:[10000,5000,500],
      percent:[
        {percent:'30%'},{percent:'20%'},{percent:'10%'},
      ],
      info:{percent:'38%',num:''}
    }
    okMsg.data=profit
    ctx.body = okMsg
  }

  @request('get', '/business/businessRun/subQuotaItem')
  @summary('次级指标查询')
  @tag
  @description('次级指标查询')
  static async subQuotaItem (ctx) {
    var items=[]
    items.push({
      name:'人均营业收入',
      value:'4500000000',
      change:0.23
    },
      {
        name:'营业成本费用',
        value:'1440000000',
        change:-0.23
      }
      );

    okMsg.data=items
    ctx.body = okMsg
  }

  @request('get', '/business/businessRun/last4ProfitByPerson')
  @summary('单位人均净利润低于网安平均水平')
  @tag
  @description('单位人均净利润低于网安平均水平')
  static async last4ProfitByPerson (ctx) {
    var items=[     {"name": "欣瑞",  //单位名称
      "value": 0.5,  //人均利润
    },
      {
        "name": "神风",
        "targetValue": -1.5,
      },
      {
        "name": "凯天",
        "targetValue": -8,
      },
      {
        "name": "凯峰",
        "targetValue": -8,
      }]


    okMsg.data=items
    ctx.body = okMsg
  }

  @request('get', '/business/businessRun/notifyInfoList')
  @summary('滚动提示信息')
  @tag
  @description('滚动提示信息')
  static async notifyInfoList (ctx) {
    var items=[
      {
        id:1,
        info:'中国网安1-9月整体营收10.5亿，完成率不足 50%，距离考核周期结束不足 90 天'
      },
      {
        id:1,
        info:'卫士通1-10月整体营收10.5亿，完成率不足 50%，距离考核周期结束不足 90 天'},
      {
        id:1,
        info:'神风1-11月整体营收10.5亿，完成率不足 50%，距离考核周期结束不足 90 天'
      }]

    okMsg.data=items
    ctx.body = okMsg
  }
}
